package controller

import (
	"edm-back/entity"
	"edm-back/service"
	"github.com/gin-gonic/gin"
)

type TypeDefinitionController interface {
	Save(ctx *gin.Context) entity.TypeDefinition
	FindAll() []entity.TypeDefinition
}

type controller struct {
	service service.TypeDefinitionService
}

func New(s service.TypeDefinitionService) TypeDefinitionController{
	return &controller{
		service: s,
	}
}

func (c *controller) Save(ctx *gin.Context) entity.TypeDefinition{
	var td entity.TypeDefinition
	_ = ctx.BindJSON(&td)
	c.service.Save(td)
	return td
}
func (c *controller) FindAll() []entity.TypeDefinition{
	return c.service.FindAll()
}
