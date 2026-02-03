package model

import (
	"time"

	"gorm.io/gorm"
)

// 用户表
type User struct {
	ID         uint   `json:"id"`
	Username   string `json:"username"`
	Password   string `json:"-"`
	Nickname   string `json:"nickname"`
	Email      string `json:"email,omitempty"`
	Role       string `gorm:"default:'student'" json:"role"`
	Department string `json:"department"`

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}

// 用户注册请求
type UserRegisterRequest struct {
	Username   string `json:"username" binding:"required,min=3,max=50"`
	Password   string `json:"password" binding:"required,min=6,max=100"`
	Nickname   string `json:"nickname" binding:"required,min=2,max=50"`
	Department string `json:"department" binding:"required,oneof=backend frontend sre product design android ios"`
	Role       string `json:"role"`
}

// 用户登录请求
type UserLoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// 用户注销账户结构体
type UserDeleteRequest struct {
	Password string `json:"password" binding:"required"`
}

// 用户刷新token结构体
type RefreshRequest struct {
	RefreshToken string `json:"refresh_token" binding:"required"`
}

// 用户响应
type UserResponse struct {
	ID              uint   `json:"id"`
	Username        string `json:"username"`
	Nickname        string `json:"nickname"`
	Email           string `json:"email"`
	Role            string `json:"role"`
	Department      string `json:"department"`
	DepartmentLabel string `json:"department_label"`
}

// 获取部门中文标签
func GetDepartmentLabel(Department string) string {
	labels := map[string]string{
		"backend":  "后端",
		"frontend": "前端",
		"sre":      "SRE",
		"product":  "产品",
		"design":   "视觉设计",
		"android":  "Android",
		"ios":      "iOS",
	}
	return labels[Department]
}

// 转换为响应格式
func (u *User) ToResponse() *UserResponse {
	return &UserResponse{
		ID:              u.ID,
		Username:        u.Username,
		Nickname:        u.Nickname,
		Email:           u.Email,
		Role:            u.Role,
		Department:      u.Department,
		DepartmentLabel: GetDepartmentLabel(u.Department),
	}
}
