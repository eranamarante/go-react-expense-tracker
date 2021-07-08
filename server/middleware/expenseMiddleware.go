package middleware

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/eranamarante/go-react-expense-tracker/server/database"
	"github.com/eranamarante/go-react-expense-tracker/server/helper"
	"github.com/eranamarante/go-react-expense-tracker/server/models"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var expensesCollection *mongo.Collection = database.OpenCollection(database.Client, "expenses")

func GetAllExpenses(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

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

		results = append(results, result)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	cur.Close(context.Background())

	json.NewEncoder(w).Encode(results)
}

func GetExpense(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	id, _ := primitive.ObjectIDFromHex(params["id"])

	var expense bson.M
	filter := bson.M{"_id": id}
	if err := expensesCollection.FindOne(context.Background(), filter).Decode(&expense); err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(expense)
}

func AddExpense(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var expense models.Expense
	json.NewDecoder(r.Body).Decode(&expense)

	insertResult, err := expensesCollection.InsertOne(context.Background(), expense)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted a Single Record ", insertResult.InsertedID)
	json.NewEncoder(w).Encode(expense)
}

func UpdateExpense(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	id, _ := primitive.ObjectIDFromHex(params["id"])

	var expense models.Expense

	filter := bson.M{"_id": id}

	_ = json.NewDecoder(r.Body).Decode(&expense)

	update := bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "title", Value: expense.Title},
			{Key: "amount", Value: expense.Amount},
			{Key: "due_date", Value: expense.DueDate},
			{Key: "is_paid", Value: expense.IsPaid},
			{Key: "updated_at", Value: time.Now()},
		}},
	}

	err := expensesCollection.FindOneAndUpdate(context.Background(), filter, update).Decode(&expense)

	if err != nil {
		helper.GetError(err, w)
		return
	}

	expense.Id = id

	json.NewEncoder(w).Encode(expense)
}

func DeleteExpense(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	id, _ := primitive.ObjectIDFromHex(params["id"])

	filter := bson.M{"_id": id}
	deleteResult, err := expensesCollection.DeleteOne(context.Background(), filter)

	if err != nil {
		helper.GetError(err, w)
		return
	}

	json.NewEncoder(w).Encode(deleteResult)
}

func MarkExpenseAsPaid(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	id, _ := primitive.ObjectIDFromHex(params["id"])

	result, err := expensesCollection.UpdateOne(
		context.Background(),
		bson.M{"_id": id},
		bson.D{
			{Key: "$set", Value: bson.D{
				{Key: "is_paid", Value: true},
				{Key: "updated_at", Value: time.Now()},
			}},
		},
	)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Updated %v Documents!\n", result.ModifiedCount)
	json.NewEncoder(w).Encode(result)
}

func MarkExpenseAsUnpaid(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	id, _ := primitive.ObjectIDFromHex(params["id"])

	result, err := expensesCollection.UpdateOne(
		context.Background(),
		bson.M{"_id": id},
		bson.D{
			{Key: "$set", Value: bson.D{
				{Key: "is_paid", Value: false},
				{Key: "updated_at", Value: time.Now()},
			}},
		},
	)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Updated %v Documents!\n", result.ModifiedCount)
	json.NewEncoder(w).Encode(result)
}
