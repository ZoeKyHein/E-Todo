package dto

// CreateTaskReq 定义请求数据结构
type CreateTaskReq struct {
	Title       string `json:"title" binding:"required"`    // 任务标题，必填
	Description string `json:"description"`                 // 任务描述，选填
	Category    string `json:"category"`                    // 任务分类，选填
	Color       string `json:"color"`                       // 颜色标记，选填
	DueDate     string `json:"due_date" binding:"required"` // 截止日期，必填 (格式：yyyy-MM-ddTHH:mmZ)
}

// FetchAllTasksReq 获取所有任务请求参数
type FetchAllTasksReq struct {
	Page          int    `form:"page"`                           // 页码
	Limit         int    `form:"limit"`                          // 每页数量
	KeyWords      string `form:"keywords"`                       // 关键字搜索
	Category      string `form:"category"`                       // 分类搜索
	Status        string `form:"status"`                         // 状态搜索
	Color         string `form:"color"`                          // 颜色搜索
	RemainingDays int    `form:"remaining_days" binding:"gte=0"` // 剩余天数搜索
}

// FetchAllTasksResp 获取所有任务响应参数
type FetchAllTasksResp struct {
	Tasks []TaskDTO `json:"tasks"`
	Total int64     `json:"total"`
	Page  int       `json:"page"`
	Limit int       `json:"limit"`
}

// TaskDTO 任务数据传输对象
type TaskDTO struct {
	ID          uint   `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Category    string `json:"category"`
	Color       string `json:"color"`
	DueDate     string `json:"due_date"`
	Status      string `json:"status"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

// UpdateTaskReq 更新任务请求参数
type UpdateTaskReq struct {
	ID          uint   `json:"id" binding:"required"` // 任务 ID，必填
	Title       string `json:"title"`                 // 任务标题，选填
	Description string `json:"description"`           // 任务描述，选填
	Category    string `json:"category"`              // 任务分类，选填
	Color       string `json:"color"`                 // 颜色标记，选填
	DueDate     string `json:"due_date"`              // 截止日期，选填 (格式：yyyy-MM-ddTHH:mmZ)
	Status      string `json:"status"`                // 任务状态，选填
}

// BatchTaskActionReq 批量任务操作请求参数
type BatchTaskActionReq struct {
	IDs []uint `json:"ids" binding:"required"`
}
