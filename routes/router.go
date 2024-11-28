package routes

import (
	"E-Todo/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	tasks := r.Group("tasks")
	{
		tasks.POST("", controllers.CreateTask)
		tasks.GET("", controllers.FetchAllTasks)
		tasks.PUT("/:id", controllers.UpdateTask)
		tasks.DELETE("/:id", controllers.DeleteTask)
		tasks.PATCH("/:id", controllers.SoftDelete)
		tasks.PATCH("/:id/restore", controllers.RestoreTask)
		tasks.PATCH("/:id/complete", controllers.CompleteTask)

		batchTasks := tasks.Group("batch")
		{
			batchTasks.DELETE("", controllers.BatchDeleteTasks)
			batchTasks.PATCH("", controllers.BatchSoftDeleteTasks)
			batchTasks.PATCH("complete", controllers.BatchCompleteTasks)
			batchTasks.PATCH("restore", controllers.BatchRestoreTasks)
		}
	}
	return r
}
