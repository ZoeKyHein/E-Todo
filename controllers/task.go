package controllers

import (
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
