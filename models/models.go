package models

import (
	"time"
)

// User represents a user in the system
type User struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	Username string `json:"username" gorm:"unique;not null"`
	Email    string `json:"email" gorm:"unique;not null"`
	Type     string `json:"type" gorm:"default:normal"` // Type can be "normal" or "admin"
}

// Task represents a task in the system
type Task struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	Title       string    `json:"title" gorm:"not null"`
	Description string    `json:"description"`
	Priority    string    `json:"priority" gorm:"not null"`
	DueDate     time.Time `json:"due_date" gorm:"not null"`
	Completed   bool      `json:"completed" gorm:"default:false"`
	UserID      uint      `json:"user_id" gorm:"not null"` // Foreign key to User ID
}

// Event represents an event in the system (e.g., task completion event)
type Event struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	TaskID    uint      `json:"task_id" gorm:"not null"` // Foreign key to Task ID
	EventType string    `json:"event_type" gorm:"not null"`
	Timestamp time.Time `json:"timestamp" gorm:"not null"`
}
