package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"runtime/debug"
	"samplego/dao/implementations"
	"strconv"
)

func GetResourceById(ctx *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			ErrorString := "Panic Recoved , Error in getResourceById and Error = " + fmt.Sprint(r) + " and stack trace = " + string(debug.Stack())
			log.Fatal(ErrorString)
		}
	}()

	var resourceDao implementations.ResourceDaoImpl
	id, _ := strconv.Atoi(ctx.Query("id"))
	ctx.JSON(200,resourceDao.GetResourceById(int64(id)))
}
