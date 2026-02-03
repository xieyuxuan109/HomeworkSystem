package service

import (
	"errors"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/xieyuxuan109/homeworksystem/configs"
	"github.com/xieyuxuan109/homeworksystem/model"
)

var Lock sync.Mutex

func SubmitHomework(req model.SubmissionRequest, Department string, ID uint) (res *model.SubmissionResponse, err error) {
	var homework model.Homework
	var submission model.Submission
	var submissionExist model.Submission
	result := configs.DB.First(&homework, req.HomeworkID)
	if result.Error != nil {
		return nil, result.Error
	}
	if homework.Department == Department {
		isLate := time.Now().After(homework.Deadline)
		submission.IsLate = isLate
		if isLate {
			if homework.AllowLate {
				submission.StudentID = ID
				submission.HomeworkID = req.HomeworkID
				submission.Content = req.Content
				submission.FileURL = req.FileURL
			} else {
				return nil, errors.New("作业截止时间已过且不允许补交")
			}
		} else {
			submission.StudentID = ID
			submission.HomeworkID = req.HomeworkID
			submission.Content = req.Content
			submission.FileURL = req.FileURL
		}
	} else {
		return nil, errors.New("该作业不是所在部门的作业")
	}
	result = configs.DB.Create(&submission)
	if result.Error != nil {
		return nil, result.Error
	}
	result = configs.DB.Where("homework_id=? AND student_id=?", req.HomeworkID, ID).First(&submissionExist)
	if result.Error != nil {
		return nil, result.Error
	}
	res = &model.SubmissionResponse{
		ID:          submissionExist.ID,
		HomeworkID:  submissionExist.HomeworkID,
		IsLate:      submissionExist.IsLate,
		SubmittedAt: submissionExist.UpdatedAt,
	}
	return res, nil
}

func SubmitHomeworkList(ID uint, page int, offset int) (res []map[string]interface{}, total int64, err error) {
	var submissions []model.Submission
	res = make([]map[string]interface{}, 0)
	configs.DB.Model(&model.Submission{}).Where("student_id=?", ID).Count(&total)
	result := configs.DB.Where("student_id=?", ID).Preload("Homework").Offset(offset).Limit(page).Find(&submissions)

	if result.Error != nil {
		return nil, 0, result.Error
	}
	for _, v := range submissions {
		res = append(res, map[string]interface{}{
			"id": v.ID,
			"homework": gin.H{
				"id":               v.Homework.ID,
				"title":            v.Homework.Title,
				"department":       v.Homework.Department,
				"department_label": model.GetDepartmentLabel(v.Homework.Department),
			},
			"score":        v.Score,
			"comment":      v.Comment,
			"is_excellent": v.IsExcellent,
			"submitted_at": v.UpdatedAt,
		})
	}
	return res, total, nil
}

func MarkExcellent(req model.Excellent, id uint) (res gin.H, err error) {
	result := configs.DB.Model(&model.Submission{}).Where("id=?", id).Updates(model.Submission{
		IsExcellent: req.IsExcellent,
	})
	if result.Error != nil {
		return nil, result.Error
	}
	var responese model.Submission
	result = configs.DB.Where("id = ?", id).Find(&responese)
	if result.Error != nil {
		return nil, result.Error
	}
	res = gin.H{
		"id":           responese.ID,
		"is_excellent": responese.IsExcellent,
	}
	return res, nil
}
func CorrectHomework(req model.CorrectHomework, id uint) (res gin.H, err error) {
	Lock.Lock()
	defer Lock.Unlock()
	now := time.Now()
	result := configs.DB.Model(&model.Submission{}).Where("id=?", id).Updates(model.Submission{
		IsExcellent: req.IsExcellent,
		Score:       req.Score,
		Comment:     req.Comment,
		ReviewedAt:  &now,
	})
	if result.Error != nil {
		return nil, result.Error
	}
	var responese model.Submission
	result = configs.DB.Model(&model.Submission{}).Where("id = ?", id).First(&responese)
	if result.Error != nil {
		return nil, result.Error
	}
	res = gin.H{
		"id":           responese.ID,
		"comment":      responese.Comment,
		"reviewed_at":  responese.ReviewedAt,
		"is_excellent": responese.IsExcellent,
	}
	return res, nil
}

func ExcellentHomeworks(department string, page int, pageSize int) (res []map[string]interface{}, total int64, err error) {
	var submissions []model.Submission
	res = make([]map[string]interface{}, 0)
	offset := (page - 1) * pageSize
	// 先查询该院系的学生ID
	var studentIDs []uint
	configs.DB.Model(&model.User{}).
		Select("id").
		Where("department = ?", department).
		Find(&studentIDs)

	// 然后查询这些学生的优秀提交
	query := configs.DB.Model(&model.Submission{}).
		Where("is_excellent = ? AND student_id IN (?)", true, studentIDs)

	query.Count(&total)
	result := query.Preload("Homework").Preload("Student").Offset(offset).Limit(page).Find(&submissions)

	if result.Error != nil {
		return nil, 0, result.Error
	}
	for _, v := range submissions {
		res = append(res, map[string]interface{}{
			"id": v.ID,
			"homework": gin.H{
				"id":               v.Homework.ID,
				"title":            v.Homework.Title,
				"department":       v.Homework.Department,
				"department_label": model.GetDepartmentLabel(v.Homework.Department),
			},
			"student": gin.H{
				"id":       v.Student.ID,
				"nickname": v.Student.Nickname,
			},
			"score":   v.Score,
			"comment": v.Comment,
		})
	}
	return res, total, nil
}

func GetSubmissions(id string, offset int, page int, pageSize int) (gin.H, error) {
	var homework model.Homework
	var total int
	err := configs.DB.First(&homework, id).Error
	if err != nil {
		return nil, errors.New("该作业不存在")
	}
	var submissions []model.Submission
	err = configs.DB.Where("homework_id=?", id).Preload("Student").Order("created_at DESC").Offset(offset).Limit(pageSize).Find(&submissions).Error
	total = len(submissions)
	if err != nil {
		return nil, err
	}
	results := make([]map[string]interface{}, len(submissions))
	for i, v := range submissions {
		results[i] = gin.H{
			"id": v.ID,
			"student": gin.H{
				"id":               v.StudentID,
				"nickname":         v.Student.Nickname,
				"department":       v.Student.Department,
				"department_label": model.GetDepartmentLabel(v.Student.Department),
			},
			"content":      v.Content,
			"is_late":      v.IsLate,
			"score":        v.IsLate,
			"comment":      v.Score,
			"submitted_at": v.UpdatedAt,
		}
	}
	response := gin.H{
		"list":      results,
		"total":     total,
		"page":      page,
		"page_size": pageSize,
	}
	return response, nil
}
