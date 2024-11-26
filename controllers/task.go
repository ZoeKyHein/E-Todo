package controllers

import (
	"E-Todo/dto"
	"E-Todo/services"
	"E-Todo/utils"
	"github.com/gin-gonic/gin"
	"strconv"
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

	// 返回成功响应
	utils.Success(c, task, "Task created successfully")
}

// FetchAllTasks 获取所有任务
func FetchAllTasks(c *gin.Context) {
	var req dto.FetchAllTasksReq

	// 绑定查询参数到 FetchAllTasksReq
	if err := c.ShouldBindQuery(&req); err != nil {
		utils.Fail(c, nil, 1001, err.Error())
		return
	}

	// 设置默认值
	if req.Page <= 0 {
		req.Page = 1
	}
	if req.Limit <= 0 {
		req.Limit = 50
	}

	tasks, total, err := services.FetchAllTasks(req)
	if err != nil {
		utils.Fail(c, nil, 1002, "Failed to fetch tasks")
		return
	}

	// 返回成功响应
	utils.Success(c, dto.FetchAllTasksResp{
		Tasks: tasks,
		Total: total,
		Page:  req.Page,
		Limit: req.Limit,
	}, "Tasks fetched successfully")
}

// UpdateTask 更新任务
func UpdateTask(c *gin.Context) {
	var req dto.UpdateTaskReq

	// 获取ID
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id <= 0 {
		utils.Fail(c, nil, 1001, "Invalid task ID")
		return
	}

	// 绑定 JSON 数据到 UpdateTaskReq
	req.ID = uint(id)
	if err = c.ShouldBindJSON(&req); err != nil {
		utils.Fail(c, nil, 1001, err.Error())
		return
	}

	task, err := services.UpdateTask(req)
	if err != nil {
		utils.Fail(c, nil, 1002, "Failed to update task")
		return
	}

	// 返回成功响应
	utils.Success(c, task, "Task updated successfully")
}
