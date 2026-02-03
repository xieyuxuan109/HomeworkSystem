package router

import (
	"github.com/gin-gonic/gin"

	"github.com/xieyuxuan109/homeworksystem/handler"
	"github.com/xieyuxuan109/homeworksystem/middleware"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	//公共路由
	//用户注册
	r.POST("/user/register", handler.Register)
	//用户登录
	r.POST("/user/login", handler.Login)
	//刷新token
	r.POST("/user/refresh", handler.RefreshTokens)
	auth := r.Group("/")
	auth.Use(middleware.AuthMiddleware())
	{
		//需认证即可使用的权限
		//获取用户信息
		auth.GET("/user/profile", handler.GetProfile)
		//注销账号
		auth.DELETE("/user/account", handler.DeleteAccount)
		//获取作业列表
		r.GET("/homework", handler.GetHomeworks)
		//获取作业详情
		r.GET("/homework/:id", handler.GetHomework)
		//获取优秀作业列表
		r.GET("/submission/excellent", handler.ExcellentHomeworks)
		//老登
		//发布作业
		auth.POST("/homework", middleware.RequireRole("admin"), handler.CreateHomework)
		auth.POST("/submission/:id/aiReview", middleware.RequireRole("admin"), handler.AIcomment)
		auth.POST("/submission/:id/localaiReview", middleware.RequireRole("admin"), handler.LocalAIcomment)
		hw := auth.Group("")
		hw.Use(middleware.RequireRole("admin"), middleware.RequireSameHomeworkDepartment())
		{
			//老登+同部门
			//修改作业
			hw.PUT("/homework/:id", handler.UpdateHomework)
			// 删除作业
			hw.DELETE("/homework/:id", handler.DeleteHomework)
			//获取作业所有提交
			hw.GET("/submission/homework/:id", handler.GetSubmissions)
		}
		//标记优秀
		auth.PUT("/submission/:id/excellent", middleware.RequireRole("admin"), middleware.RequireSameHomeworkDepartmentEx(), handler.MarkExcellent)
		//批改作业
		auth.PUT("/submission/:id/review", middleware.RequireRole("admin"), middleware.RequireSameHomeworkDepartmentEx(), handler.CorrectHomework)
		//小登
		//提交作业
		auth.POST("/submission", middleware.RequireRole("student"), handler.SubmitHomework)
		//我的提交列表
		auth.GET("/submission/my", middleware.RequireRole("student"), handler.SubmitHomeworkList)
		//}
		return r
	}

}
