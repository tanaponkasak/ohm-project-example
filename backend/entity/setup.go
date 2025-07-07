package entity

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func DB() *gorm.DB {
	return db
}

func SetupDatabase() {
	database, err := gorm.Open(sqlite.Open("sa-65.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate schema
	database.AutoMigrate(
		&User{},
		&WorkRequest{},
		&Approval{},
	)

	db = database

	// เรียกใช้ seed function
	seedUsers(db)
}

// ✅ ฟังก์ชันอยู่นอก SetupDatabase()
func seedUsers(db *gorm.DB) {

	users := []User{
		{
			Name:       "Tanapon Kasak",
			Email:      "tanapon.k@example.com",
			Position:   "Developer",
			Department: "Technology",
			Role:       "employee",
			IsActive:   true,
			WorkerID:   "W001",
		},
		{
			Name:       "Jane Doe",
			Email:      "jane.doe@example.com",
			Position:   "Project Manager",
			Department: "Business",
			Role:       "manager",
			IsActive:   true,
			WorkerID:   "W002",
		},
		{
			Name:       "John Smith",
			Email:      "john.smith@example.com",
			Position:   "UX Designer",
			Department: "Design",
			Role:       "employee",
			IsActive:   true,
			WorkerID:   "W003",
		},
		{
			Name:       "Admin User",
			Email:      "admin@example.com",
			Position:   "System Admin",
			Department: "IT",
			Role:       "admin",
			IsActive:   true,
			WorkerID:   "W004",
		},
	}

	for _, user := range users {
		// เช็คว่าอีเมลซ้ำก่อนสร้าง
		var existing User
		if err := db.Where("email = ?", user.Email).First(&existing).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				db.Create(&user)
			}
		}
	}
}
