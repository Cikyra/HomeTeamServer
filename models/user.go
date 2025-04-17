package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `gorm:"type:uuid;default:gen_random_uuid()" json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	PhotoUrls []string  `gorm:"type:text[]" json:"photo_urls"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	// TODO: Add Children
	// TODO: Add RegisteredEvents
}
