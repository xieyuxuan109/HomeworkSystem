package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/xieyuxuan109/homeworksystem/configs"
	"github.com/xieyuxuan109/homeworksystem/dao"
	"github.com/xieyuxuan109/homeworksystem/model"
	"github.com/xieyuxuan109/homeworksystem/pkg"
)

func RequireRole(allowedRoles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		userRole, exists := c.Get("role")
		if !exists {
			pkg.BadResponse(c, "未识别用户角色", nil)
			c.Abort()
			return
		}
		allowed := false
		for _, role := range allowedRoles {
			if userRole == role {
				allowed = true
				break
			}
		}
		if !allowed {
			pkg.BadResponse(c, "权限不足", nil)
			c.Abort()
			return
		}
		c.Next()
	}
}

// 验证老师与作业是否同部门
func RequireSameHomeworkDepartment() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 1. 获取当前老师部门
		teacherDept, exists := c.Get("department")
		if !exists {
			pkg.BadResponse(c, "教师部门信息缺失", nil)
			c.Abort()
			return
		}
		// 2. 获取作业ID
		homeworkID := c.Param("id")
		if homeworkID == "" {
			pkg.BadResponse(c, "作业ID不能为空", nil)
			c.Abort()
			return
		}
		// 3. 查询作业部门
		department, err := dao.SearchHWDepartment(homeworkID)
		if err != nil {
			pkg.BadResponse(c, "查询失败", err)
			c.Abort()
			return
		}
		if teacherDept.(string) != *department {
			pkg.BadResponse(c, "权限不足", nil)
			c.Abort()
			return
		}
		c.Next()
	}
}

func RequireSameHomeworkDepartmentEx() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 1. 获取当前老师部门
		teacherDept, exists := c.Get("department")
		if !exists {
			pkg.BadResponse(c, "教师部门信息缺失", nil)
			c.Abort()
			return
		}
		// 2. 获取提交作业ID
		submissionID := c.Param("id")
		if submissionID == "" {
			pkg.BadResponse(c, "作业ID不能为空", nil)
			c.Abort()
			return
		}
		// 3. 查询作业部门
		var submission model.Submission
		err := configs.DB.Where("id=?", submissionID).Preload("Student").Find(&submission).Error
		department := submission.Student.Department
		if err != nil {
			pkg.BadResponse(c, "查询失败", err)
			c.Abort()
			return
		}
		if teacherDept.(string) != department {
			pkg.BadResponse(c, "权限不足", nil)
			c.Abort()
			return
		}
		c.Next()
	}
}
