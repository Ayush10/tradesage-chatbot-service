package notification

import (
	"context"
	"time"

	"github.com/google/uuid"
	pb "github.com/tradesage-chatbot-service/internal/pb"
	"github.com/tradesage-chatbot-service/models"
)

type Service struct {
	pb.UnimplementedNotificationServiceServer
	store *models.NotificationStore
}

func NewService() *Service {
	return &Service{
		store: models.NewNotificationStore(),
	}
}

func (s *Service) SendNotification(ctx context.Context, req *pb.SendNotificationRequest) (*pb.SendNotificationResponse, error) {
	notification := &models.Notification{
		ID:        uuid.New().String(),
		UserID:    req.UserId,
		Title:     req.Title,
		Content:   req.Content,
		Read:      false,
		CreatedAt: time.Now(),
	}

	if err := s.store.Save(notification); err != nil {
		return nil, err
	}

	return &pb.SendNotificationResponse{
		Success: true,
		Message: "Notification sent successfully",
	}, nil
}

func (s *Service) GetNotifications(ctx context.Context, req *pb.GetNotificationsRequest) (*pb.GetNotificationsResponse, error) {
	notifications, total := s.store.GetUserNotifications(req.UserId, int(req.Page), int(req.Limit))

	pbNotifications := make([]*pb.Notification, len(notifications))
	for i, n := range notifications {
		pbNotifications[i] = &pb.Notification{
			Id:        n.ID,
			UserId:    n.UserID,
			Title:     n.Title,
			Content:   n.Content,
			Read:      n.Read,
			CreatedAt: n.CreatedAt.Unix(),
		}
	}

	return &pb.GetNotificationsResponse{
		Notifications: pbNotifications,
		Total:         int32(total),
	}, nil
}

func (s *Service) MarkAsRead(ctx context.Context, req *pb.MarkAsReadRequest) (*pb.MarkAsReadResponse, error) {
	err := s.store.MarkAsRead(req.NotificationId, req.UserId)
	if err != nil {
		return nil, err
	}

	return &pb.MarkAsReadResponse{
		Success: true,
	}, nil
}
