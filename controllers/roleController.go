package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"runtime/debug"
	"samplego/dao/implementations"
	"samplego/models"
	"samplego/services"
	"strconv"
)

func GetRoleById(ctx *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			ErrorString := "Panic Recovered , Error in GetRoleById and Error = " + fmt.Sprint(r) + " and stack trace = " + string(debug.Stack())
			log.Fatal(ErrorString)
		}
	}()

	var roleDoaImpl implementations.RoleDaoImpl
	id, _ := strconv.Atoi(ctx.Query("id"))
	ctx.JSON(200,roleDoaImpl.GetRoleById(int64(id)))
	return
}

func AddRoleToUser(ctx *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			ErrorString := "Panic Recovered , Error in AddRoleToUser and Error = " + fmt.Sprint(r) + " and stack trace = " + string(debug.Stack())
			log.Fatal(ErrorString)
		}
	}()

	var roleService services.RoleService
	roleId, _ := strconv.Atoi(ctx.Query("role_id"))
	userId, _ := strconv.Atoi(ctx.Query("user_id"))
	ctx.JSON(200,roleService.AddRoleToUser(int64(roleId), int64(userId)))
	return
}

func CreateRole(ctx *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			ErrorString := "Panic Recovered , Error in CreateRole and Error = " + fmt.Sprint(r) + " and stack trace = " + string(debug.Stack())
			log.Fatal(ErrorString)
		}
	}()

	var roleService services.RoleService
	request := models.CreateRoleRequest{}
	_ = ctx.BindJSON(&request)
	ctx.JSON(200, roleService.CreateRole(request.ResourceId,request.Perms,request.RoleName))
	return
}