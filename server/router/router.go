package router

import (
	"github.com/eranamarante/go-react-expense-tracker/server/middleware"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	r := mux.NewRouter()

	// Auth Endpoints
	// r.HandleFunc("api/auth/login", authMiddleware.Login).Methods("POST", "OPTIONS")
	// r.HandleFunc("api/auth/login", authMiddleware.Signup).Methods("POST", "OPTIONS")

	// Expense Endpoints
	r.HandleFunc("/api/expenses", middleware.GetAllExpenses).Methods("GET")
	r.HandleFunc("/api/expenses/{id}", middleware.GetExpense).Methods("GET")
	r.HandleFunc("/api/expenses/new", middleware.AddExpense).Methods("POST")
	r.HandleFunc("/api/expenses/{id}/edit", middleware.UpdateExpense).Methods("PUT")
	r.HandleFunc("/api/expenses/{id}/delete", middleware.DeleteExpense).Methods("DELETE")
	r.HandleFunc("/api/expenses/{id}/markAsPaid", middleware.MarkExpenseAsPaid).Methods("PUT")
	r.HandleFunc("/api/expenses/{id}/markAsUnpaid", middleware.MarkExpenseAsUnpaid).Methods("PUT")

	return r
}
