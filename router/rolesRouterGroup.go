package router

import (
	roleController "samplego/controllers"
	"github.com/gin-gonic/gin"
)


func AddRolesApisRoutes(group *gin.RouterGroup) {
	group.GET("/get_role", roleController.GetRoleById)
	group.PUT("/add_role_to_user", roleController.AddRoleToUser)
	group.POST("/create_role", roleController.CreateRole)
}
