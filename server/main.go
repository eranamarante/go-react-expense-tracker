package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/eranamarante/go-react-expense-tracker/server/helper"
	"github.com/eranamarante/go-react-expense-tracker/server/router"
)

func main() {
	port := helper.GetConfiguration().Port

	r := router.Router()

	fmt.Printf("Starting server on the port %v ... \n", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
