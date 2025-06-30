package entity

import (
    "time"
    "gorm.io/gorm"
)

type User struct {
    gorm.Model
    Name       string    `gorm:"size:100;not null"`
    Email      string    `gorm:"size:100;not null;unique"`
    Position   string    `gorm:"size:100;not null"`
    Department string    `gorm:"size:100;not null"`
    Role       string    `gorm:"type:text;default:'employee'"` // เปลี่ยนจาก enum เป็น string
    IsActive   bool      `gorm:"default:true"`
    CreatedAt  time.Time
    UpdatedAt  time.Time
    DeletedAt  gorm.DeletedAt `gorm:"index"`
}
