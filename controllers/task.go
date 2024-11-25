package controllers

import (
	"E-Todo/config"
	"E-Todo/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// CreateTaskReq 定义请求数据结构
type CreateTaskReq struct {
	Title       string `json:"title" binding:"required"`    // 任务标题，必填
	Description string `json:"description"`                 // 任务描述，选填
	Category    string `json:"category"`                    // 任务分类，选填
	Color       string `json:"color"`                       // 颜色标记，选填
	DueDate     string `json:"due_date" binding:"required"` // 截止日期，必填 (格式：yyyy-MM-ddTHH:mmZ)
}

// CreateTask 创建任务
func CreateTask(c *gin.Context) {
	var req CreateTaskReq

	// 绑定 JSON 数据到 CreateTaskReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 解析截止日期为 time.Time 类型，支持到分钟精度
	dueDate, err := time.Parse("2006-01-02T15:04Z", req.DueDate)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid due_date format. Use yyyy-MM-ddTHH:mmZ (e.g., 2025-01-06T12:30Z)"})
		return
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
	if err := task.Create(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 返回成功响应
	c.JSON(http.StatusOK, gin.H{
		"message": "Task created successfully",
		"task":    task,
	})
}

type FetchAllTasksReq struct {
	Page  int `form:"page" binding:"gte=1"`   // 页码
	Limit int `form:"limit" binding:"gte=50"` // 每页数量

	KeyWords string `form:"keywords"` // 关键字搜索
	Category string `form:"category"` // 分类搜索
	Status   string `form:"status"`   // 状态搜索
	Color    string `form:"color"`    // 颜色搜索

	RemainingDays int `form:"remaining_days" binding:"gte=0"` // 剩余天数搜索
}

// FetchAllTasks 获取所有任务
func FetchAllTasks(c *gin.Context) {
	var req FetchAllTasksReq

	// 绑定查询参数到 FetchAllTasksReq
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 初始化任务模型
	task := models.Task{}

	// 获取所有任务
	query := config.DB.Model(&task)

	// 关键字搜索
	if req.KeyWords != "" {
		query = query.Where("title LIKE ?", "%"+req.KeyWords+"%")
	}

	// 分类搜索
	if req.Category != "" {
		query = query.Where("category = ?", req.Category)
	}

	// 状态搜索
	if req.Status != "" {
		query = query.Where("status = ?", req.Status)
	}

	// 颜色搜索
	if req.Color != "" {
		query = query.Where("color = ?", req.Color)
	}

	// 剩余天数搜索
	if req.RemainingDays > 0 {
		currentDate := time.Now()
		targetDate := currentDate.AddDate(0, 0, req.RemainingDays)
		query = query.Where("due_date <= ?", targetDate)
	}

	// 分页
	page := req.Page
	limit := req.Limit
	offset := (page - 1) * limit
	var tasks []models.Task
	var total int64
	query.Count(&total).Offset(offset).Limit(limit).Find(&tasks)

	// 返回成功响应
	c.JSON(http.StatusOK, gin.H{
		"tasks": tasks,
		"total": total,
		"page":  page,
		"limit": limit,
	})
}
