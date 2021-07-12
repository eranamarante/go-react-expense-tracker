package main

import (
	"github.com/eranamarante/go-expense-tracker-api/helper"
	"github.com/eranamarante/go-expense-tracker-api/middleware"
	"github.com/eranamarante/go-expense-tracker-api/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	port := helper.GetConfiguration().Port

	router := gin.New()
	router.Use(gin.Logger())
	routes.UserRoutes(router)

	router.Use(middleware.Authentication())

	router.Run(":" + port)
}
