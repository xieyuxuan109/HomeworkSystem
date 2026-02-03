package pkg

import (
	"github.com/xieyuxuan109/homeworksystem/model"
	"golang.org/x/crypto/bcrypt"
)

// SetPassword 加密并设置密码
func SetPassword(u *model.User, password string) error {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashed)
	return nil
}

// CheckPassword 验证密码
func CheckPassword(u *model.User, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password)) //顺序不能反
	return err == nil
}
