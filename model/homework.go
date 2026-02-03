package model

import (
	"time"

	"gorm.io/gorm"
)

// 学生作业结构体
type Homework struct {
	ID          uint           `json:"id"`
	Title       string         `json:"title"`
	Description string         `json:"description"`
	Department  string         `json:"department"`
	CreatorID   uint           `json:"creator_id"`
	Deadline    time.Time      `json:"deadline"`
	AllowLate   bool           `json:"allow_late"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"-" gorm:"index"`
	Creator     User           `gorm:"foreignKey:CreatorID;references:ID"`
	Submissions []Submission   `gorm:"foreignKey:HomeworkID"`
}

// 发布作业请求结构体
type PostHomework struct {
	Title       string    `json:"title" binding:"required"`
	Description string    `json:"description"`
	Department  string    `json:"department" binding:"required,oneof=backend frontend sre product design android ios"`
	Deadline    time.Time `json:"deadline" binding:"required"`
	AllowLate   bool      `json:"allow_late" binding:"required"`
}

// 发布作业响应结构体
type HomeworkResponse struct {
	ID              uint      `json:"id"`
	Title           string    `json:"title"`
	Department      string    `json:"department" binding:"required,oneof=backend frontend sre product design android ios"`
	DepartmentLabel string    `json:"department_label"`
	Deadline        time.Time `json:"deadline" binding:"required"`
	AllowLate       bool      `json:"allow_late" binding:"required"`
}

// 修改作业请求结构体
type UpdateHomework struct {
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Deadline    time.Time `json:"deadline"`
	AllowLate   bool      `json:"allow_late"`
}

// 转换为响应格式
func (u *Homework) ToResponse() *HomeworkResponse {
	return &HomeworkResponse{
		ID:              u.ID,
		Title:           u.Title,
		Department:      u.Department,
		DepartmentLabel: GetDepartmentLabel(u.Department),
		Deadline:        u.Deadline,
		AllowLate:       u.AllowLate,
	}
}
