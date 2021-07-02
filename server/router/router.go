package router

import "github.com/gorilla/mux"

func Router() *mux.Router {
	r := mux.NewRouter()

	// Auth Endpoints
	r.HandleFunc("api/auth/login", authMiddleware.Login).Methods("POST", "OPTIONS")
	r.HandleFunc("api/auth/login", authMiddleware.Signup).Methods("POST", "OPTIONS")
	
	// Expense Endpoints
	r.HandleFunc("api/expenses", expenseMiddleware.GetAllExpenses).Methods("GET", "OPTIONS")
	r.HandleFunc("api/expenses/{id}", expenseMiddleware.GetExpense).Methods("GET", "OPTIONS")
	r.HandleFunc("api/expenses/new", expenseMiddleware.AddExpense).Methods("POST", "OPTIONS")
	r.HandleFunc("api/expenses/{id}/edit", expenseMiddleware.UpdateExpense).Methods("PUT", "OPTIONS")
	r.HandleFunc("api/expenses/{id}/delete", expenseMiddleware.DeleteExpense).Methods("DELETE", "OPTIONS")
}