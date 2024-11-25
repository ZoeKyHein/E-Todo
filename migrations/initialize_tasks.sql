CREATE TABLE tasks (
                       id INT AUTO_INCREMENT PRIMARY KEY,        -- 任务唯一 ID
                       title VARCHAR(255) NOT NULL,             -- 任务标题
                       description TEXT,                        -- 任务描述
                       category VARCHAR(100),                   -- 任务分类
                       color VARCHAR(20),                       -- 颜色标记（如 #FF0000）
                       due_date DATETIME,                       -- 到期时间
                       status ENUM('pending', 'completed') DEFAULT 'pending', -- 任务状态
                       created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,        -- 创建时间
                       updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP -- 更新时间
);