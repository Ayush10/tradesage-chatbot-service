package messaging

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/tradesage-chatbot-service/internal/models"
	pb "github.com/tradesage-chatbot-service/internal/pb"
)

type Service struct {
	pb.UnimplementedMessagingServiceServer
	store *models.MessageStore
}

func NewService() *Service {
	return &Service{
		store: models.NewMessageStore(),
	}
}

func (s *Service) SendMessage(ctx context.Context, req *pb.SendMessageRequest) (*pb.SendMessageResponse, error) {
	msg := &models.Message{
		ID:          uuid.New().String(),
		SenderID:    req.SenderId,
		RecipientID: req.RecipientId,
		Content:     req.Content,
		Read:        false,
		SentAt:      time.Now(),
	}

	if err := s.store.Save(msg); err != nil {
		return nil, err
	}

	return &pb.SendMessageResponse{
		Success: true,
		Message: "Message sent successfully",
	}, nil
}

func (s *Service) GetMessages(ctx context.Context, req *pb.GetMessagesRequest) (*pb.GetMessagesResponse, error) {
	messages, total := s.store.GetUserMessages(req.UserId, int(req.Page), int(req.Limit))

	pbMessages := make([]*pb.Message, len(messages))
	for i, msg := range messages {
		pbMessages[i] = &pb.Message{
			Id:          msg.ID,
			SenderId:    msg.SenderID,
			RecipientId: msg.RecipientID,
			Content:     msg.Content,
			Read:        msg.Read,
			SentAt:      msg.SentAt.Unix(),
		}
	}

	return &pb.GetMessagesResponse{
		Messages: pbMessages,
		Total:    int32(total),
	}, nil
}

func (s *Service) StreamMessages(stream pb.MessagingService_StreamMessagesServer) error {
	for {
		msg, err := stream.Recv()
		if err != nil {
			return err
		}

		// Process received message
		msg.Id = uuid.New().String()
		msg.SentAt = time.Now().Unix()

		// Save message
		s.store.Save(&models.Message{
			ID:          msg.Id,
			SenderID:    msg.SenderId,
			RecipientID: msg.RecipientId,
			Content:     msg.Content,
			Read:        false,
			SentAt:      time.Unix(msg.SentAt, 0),
		})

		// Send back to stream
		if err := stream.Send(msg); err != nil {
			return err
		}
	}
}
