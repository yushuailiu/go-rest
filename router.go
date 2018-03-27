package main

import (
	. "handlers"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	router.POST("/person", AddPersonHandler)
	router.DELETE("/person/:id", DeletePersonHandler)
	router.PUT("/person/:id", UpdatePersonHandler)
	router.GET("/persons/:page/:number", ListPersonHandler)
	router.GET("/", IndexHandler)
	router.GET("/person/:id", GetPersonHandler)
	return router
}
