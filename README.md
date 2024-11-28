# E-Todo 任务管理系统

E-Todo 是一个基于 Go 和 Gin 框架开发的任务管理系统，支持任务的增删改查、批量操作、颜色标记、分类、逾期提醒等功能。

E-Todo is a task management system built with Go and the Gin framework, supporting task CRUD operations, batch operations, color tagging, categorization, overdue reminders, and more.

## 功能特点 / Features

- 任务的增删改查 / CRUD operations for tasks
- 支持批量操作（批量删除、批量完成、批量恢复等） / Batch operations (delete, complete, restore, etc.)
- 颜色标记和分类管理 / Color tagging and categorization
- 软删除与恢复功能 / Soft delete and restore functionality

## 项目结构 / Project Structure

```
E-Todo/
├── controllers/  # 控制器层，负责处理 HTTP 请求 / Controllers for handling HTTP requests
├── services/     # 服务层，包含业务逻辑 / Services with business logic
├── models/       # 数据层，定义数据库模型 / Data layer with database models
├── dto/          # 数据传输对象，定义请求和响应格式 / Data Transfer Objects for request/response formats
├── utils/        # 工具函数和通用方法 / Utility functions and common methods
├── config/       # 配置文件和数据库初始化 / Configuration and database initialization
├── routes/       # 路由定义 / Route definitions
├── main.go       # 主程序入口 / Main program entry point
```

## 技术栈 / Tech Stack

- 编程语言 / Programming Language: Go
- Web 框架 / Web Framework: Gin
- 数据库 / Database: MySQL
- ORM 工具 / ORM Tool: GORM

## 安装与运行 / Installation and Running

### 环境要求 / Requirements

- Go 1.18+
- MySQL

### 本地运行 / Running Locally

1. 克隆项目 / Clone the repository:
   ```bash
   git clone https://github.com/your-repo/E-Todo.git
   cd E-Todo
   ```

2. 配置环境变量 / Set up environment variables:
   创建一个 `.env` 文件，并根据需要填写数据库连接信息。/ Create a `.env` file and set up your database connection information.

3. 安装依赖 / Install dependencies:
   ```bash
   go mod tidy
   ```

4. 运行项目 / Run the project:
   ```bash
   go run main.go
   ```

5. 打开浏览器访问 / Open your browser and visit:
   ```
    http://localhost:8080
   ```

## API 文档 / API Documentation

API 文档使用 [APIFOX](https://apifox.com/apidoc/shared-f53cc96e-57d6-4bf1-9341-795011f7d9b7/237933623e0) 生成。/ API documentation is generated with [APIFOX](https://apifox.com/apidoc/shared-f53cc96e-57d6-4bf1-9341-795011f7d9b7/237933623e0).


## 许可证 / License

MIT License. Feel free to use and modify.
