package services

import (
	"E-Todo/dto"
	"E-Todo/models"
	"time"
)

func CreateTask(req dto.CreateTaskReq) (models.Task, error) {
	// 解析截止日期为 time.Time 类型，支持到分钟精度
	dueDate, err := time.Parse("2006-01-02T15:04Z", req.DueDate)
	if err != nil {
		return models.Task{}, err
	}

	// 初始化任务模型
	task := models.Task{
		Title:       req.Title,
		Description: req.Description,
		Category:    req.Category,
		Color:       req.Color,
		DueDate:     dueDate,
	}

	// 保存任务到数据库
	err = task.Create()

	return task, err
}

// FetchAllTasks 获取所有任务
func FetchAllTasks(req dto.FetchAllTasksReq) ([]dto.TaskDTO, int64, error) {

	params := models.TaskQueryParams{
		Page:          req.Page,
		Limit:         req.Limit,
		KeyWords:      req.KeyWords,
		Category:      req.Category,
		Status:        req.Status,
		Color:         req.Color,
		RemainingDays: req.RemainingDays,
	}

	var task models.Task

	// 获取所有任务
	tasks, total, err := task.FetchAll(params)
	if err != nil {
		return nil, 0, err
	}

	var taskDTOs []dto.TaskDTO
	for _, t := range tasks {
		taskDTOs = append(taskDTOs, dto.TaskDTO{
			ID:          t.ID,
			Title:       t.Title,
			Description: t.Description,
			Category:    t.Category,
			Color:       t.Color,
			DueDate:     t.DueDate.Format("2006-01-02T15:04Z"),
			Status:      t.Status,
		})
	}

	return taskDTOs, total, nil
}
