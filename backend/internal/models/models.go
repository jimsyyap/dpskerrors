package models

import (
	"time"
	"github.com/google/uuid"
)

type User struct {
	ID           uuid.UUID `gorm:"type:uuid;primaryKey"`
	Username     string    `gorm:"unique"`
	PasswordHash string
}

type MatchSession struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey"`
	UserID    uuid.UUID 
	StartTime time.Time 
	EndTime   *time.Time 
}

type ErrorType struct {
	ID   int    `gorm:"primaryKey"`
	Name string 
}

type ErrorLog struct {
	ID           uuid.UUID `gorm:"type:uuid;primaryKey"`
	SessionID    uuid.UUID 
	ErrorTypeID  int       
	Timestamp    time.Time 
}
