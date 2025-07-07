package entity

import (
    "time"
    "gorm.io/gorm"
)

type User struct {
    gorm.Model
    Name       string         `gorm:"size:100;not null"`
    Email      string         `gorm:"size:100;not null;unique"`
    Position   string         `gorm:"size:100;not null"`
    Department string         `gorm:"size:100;not null"`
    Role       string         `gorm:"type:varchar(50);default:'employee'"` // เช่น: employee, supervisor, approver
    IsActive   bool           `gorm:"default:true"`
    WorkerID   string         `gorm:"size:50"`
    DeletedAt  gorm.DeletedAt `gorm:"index"`
}

type WorkRequest struct {
    gorm.Model
    UserID    uint
    User      User           `gorm:"foreignKey:UserID"`
    Reason    string         `gorm:"type:text"`
    StartTime time.Time      // เวลาขอออก
    EndTime   time.Time      // เวลาขอเข้า
    Status    string         `gorm:"type:varchar(50);default:'pending_supervisor'"` // pending_supervisor, pending_approver, approved, rejected
    Approvals []Approval     `gorm:"foreignKey:WorkRequestID"`
    LeaveLog  *LeaveLog       `gorm:"foreignKey:WorkRequestID"`
}

type Approval struct {
    gorm.Model
    WorkRequestID uint
    WorkRequest   WorkRequest
    ApproverID    uint
    Approver      User       `gorm:"foreignKey:ApproverID"`
    Level         int        // 1 = Supervisor, 2 = Approver
    Status        string     `gorm:"type:varchar(20)"` // approved / rejected / pending
    Comment       string     `gorm:"type:text"`
    ApprovedAt    *time.Time // ใช้ pointer เพื่อบ่งบอกว่าอาจยังไม่อนุมัติ
}

type LeaveLog struct {
    gorm.Model
    WorkRequestID uint
    WorkRequest   WorkRequest
    OutTime       *time.Time // เวลาจริงที่ออก
    InTime        *time.Time // เวลาจริงที่เข้า
}
