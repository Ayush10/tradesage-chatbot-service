package main

import (
	"log"
	"net"
	"os"

	pb "github.com/Ayush10/tradesage-chatbot-service/internal/pb"
	"github.com/Ayush10/tradesage-chatbot-service/internal/services"
	messagingService "github.com/Ayush10/tradesage-chatbot-service/internal/services/messaging"
	notificationService "github.com/Ayush10/tradesage-chatbot-service/internal/services/notification"
	settingsService "github.com/Ayush10/tradesage-chatbot-service/internal/services/settings"
	userProfileService "github.com/Ayush10/tradesage-chatbot-service/internal/services/user_profile"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	port := os.Getenv("CHATBOT_SERVICE_PORT")
	if port == "" {
		port = "50051"
	}

	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	chatbotServer, err := services.NewChatbotServiceServer()
	if err != nil {
		log.Fatalf("Failed to create chatbot server: %v", err)
	}
	pb.RegisterChatbotServiceServer(grpcServer, chatbotServer)

	// Register Messaging Service
	msgServer := messagingService.NewService()
	pb.RegisterMessagingServiceServer(grpcServer, msgServer)

	// Register Notification Service
	notifyServer := notificationService.NewService()
	pb.RegisterNotificationServiceServer(grpcServer, notifyServer)

	// Register Settings Service
	settingsServer := settingsService.NewService()
	pb.RegisterSettingsServiceServer(grpcServer, settingsServer)

	// Register User Profile Service
	userProfileServer := userProfileService.NewService()
	pb.RegisterUserProfileServiceServer(grpcServer, userProfileServer)

	// Register reflection service on gRPC server
	reflection.Register(grpcServer)

	log.Printf("TradeSage AI - Chatbot Integration Service is running on port %s.", port)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
