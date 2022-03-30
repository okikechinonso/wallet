package entities

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Model struct {
	ID        string    `sql:"type:uuid; default:uuid_generate_v4();size:100; not null"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

func (u *Model) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New().String()
	if u.ID == "" {
		err = errors.New("can't save invalid data")
	}
	return
}

type Wallet struct {
	Model
	Address string
	Balance float64
}
