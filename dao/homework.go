package dao

import (
	"time"

	"github.com/xieyuxuan109/homeworksystem/configs"
	"github.com/xieyuxuan109/homeworksystem/model"
)

func SearchHomework(title string) (*model.Homework, error) {
	var homework model.Homework
	query := configs.DB
	result := query.Where("title=?", title).First(&homework)
	if result.Error != nil {
		return nil, result.Error
	}
	return &homework, nil
}

func SearchHomeworkById(id uint) (*model.Homework, error) {
	var homework model.Homework
	query := configs.DB
	result := query.Where("id=?", id).Preload("Creator").First(&homework)
	if result.Error != nil {
		return nil, result.Error
	}
	return &homework, nil
}
func UpdateHomework(req model.UpdateHomework, id uint) error {
	query := configs.DB.Model(model.Homework{})
	result := query.Where("id=?", id).Updates(model.Homework{
		Title:       req.Title,
		Description: req.Description,
		Deadline:    req.Deadline,
		AllowLate:   req.AllowLate,
	})
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func DeleteHomework(id uint) error {
	query := configs.DB.Model(model.Homework{})
	result := query.Delete(&model.Homework{}, id)
	if result.Error != nil {
		return result.Error
	}
	result = configs.DB.Model(&model.Submission{}).
		Where("homework_id = ?", id).
		Update("deleted_at", time.Now())
	if result.Error != nil {
		return result.Error
	}
	return nil
}
func SearchHWDepartment(id string) (*string, error) {
	var department string
	err := configs.DB.Table("homeworks").
		Where("id = ?", id).
		Pluck("department", &department).Error
	if err != nil {
		return nil, err
	}
	return &department, nil
}
func AddHW(hw model.Homework) error {
	result := configs.DB.Create(&hw)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
