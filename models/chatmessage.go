package models

import "time"

type ChatMessage struct {
	Content     string    `json:"content"`
	Author      string    `json:"author"`
	SubmittedAt time.Time `json:"publishedAt"`
}
