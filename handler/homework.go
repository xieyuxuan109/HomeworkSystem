package handler

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/xieyuxuan109/homeworksystem/dao"
	"github.com/xieyuxuan109/homeworksystem/model"
	"github.com/xieyuxuan109/homeworksystem/pkg"
	"github.com/xieyuxuan109/homeworksystem/service"
)

func CreateHomework(c *gin.Context) {
	var req model.PostHomework
	if err := c.ShouldBindJSON(&req); err != nil {
		pkg.BadResponse(c, "参数错误", err)
		return
	}
	id, exists := c.Get("user_id")
	if !exists {
		pkg.BadResponse(c, "请先登录", nil)
		return
	}
	Userid := id.(uint)
	res, err := service.CreateHomework(req, Userid)
	if err != nil {
		pkg.BadResponse(c, "发布失败", err)
		return
	}
	pkg.GoodResponse(c, "发布成功", res)
}

func GetHomeworks(c *gin.Context) {
	department := c.Query("department")
	pageStr := c.DefaultQuery("page", "1")
	pageSizeStr := c.DefaultQuery("page_size", "10")

	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = 1
	}

	pageSize, err := strconv.Atoi(pageSizeStr)
	if err != nil || pageSize < 1 {
		pageSize = 10
	}

	if pageSize > 100 {
		pageSize = 100
	}
	offset := (page - 1) * pageSize

	results, total, err := service.GetHomework(department, offset, page, pageSize)
	if err != nil {
		pkg.BadResponse(c, "获取失败", err)
		return
	}
	response := gin.H{
		"list":      results,
		"total":     total,
		"page":      page,
		"page_size": pageSize,
	}
	pkg.GoodResponse(c, "success", response)
}

func GetHomework(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		pkg.BadResponse(c, "无效id", err)
		return
	}
	res, err := dao.SearchHomeworkById(uint(id))
	if err != nil {
		pkg.BadResponse(c, "获取失败", err)
		return
	}
	sub, err := dao.SearchSubByHW(res.ID)
	if err != nil {
		pkg.BadResponse(c, "查询错误", err)
		return
	}
	result := gin.H{
		"id":              res.ID,
		"title":           res.Title,
		"description":     res.Description,
		"department":      res.Department,
		"deparment_label": model.GetDepartmentLabel(res.Department),
		"creator": gin.H{
			"id":       res.Creator.ID,
			"nickname": res.Creator.Nickname,
		},
		"deadline":         res.Deadline,
		"allow_late":       res.AllowLate,
		"submission_count": sub.SubmissionCount,
	}
	results := gin.H{
		"id":              res.ID,
		"title":           res.Title,
		"description":     res.Description,
		"department":      res.Department,
		"deparment_label": model.GetDepartmentLabel(res.Department),
		"creator": gin.H{
			"id":       res.Creator.ID,
			"nickname": res.Creator.Nickname,
		},
		"deadline":         res.Deadline,
		"allow_late":       res.AllowLate,
		"submission_count": sub.SubmissionCount,
		"my_submission": gin.H{
			"id":           sub.StudentID,
			"score":        sub.Score,
			"is_excellent": sub.IsExcellent,
		},
	}
	role, _ := c.Get("role")
	if role == "admin" {
		pkg.GoodResponse(c, "success", result)
	} else {
		pkg.GoodResponse(c, "success", results)
	}
}

func UpdateHomework(c *gin.Context) {
	id := c.Param("id")
	var req model.UpdateHomework
	if err := c.ShouldBindJSON(&req); err != nil {
		pkg.BadResponse(c, "参数错误", err)
	}
	UerID, _ := strconv.Atoi(id)
	res, err := service.UpdateHomework(req, uint(UerID))
	if err != nil {
		pkg.BadResponse(c, "更新失败", err)
		return
	}
	result := gin.H{
		"id":       res.ID,
		"title":    res.Title,
		"deadline": res.Deadline,
	}
	pkg.GoodResponse(c, "修改成功", result)
}

func DeleteHomework(c *gin.Context) {
	id := c.Param("id")
	UerID, _ := strconv.Atoi(id)
	dao.DeleteHomework(uint(UerID))
	pkg.GoodResponse(c, "删除成功", nil)
}
