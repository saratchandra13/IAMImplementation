package main

import (
	"log"
	"github.com/gin-gonic/gin"
	"samplego/router"
)


func main() {
	mainRouter := gin.Default()

	// add all the routes present in the search here.
	rolesRouterGroup := mainRouter.Group("/role")
	userRouterGroup := mainRouter.Group("/user")
	resourceRouterGroup := mainRouter.Group("/resource")

	router.AddResourceApiRoutes(resourceRouterGroup)
	router.AddUsersApiRoutes(userRouterGroup)
	router.AddRolesApisRoutes(rolesRouterGroup)
	// can add more functionalities here.

	if err := mainRouter.Run(":8080");err!=nil{
		log.Fatal("Fatal Error! Error while starting the server. Shutting Down the server. Goodbye From Server. :( ")
	}else{
		log.Println("Hi! From the server, :D")
	}
}
