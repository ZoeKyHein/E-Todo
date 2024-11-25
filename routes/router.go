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
	}
	return r
}
