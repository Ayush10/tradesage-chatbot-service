package models

import "time"

type Message struct {
	ID          string    `json:"id"`
	SenderID    string    `json:"sender_id"`
	RecipientID string    `json:"recipient_id"`
	Content     string    `json:"content"`
	Read        bool      `json:"read"`
	SentAt      time.Time `json:"sent_at"`
}

type MessageStore struct {
	messages map[string][]*Message // key: userID
}

func NewMessageStore() *MessageStore {
	return &MessageStore{
		messages: make(map[string][]*Message),
	}
}

func (s *MessageStore) Save(msg *Message) error {
	// Store message for both sender and recipient
	s.messages[msg.SenderID] = append(s.messages[msg.SenderID], msg)
	s.messages[msg.RecipientID] = append(s.messages[msg.RecipientID], msg)
	return nil
}

func (s *MessageStore) GetUserMessages(userID string, page, limit int) ([]*Message, int) {
	messages := s.messages[userID]
	total := len(messages)

	start := page * limit
	if start >= total {
		return []*Message{}, total
	}

	end := start + limit
	if end > total {
		end = total
	}

	return messages[start:end], total
}
