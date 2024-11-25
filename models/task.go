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
