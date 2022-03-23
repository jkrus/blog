package models

import (
	"time"
)

type (
	Note struct {
		ID        uint64    `json:"id" gorm:"primaryKey"`
		Header    string    `json:"header" gorm:"header"`
		Body      string    `json:"body" gorm:"body"`
		Tags      string    `json:"tags" gorm:"tags"`
		CreatedAt time.Time `json:"createdAt" gorm:"autoCreateTime"`
	}

	NoteDTO struct {
		Header string   `json:"header"`
		Body   string   `json:"body"`
		Tags   []string `json:"tags"`
	}
)
