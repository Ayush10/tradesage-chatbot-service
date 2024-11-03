package settings

import (
	"context"

	"github.com/tradesage-chatbot-service/models"
	pb "github.com/tradesage-chatbot-service/pb"
)

type Service struct {
	pb.UnimplementedSettingsServiceServer
	store *models.SettingsStore
}

func NewService() *Service {
	return &Service{
		store: models.NewSettingsStore(),
	}
}

func (s *Service) GetSettings(ctx context.Context, req *pb.GetSettingsRequest) (*pb.GetSettingsResponse, error) {
	settings, err := s.store.GetSettings(req.UserId)
	if err != nil {
		return nil, err
	}

	return &pb.GetSettingsResponse{
		Settings: &pb.Settings{
			UserId:               settings.UserID,
			Theme:                settings.Theme,
			Language:             settings.Language,
			NotificationsEnabled: settings.NotificationsEnabled,
			Timezone:             settings.Timezone,
		},
	}, nil
}

func (s *Service) UpdateSettings(ctx context.Context, req *pb.UpdateSettingsRequest) (*pb.UpdateSettingsResponse, error) {
	settings := &models.Settings{
		UserID:               req.Settings.UserId,
		Theme:                req.Settings.Theme,
		Language:             req.Settings.Language,
		NotificationsEnabled: req.Settings.NotificationsEnabled,
		Timezone:             req.Settings.Timezone,
	}

	if err := s.store.UpdateSettings(settings); err != nil {
		return nil, err
	}

	return &pb.UpdateSettingsResponse{
		Success: true,
		Message: "Settings updated successfully",
	}, nil
}
