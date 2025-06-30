package entity

import (
	"time"
	"gorm.io/gorm"
)

type Approval struct {
	gorm.Model
	WorkRequestID uint
	ApproverID    uint
	Level         int
	Status        string
	Comment       string
	ApprovedAt    time.Time
}