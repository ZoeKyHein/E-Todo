package services

import (
	"E-Todo/dto"
	"E-Todo/models"
	"fmt"
	"time"
)

// CreateTask 创建任务
func CreateTask(req dto.CreateTaskReq) (dto.TaskDTO, error) {
	// 解析截止日期
	dueDate, err := time.Parse("2006-01-02T15:04Z", req.DueDate)
	if err != nil {
		return dto.TaskDTO{}, fmt.Errorf("invalid due date format: %w", err)
	}

	// 初始化任务模型
	task := models.Task{
		Title:       req.Title,
		Description: req.Description,
		Category:    req.Category,
		Color:       req.Color,
		DueDate:     dueDate,
	}

	// 保存到数据库
	if err = task.Create(); err != nil {
		return dto.TaskDTO{}, fmt.Errorf("failed to create task: %w", err)
	}

	// 构造 TaskDTO
	return dto.TaskDTO{
		ID:          task.ID,
		Title:       task.Title,
		Description: task.Description,
		Category:    task.Category,
		Color:       task.Color,
		DueDate:     task.DueDate.Format("2006-01-02T15:04Z"),
		Status:      task.Status,
		CreatedAt:   task.CreatedAt.Format("2006-01-02T15:04:05Z"),
		UpdatedAt:   task.UpdatedAt.Format("2006-01-02T15:04:05Z"),
	}, nil
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
	tasks, total, err := task.FetchAll(params)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to fetch tasks: %w", err)
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
