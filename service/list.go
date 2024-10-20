package service

import "github.com/gin-gonic/gin"

type ListService interface {
	All(c *gin.Context)
	Page(c *gin.Context)
	Create(c *gin.Context)
	Delete(c *gin.Context)
	Update(c *gin.Context)
}

type ListServiceImpl struct {
}

func (listService ListServiceImpl) All(c *gin.Context) {

}

func (listService ListServiceImpl) Page(c *gin.Context) {

}

func (listService ListServiceImpl) Create(c *gin.Context) {

}

func (listService ListServiceImpl) Delete(c *gin.Context) {

}

func (listService ListServiceImpl) Update(c *gin.Context) {

}
