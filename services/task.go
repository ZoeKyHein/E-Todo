package services

import (
	"E-Todo/config"
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

// UpdateTask 更新任务
func UpdateTask(req dto.UpdateTaskReq) (dto.TaskDTO, error) {
	// 初始化任务模型
	var task models.Task

	// 查询任务
	if err := config.DB.Where("id = ?", req.ID).First(&task).Error; err != nil {
		return dto.TaskDTO{}, fmt.Errorf("failed to find task: %w", err)
	}

	// 更新字段
	if req.Title != "" {
		task.Title = req.Title
	}
	if req.Description != "" {
		task.Description = req.Description
	}
	if req.Category != "" {
		task.Category = req.Category
	}
	if req.Color != "" {
		task.Color = req.Color
	}
	if req.DueDate != "" {
		dueDate, err := time.Parse("2006-01-02T15:04Z", req.DueDate)
		if err != nil {
			return dto.TaskDTO{}, fmt.Errorf("invalid due date format: %w", err)
		}
		task.DueDate = dueDate
	}
	if req.Status != "" {
		task.Status = req.Status
	}

	// 更新任务
	if err := task.Update(); err != nil {
		return dto.TaskDTO{}, fmt.Errorf("failed to update task: %w", err)
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

// DeleteTask 删除任务
func DeleteTask(id uint) error {
	// 初始化任务模型
	var task models.Task

	task.ID = id

	// 删除任务
	if err := task.Delete(); err != nil {
		return fmt.Errorf("failed to hard delete task with ID %d: %w", id, err)
	}

	return nil
}

// SoftDelete 软删除任务
func SoftDelete(id uint) error {
	// 初始化任务模型
	var task models.Task

	task.ID = id

	// 软删除任务
	if err := task.SoftDelete(); err != nil {
		return fmt.Errorf("failed to soft delete task with ID %d: %w", id, err)
	}

	return nil
}

func RestoreTask(id uint) error {
	// 初始化任务模型
	var task models.Task

	task.ID = id

	// 恢复任务
	if err := task.Restore(); err != nil {
		return fmt.Errorf("service: failed to restore task with ID %d: %w", id, err)
	}

	return nil
}
