package models

import (
	"E-Todo/config"
	"time"
)

type Task struct {
	ID          uint      `gorm:"primaryKey"`        // 任务 ID
	Title       string    `gorm:"size:255;not null"` // 任务标题
	Description string    // 任务描述
	Category    string    `gorm:"size:100"`                                           // 任务分类
	Color       string    `gorm:"size:20"`                                            // 任务颜色(十六进制 如#ffffff)
	DueDate     time.Time `gorm:"type:datetime"`                                      // 截止日期
	Status      string    `gorm:"type:enum('pending','completed');default:'pending'"` // 任务状态
	CreatedAt   time.Time `gorm:"autoCreatedTime"`                                    // 创建时间
	UpdatedAt   time.Time `gorm:"autoUpdateTime"`                                     // 更新时间
}

// Create 创建任务
func (t *Task) Create() error {
	return config.DB.Create(t).Error
}

// TaskQueryParams 任务查询参数(Database)
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

	// 关键字搜索
	if params.KeyWords != "" {
		query = query.Where("title LIKE ?", "%"+params.KeyWords+"%")
	}

	// 分类搜索
	if params.Category != "" {
		query = query.Where("category = ?", params.Category)
	}

	// 状态搜索
	if params.Status != "" {
		query = query.Where("status = ?", params.Status)
	}

	// 颜色搜索
	if params.Color != "" {
		query = query.Where("color = ?", params.Color)
	}

	// 剩余天数搜索
	if params.RemainingDays > 0 {
		currentDate := time.Now()
		targetDate := currentDate.AddDate(0, 0, params.RemainingDays)
		query = query.Where("due_date <= ?", targetDate)
	}

	// 获取总数
	query.Count(&total)

	// 分页
	if params.Page != 0 || params.Limit != 0 {
		if params.Page < 0 {
			params.Page = 1
		}
		if params.Limit < 0 {
			params.Limit = 50
		}
		offset := (params.Page - 1) * params.Limit
		limit := params.Limit
		query = query.Offset(offset).Limit(limit)
	}

	err := query.Find(&tasks).Error
	return tasks, total, err
}
