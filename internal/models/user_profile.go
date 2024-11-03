package models

type UserProfile struct {
	UserID    string `json:"user_id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	AvatarURL string `json:"avatar_url"`
}

type UserProfileStore struct {
	profiles map[string]*UserProfile // key: userID
}

func NewUserProfileStore() *UserProfileStore {
	return &UserProfileStore{
		profiles: make(map[string]*UserProfile),
	}
}

func (s *UserProfileStore) GetProfile(userID string) (*UserProfile, error) {
	profile, exists := s.profiles[userID]
	if !exists {
		return nil, nil
	}
	return profile, nil
}

func (s *UserProfileStore) SaveProfile(profile *UserProfile) error {
	s.profiles[profile.UserID] = profile
	return nil
}
