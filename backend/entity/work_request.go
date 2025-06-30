package entity

import (
	"time"
	"gorm.io/gorm"
)

type WorkRequest struct {
	gorm.Model
	UserID      uint
	Reason      string
	StartTime   time.Time
	EndTime     time.Time
	Status      string
	Approver1ID *uint
	Approver2ID *uint
	Approver3ID *uint
}
