package models

import "time"

type Notification struct {
	ID        string    `json:"id"`
	UserID    string    `json:"user_id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	Read      bool      `json:"read"`
	CreatedAt time.Time `json:"created_at"`
}

type NotificationStore struct {
	notifications map[string][]*Notification // key: userID
}

func NewNotificationStore() *NotificationStore {
	return &NotificationStore{
		notifications: make(map[string][]*Notification),
	}
}

func (s *NotificationStore) Save(notification *Notification) error {
	s.notifications[notification.UserID] = append(s.notifications[notification.UserID], notification)
	return nil
}

func (s *NotificationStore) GetUserNotifications(userID string, page, limit int) ([]*Notification, int) {
	notifications := s.notifications[userID]
	total := len(notifications)

	start := page * limit
	if start >= total {
		return []*Notification{}, total
	}

	end := start + limit
	if end > total {
		end = total
	}

	return notifications[start:end], total
}

func (s *NotificationStore) MarkAsRead(notificationID, userID string) error {
	for _, notification := range s.notifications[userID] {
		if notification.ID == notificationID {
			notification.Read = true
			break
		}
	}
	return nil
}
