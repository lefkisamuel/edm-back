package main

import (
	"edm-back/controller"
	"edm-back/repository"
	"edm-back/service"
	"github.com/gin-gonic/gin"
)
var (
	typeDefinitionRepository repository.TypeDefinitionRepository = repository.New()
	typeDefinitionService service.TypeDefinitionService = service.New(typeDefinitionRepository)
	typeDefinitionController controller.TypeDefinitionController = controller.New(typeDefinitionService)
)


func main(){
	server := gin.Default()
	server.GET("/typedef", func(ctx *gin.Context) {
		ctx.JSON(200, typeDefinitionController.FindAll())
	})
	server.POST("/typedef", func(ctx *gin.Context) {
		ctx.JSON(200, typeDefinitionController.Save(ctx))
	})
	server.Run(":8080")
}