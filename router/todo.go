package router

import (
	"github.com/gin-gonic/gin"
	"todolist/service"
)

var todoService service.TodoService
var listService service.ListService

func Router(r *gin.Engine) {
	todoService = service.TodoServiceImpl{}
	listService = service.ListServiceImpl{}

	todoRouter := r.Group("/todo")
	{
		todoRouter.GET("/list", todoService.List)
		todoRouter.POST("/", todoService.Create)
		todoRouter.DELETE("/", todoService.Delete)
		todoRouter.PUT("/", todoService.Update)
	}

	listRouter := r.Group("/list")
	{
		listRouter.GET("/all", listService.All)
		listRouter.POST("/page", listService.Page)
		listRouter.POST("/", listService.Create)
		listRouter.DELETE("/", listService.Delete)
		listRouter.PUT("/", listService.Update)
	}
}
