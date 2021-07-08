package helper

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

// ErrorResponse : This is error model.
type ErrorResponse struct {
	StatusCode   int    `json:"status"`
	ErrorMessage string `json:"message"`
}

// GetError : This is helper function to prepare error model.
// If you want to export your function. You must to start upper case function name. Otherwise you won't see your function when you import that on other class.
func GetError(err error, w http.ResponseWriter) {

	log.Fatal(err.Error())
	var response = ErrorResponse{
		ErrorMessage: err.Error(),
		StatusCode:   http.StatusInternalServerError,
	}

	message, _ := json.Marshal(response)

	w.WriteHeader(response.StatusCode)
	w.Write(message)
}

// Configuration model
type Configuration struct {
	Port             string
	ConnectionString string
	DatabaseName     string
}

// GetConfiguration method basically populate configuration information from .env and return Configuration model
func GetConfiguration() Configuration {
	err := godotenv.Load(os.ExpandEnv("$GOPATH/src/github.com/eranamarante/go-react-expense-tracker/.env"))

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	configuration := Configuration{
		os.Getenv("PORT"),
		os.Getenv("MONGODB_URL"),
		os.Getenv("MONGODB_DATABASE"),
	}

	return configuration
}
