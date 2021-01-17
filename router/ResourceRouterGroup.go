package router

import (
	"github.com/gin-gonic/gin"
	resourceController "samplego/controllers"
)

func AddResourceApiRoutes(group *gin.RouterGroup) {
	group.GET("/get_resource", resourceController.GetResourceById)
}
