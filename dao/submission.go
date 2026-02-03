package dao

import (
	"github.com/xieyuxuan109/homeworksystem/configs"
	"github.com/xieyuxuan109/homeworksystem/model"
)

func SearchSubByHW(id uint) (*model.Submission, error) {
	var sub model.Submission
	query := configs.DB.Model(&model.Submission{})
	var total int64
	query.Where("homework_id=?", id).Count(&total)
	if total == 0 {
		return &model.Submission{SubmissionCount: 0}, nil
	}
	result := query.Where("homework_id=?", id).First(&sub)
	if result.Error != nil {
		return nil, result.Error
	}
	return &sub, nil
}

func SearchSubmissionById(id uint) (*model.Submission, error) {
	var submission model.Submission
	query := configs.DB
	result := query.Where("id=?", id).Find(&submission)
	if result.Error != nil {
		return nil, result.Error
	}
	return &submission, nil
}
