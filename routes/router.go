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
	}
	return r
}
