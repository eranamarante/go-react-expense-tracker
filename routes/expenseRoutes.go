package routes

import (
	"github.com/eranamarante/go-expense-tracker-api/controllers"
	"github.com/gin-gonic/gin"
)

func ExpenseRoutes(route *gin.Engine) {
	route.GET("/expenses", controllers.GetAllExpenses())
	route.POST("/expenses/new", controllers.AddExpense())
	route.GET("/expenses/:id/", controllers.GetExpense())
	route.PUT("/expenses/:id/edit", controllers.UpdateExpense())
	route.DELETE("/expenses/:id/delete", controllers.DeleteExpense())
}
