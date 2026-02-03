package model

import (
	"time"

	"gorm.io/gorm"
)

// 提交作业模型结构体
type Submission struct {
	ID              uint       `json:"id"`
	HomeworkID      uint       `json:"homework_id"`
	StudentID       uint       `json:"student_id"`
	Content         string     `json:"content"`
	FileURL         string     `json:"file_url"`
	IsLate          bool       `json:"is_late"`
	Score           int        `json:"score"`
	Comment         string     `json:"comment"`
	IsExcellent     bool       `json:"is_excellent"`
	ReviewerID      uint       `json:"reviewer_id"`
	ReviewedAt      *time.Time `json:"reviewed_at"`
	SubmissionCount int        `json:"submission_count"`
	CreatedAt       time.Time  `json:"created_at"`
	UpdatedAt       time.Time  `json:"updated_at"`

	Homework  Homework       `gorm:"foreignKey:HomeworkID;references:ID" json:"homework"`
	Student   User           `gorm:"foreignKey:StudentID;references:ID" json:"student"`
	Reviewer  User           `gorm:"foreignKey:ReviewerID;references:ID" json:"reviewer"`
	DeletedAt gorm.DeletedAt `json:"-"`
}

// 提交作业请求结构体
type SubmissionRequest struct {
	HomeworkID uint   `json:"homework_id"`
	Content    string `json:"content"`
	FileURL    string `json:"file_url"`
}

// 提交作业响应结构体
type SubmissionResponse struct {
	ID          uint      `json:"id"`
	HomeworkID  uint      `json:"homework_id"`
	IsLate      bool      `json:"is_late"`
	SubmittedAt time.Time `json:"submitted_at"`
}

// 标记优秀结构体
type Excellent struct {
	IsExcellent bool `json:"is_excellent"`
}

// 批改作业请求结构体
type CorrectHomework struct {
	Score       int    `json:"score"`
	Comment     string `json:"comment"`
	IsExcellent bool   `json:"is_excellent"`
}

// AI
type AIcommentRequest struct {
	HomeworkID   uint `json:"homework_id"`
	SubmissionID uint `json:"submission_id"`
}
