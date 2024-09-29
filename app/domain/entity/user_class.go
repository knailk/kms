package entity

import (
	"time"

	"github.com/google/uuid"
)

type UserClass struct {
	ID        uuid.UUID `gorm:"primaryKey"`
	Username  string    `gorm:"index:,unique,composite:username_class_id_unique"`
	ClassID   uuid.UUID `gorm:"index:,unique,composite:username_class_id_unique"`
	Status    string    // joined, studying, complete, ...etc
	CreatedAt time.Time
	UpdatedAt time.Time

	User        *User         `gorm:"foreignKey:Username;references:Username"`
	CheckInOuts []*CheckInOut `gorm:"foreignKey:UserClassID"`
}

type UserClassStatus string

const (
	UserClassStatusJoined   UserClassStatus = "joined"
	UserClassStatusStudying UserClassStatus = "studying"
	UserClassStatusComplete UserClassStatus = "complete"
	UserClassStatusCanceled UserClassStatus = "canceled"
)
