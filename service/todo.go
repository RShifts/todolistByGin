package service

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
	"todolist/db"
)

type TodoService interface {
	Create(c *gin.Context)
	Delete(c *gin.Context)
	Update(c *gin.Context)
	List(c *gin.Context)
}

type TodoServiceImpl struct{}

func (todoService TodoServiceImpl) Create(c *gin.Context) {
	var todo db.Todo
	err := c.ShouldBindJSON(&todo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}

	todo.Id = db.UUID()
	todo.Deleted = 0
	todo.CreateTime = time.Now()
	todo.LastUpdateTime = time.Now()

	db.DB.Create(&todo)
	c.JSON(http.StatusOK, todo)
}

func (todoService TodoServiceImpl) List(c *gin.Context) {}

func (todoService TodoServiceImpl) Delete(c *gin.Context) {}

func (todoService TodoServiceImpl) Update(c *gin.Context) {
	var todo db.Todo
	err := c.ShouldBindJSON(&todo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}
	var t db.Todo
	db.DB.Where("id = ?", todo.Id).First(&t)

	t.Content = todo.Content
	t.Done = todo.Done
	t.Sort = todo.Sort
	t.LastUpdateTime = time.Now()

	db.DB.Save(&t)

	c.JSON(http.StatusOK, t)
}
