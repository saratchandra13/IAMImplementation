package router

import (
	"github.com/gin-gonic/gin"
	"samplego/controllers"
)

func AddUsersApiRoutes(group *gin.RouterGroup) {
	group.GET("/get_user", controllers.GetUserById)
}
