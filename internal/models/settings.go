package models

type Settings struct {
	UserID               string `json:"user_id"`
	Theme                string `json:"theme"`
	Language             string `json:"language"`
	NotificationsEnabled bool   `json:"notifications_enabled"`
	Timezone             string `json:"timezone"`
}

type SettingsStore struct {
	settings map[string]*Settings // key: userID
}

func NewSettingsStore() *SettingsStore {
	return &SettingsStore{
		settings: make(map[string]*Settings),
	}
}

func (s *SettingsStore) GetSettings(userID string) (*Settings, error) {
	settings, exists := s.settings[userID]
	if !exists {
		// Return default settings if not found
		return &Settings{
			UserID:               userID,
			Theme:                "light",
			Language:             "en",
			NotificationsEnabled: true,
			Timezone:             "UTC",
		}, nil
	}
	return settings, nil
}

func (s *SettingsStore) UpdateSettings(settings *Settings) error {
	s.settings[settings.UserID] = settings
	return nil
}
