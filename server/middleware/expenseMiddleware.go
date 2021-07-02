package middleware

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/eranamarante/go-react-expense-tracker/server/database"
	"github.com/eranamarante/go-react-expense-tracker/server/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var expensesCollection *mongo.Collection = database.OpenCollection(database.Client, "expenses")

func GetAllExpenses(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	payload := getAllExpenses()
	json.NewEncoder(w).Encode(payload)
}

func getAllExpenses() []primitive.M  {
	findOptions := options.Find()
	findOptions.SetSort(bson.D{primitive.E{Key: "_id", Value: -1}})

	cur, err := expensesCollection.Find(context.Background(), bson.D{}, findOptions)
	if err != nil {
		log.Fatal(err)
	}

	var results []primitive.M
	for cur.Next(context.Background()) {
		var result bson.M
		e := cur.Decode(&result)
		if e != nil {
			log.Fatal(e)
		}

		results =  append(results, result)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	cur.Close(context.Background())
	return results
}

func AddExpense(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	
	var expense models.Expense
	json.NewDecoder(r.Body).Decode(&expense)

	addExpense(expense)
	json.NewEncoder(w).Encode(expense)
}

func addExpense(task models.Expense) {
	insertResult, err := expensesCollection.InsertOne(context.Background(), task)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted a Single Record ", insertResult.InsertedID)
}
