package helper

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

type ErrorResponse struct {
	StatusCode   int    `json:"status"`
	ErrorMessage string `json:"message"`
}

type Configuration struct {
	Port             string
	ConnectionString string
	DatabaseName     string
	SecretKey        string
}

func GetError(err error, w http.ResponseWriter) {
	log.Fatal(err.Error())
	var response = ErrorResponse{
		StatusCode:   http.StatusInternalServerError,
		ErrorMessage: err.Error(),
	}

	message, _ := json.Marshal(response)

	w.WriteHeader(response.StatusCode)
	w.Write(message)
}

func GetConfiguration() Configuration {
	err := godotenv.Load(os.ExpandEnv("$GOPATH/src/github.com/eranamarante/go-expense-tracker-api/.env"))
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	configuration := Configuration{
		os.Getenv("PORT"),
		os.Getenv("MONGO_URL"),
		os.Getenv("MONGO_DATABASE"),
		os.Getenv("SECRET_KEY"),
	}

	return configuration
}
