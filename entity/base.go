package entity

import (
	"time"

	"github.com/google/uuid"
)

type Base struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey;not null" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}

func (b *Base) beforeCreate() error {
	id, err := uuid.NewUUID()
	if err != nil {
		return err
	}
	b.ID = id
	return nil
}
