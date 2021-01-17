package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"runtime/debug"
	"samplego/dao/implementations"
	"strconv"
)

func GetUserById(ctx *gin.Context) {
	defer func() {
		if r := recover(); r!=nil {
			ErrorString := "Panic Recoved , Error in GetUserById and Error = " + fmt.Sprint(r) + " and stack trace = " + string(debug.Stack())
			log.Fatal(ErrorString)
		}
	}()

	var userImpl implementations.UserImpl
	id, _ := strconv.Atoi(ctx.Query("id"))
	ctx.JSON(200,userImpl.GetUserById(int64(id)))
	return
}

