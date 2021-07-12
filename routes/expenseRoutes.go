package routes

import (
	"github.com/eranamarante/go-expense-tracker-api/controllers"
	"github.com/gin-gonic/gin"
)

func ExpenseRoutes(route *gin.Engine) {
	route.GET("/expenses", controllers.GetAllExpenses())
	route.POST("/expenses/new", controllers.AddExpense())
}