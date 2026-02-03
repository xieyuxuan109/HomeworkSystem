package configs

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error
var JWT_SECRET string = "zXk9PqR2sV5wY8tL3mN6bC1dF4gH7jK0pO2iU5yA8tQ3wE6r"

func init() { //init方法
	dsn := "root:root@tcp(localhost:3306)/homework?charset=utf8mb4&parseTime=True&loc=Local"
	//格式：数据库用户名:密码@tcp(数据库所在电脑ip地址:3306)/xieyuxuan?charset=utf8mb4&parseTime=True&loc=Local
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err) // 或者 log.Fatal(err)
	}
}
