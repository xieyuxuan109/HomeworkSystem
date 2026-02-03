package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/xieyuxuan109/homeworksystem/dao"
	"github.com/xieyuxuan109/homeworksystem/model"
	"github.com/xieyuxuan109/homeworksystem/pkg"

	"github.com/xieyuxuan109/homeworksystem/service"
)

func Register(c *gin.Context) {
	var req model.UserRegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		pkg.BadResponse(c, "参数错误", err)
		return
	}
	res, err := service.Register(req)
	if err != nil {
		pkg.BadResponse(c, "注册失败", err)
		return
	}
	pkg.GoodResponse(c, "注册成功", res)

}

func Login(c *gin.Context) {
	var req model.UserLoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		pkg.BadResponse(c, "参数错误", err)
		return
	}
	response, err := service.Login(req)
	if err != nil {
		pkg.BadResponse(c, "登录失败", err)
		return
	}
	result, err := pkg.GenerateTokens(response.ID, response.Username, response.Role, response.Department)
	if err != nil {
		pkg.BadResponse(c, "token生成失败", err)
		return
	}
	pkg.GoodResponse(c, "登录成功", gin.H{
		"access_token":  result["access_token"],
		"refresh_token": result["refresh_token"],
		"user":          response,
	})
}

func RefreshTokens(c *gin.Context) {

	var req model.RefreshRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		pkg.BadResponse(c, "未传递refresh_token", err)
		return
	}
	tokens, err := service.RefreshTokens(req)
	if err != nil {
		pkg.BadResponse(c, "token刷新失败", err)
		return
	}
	pkg.GoodResponse(c, "token刷新成功", tokens)
}

func GetProfile(c *gin.Context) {
	username, exists := c.Get("username")
	if !exists {
		pkg.BadResponse(c, "请先登录",nil)
		return
	}
	user, err := dao.Search(username.(string))
	if err != nil {
		pkg.BadResponse(c, "用户不存在",err)
		return
	}
	userResponse := user.ToResponse()
	pkg.GoodResponse(c, "success", userResponse)
}

func DeleteAccount(c *gin.Context) {
	var req model.UserDeleteRequest
	if err := c.ShouldBindBodyWithJSON(&req); err != nil {
		pkg.BadResponse(c, "参数错误", err)
		return
	}
	username, exists := c.Get("username")
	if !exists {
		pkg.BadResponse(c, "请先登录",nil)
		return
	}
	err := service.DeleteAccount(req, username.(string))
	if err != nil {
		pkg.BadResponse(c, "注销失败", err)
		return
	}
	pkg.GoodResponse(c, "账号已注销", nil)
}
