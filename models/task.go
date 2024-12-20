package models

import (
	"E-Todo/config"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"time"
)

// Task 任务模型
type Task struct {
	ID          uint   `gorm:"primaryKey"`
	Title       string `gorm:"size:255;not null"`
	Description string
	Category    string `gorm:"size:100"`
	Color       string `gorm:"size:20"`
	DueDate     time.Time
	Status      string         `gorm:"type:enum('pending','completed');default:'pending'"`
	CreatedAt   time.Time      `gorm:"autoCreateTime"`
	UpdatedAt   time.Time      `gorm:"autoUpdateTime"`
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}

// Create 保存任务到数据库
func (t *Task) Create() error {
	return config.DB.Create(t).Error
}

// TaskQueryParams 查询参数结构体
type TaskQueryParams struct {
	Page          int
	Limit         int
	KeyWords      string
	Category      string
	Status        string
	Color         string
	RemainingDays int
}

// FetchAll 获取所有任务
func (t *Task) FetchAll(params TaskQueryParams) ([]Task, int64, error) {
	var tasks []Task
	var total int64

	query := config.DB.Model(&t)

	// 动态查询条件
	if params.KeyWords != "" {
		query = query.Where("title LIKE ?", "%"+params.KeyWords+"%")
	}
	if params.Category != "" {
		query = query.Where("category = ?", params.Category)
	}
	if params.Status != "" {
		query = query.Where("status = ?", params.Status)
	}
	if params.Color != "" {
		query = query.Where("color = ?", params.Color)
	}
	if params.RemainingDays >= 0 {
		targetDate := time.Now().AddDate(0, 0, params.RemainingDays)
		query = query.Where("due_date <= ?", targetDate)
	}

	// 分页
	query.Scopes(Paginate(params.Page, params.Limit)).Find(&tasks).Count(&total)

	return tasks, total, query.Error
}

// Update 更新任务
func (t *Task) Update() error {
	return config.DB.Model(&t).Updates(map[string]interface{}{
		"title":       t.Title,
		"description": t.Description,
		"category":    t.Category,
		"color":       t.Color,
		"due_date":    t.DueDate,
		"status":      t.Status,
	}).Error
}

// Delete 删除任务
func (t *Task) Delete() error {
	// 检查任务是否存在
	if err := t.FindTaskByID(); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("task not found: %w", err)
		}
		return fmt.Errorf("failed to query task: %w", err)
	}
	return config.DB.Unscoped().Delete(&t).Error
}

// SoftDelete 软删除任务
func (t *Task) SoftDelete() error {
	// 检查任务是否存在
	if err := t.FindTaskByID(); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("task not found: %w", err)
		}
		return fmt.Errorf("failed to query task: %w", err)
	}
	return config.DB.Delete(&t).Error
}

// Paginate 分页
func Paginate(page, limit int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if page <= 0 {
			page = 1
		}
		if limit <= 0 {
			limit = 50
		}
		offset := (page - 1) * limit
		return db.Offset(offset).Limit(limit)
	}
}

// FindTaskByID 根据 ID 查询任务
func (t *Task) FindTaskByID() error {
	return config.DB.Unscoped().First(&t, t.ID).Error
}

// Restore 恢复软删除的任务
func (t *Task) Restore() error {
	// 确保只查询软删除的记录
	if err := config.DB.Unscoped().Where("id = ? AND deleted_at IS NOT NULL", t.ID).First(&t).Error; err != nil {
		return fmt.Errorf("restore failed: task not found or not soft-deleted: %w", err)
	}

	// 恢复软删除的记录
	if err := config.DB.Unscoped().Model(&t).UpdateColumn("deleted_at", nil).Error; err != nil {
		return fmt.Errorf("restore failed: unable to update deleted_at: %w", err)
	}

	return nil
}

// Complete 完成任务
func (t *Task) Complete() error {
	// 确保只查询未完成的任务
	if err := config.DB.Where("id = ? AND status = ?", t.ID, TaskStatusPending).First(&t).Error; err != nil {
		return fmt.Errorf("complete failed: task not found or already completed: %w", err)
	}

	// 完成任务
	if err := config.DB.Model(&t).Update("status", TaskStatusCompleted).Error; err != nil {
		return fmt.Errorf("complete failed: unable to update status: %w", err)
	}

	return nil
}

// BatchDelete 批量硬删除任务
func (t *Task) BatchDelete(ids []uint) error {
	return config.DB.Unscoped().Where("id IN ?", ids).Delete(&t).Error
}

// BatchComplete 批量完成任务
func (t *Task) BatchComplete(ids []uint) error {
	return config.DB.Model(&t).Where("id IN ? AND status = ?", ids, TaskStatusPending).Update("status", TaskStatusCompleted).Error
}

// BatchSoftDelete 批量软删除任务
func (t *Task) BatchSoftDelete(ids []uint) error {
	return config.DB.Where("id IN ? AND deleted_at IS NULL", ids).Delete(&t).Error
}

// BatchRestore 批量恢复任务
func (t *Task) BatchRestore(ids []uint) error {
	return config.DB.Unscoped().Model(&t).Where("id IN ? AND deleted_at IS NOT NULL", ids).Update("deleted_at", nil).Error
}
