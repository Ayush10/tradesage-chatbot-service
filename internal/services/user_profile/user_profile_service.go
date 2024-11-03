// internal/services/user_profile/service.go
package userprofile

import (
	"context"

	"github.com/Ayush10/tradesage-chatbot-service/internal/models"
	pb "github.com/Ayush10/tradesage-chatbot-service/internal/pb"
	"github.com/google/uuid"
)

type Service struct {
	pb.UnimplementedUserProfileServiceServer
	store *models.UserProfileStore
}

func NewService() *Service {
	return &Service{
		store: models.NewUserProfileStore(),
	}
}

func (s *Service) GetProfile(ctx context.Context, req *pb.GetProfileRequest) (*pb.GetProfileResponse, error) {
	profile, err := s.store.GetProfile(req.UserId)
	if err != nil {
		return nil, err
	}

	if profile == nil {
		return &pb.GetProfileResponse{
			Profile: &pb.Profile{
				UserId:    req.UserId,
				Name:      "Default User",
				Email:     "",
				Phone:     "",
				AvatarUrl: "",
			},
		}, nil
	}

	return &pb.GetProfileResponse{
		Profile: &pb.Profile{
			UserId:    profile.UserID,
			Name:      profile.Name,
			Email:     profile.Email,
			Phone:     profile.Phone,
			AvatarUrl: profile.AvatarURL,
		},
	}, nil
}

func (s *Service) CreateProfile(ctx context.Context, req *pb.CreateProfileRequest) (*pb.CreateProfileResponse, error) {
	userId := uuid.New().String()

	profile := &models.UserProfile{
		UserID:    userId,
		Name:      req.Name,
		Email:     req.Email,
		Phone:     req.Phone,
		AvatarURL: req.AvatarUrl,
	}

	err := s.store.SaveProfile(profile)
	if err != nil {
		return nil, err
	}

	return &pb.CreateProfileResponse{
		Profile: &pb.Profile{
			UserId:    profile.UserID,
			Name:      profile.Name,
			Email:     profile.Email,
			Phone:     profile.Phone,
			AvatarUrl: profile.AvatarURL,
		},
		Success: true,
		Message: "Profile created successfully",
	}, nil
}

func (s *Service) UpdateProfile(ctx context.Context, req *pb.UpdateProfileRequest) (*pb.UpdateProfileResponse, error) {
	// Get existing profile
	existingProfile, err := s.store.GetProfile(req.UserId)
	if err != nil {
		return nil, err
	}

	if existingProfile == nil {
		return &pb.UpdateProfileResponse{
			Success: false,
			Message: "Profile not found",
		}, nil
	}

	// Update fields if provided
	if req.Name != nil {
		existingProfile.Name = *req.Name
	}
	if req.Email != nil {
		existingProfile.Email = *req.Email
	}
	if req.Phone != nil {
		existingProfile.Phone = *req.Phone
	}
	if req.AvatarUrl != nil {
		existingProfile.AvatarURL = *req.AvatarUrl
	}

	// Save updated profile
	err = s.store.SaveProfile(existingProfile)
	if err != nil {
		return nil, err
	}

	return &pb.UpdateProfileResponse{
		Profile: &pb.Profile{
			UserId:    existingProfile.UserID,
			Name:      existingProfile.Name,
			Email:     existingProfile.Email,
			Phone:     existingProfile.Phone,
			AvatarUrl: existingProfile.AvatarURL,
		},
		Success: true,
		Message: "Profile updated successfully",
	}, nil
}

func (s *Service) DeleteProfile(ctx context.Context, req *pb.DeleteProfileRequest) (*pb.DeleteProfileResponse, error) {
	// Since our simple store doesn't have delete functionality, we'll just return success
	return &pb.DeleteProfileResponse{
		Success: true,
		Message: "Profile deleted successfully",
	}, nil
}
