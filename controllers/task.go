package controllers

import (
	"E-Todo/dto"
	"E-Todo/services"
	"E-Todo/utils"
	"github.com/gin-gonic/gin"
)

// CreateTask 创建任务
func CreateTask(c *gin.Context) {
	var req dto.CreateTaskReq

	// 绑定 JSON 数据到 CreateTaskReq
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.Fail(c, nil, 1001, err.Error())
		return
	}

	task, err := services.CreateTask(req)
	if err != nil {
		utils.Fail(c, nil, 1002, "Failed to create task")
		return
	}

	// 构造响应数据
	resp := dto.CreateTaskResp{
		ID:          task.ID,
		Title:       task.Title,
		Description: task.Description,
		Category:    task.Category,
		Color:       task.Color,
		DueDate:     task.DueDate.Format("2006-01-02T15:04Z"),
		Status:      task.Status,
	}

	// 返回成功响应
	utils.Success(c, resp, "Task created successfully")
}

// FetchAllTasks 获取所有任务
func FetchAllTasks(c *gin.Context) {
	var req dto.FetchAllTasksReq

	// 绑定查询参数到 FetchAllTasksReq
	if err := c.ShouldBindQuery(&req); err != nil {
		utils.Fail(c, nil, 1001, err.Error())
		return
	}

	tasks, total, err := services.FetchAllTasks(req)
	if err != nil {
		utils.Fail(c, nil, 1002, "Failed to fetch tasks")
	}

	// 返回成功响应
	utils.Success(c, dto.FetchAllTasksResp{
		Tasks: tasks,
		Total: total,
		Page:  req.Page,
		Limit: req.Limit,
	}, "Tasks fetched successfully")
}
