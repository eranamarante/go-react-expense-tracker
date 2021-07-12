package controllers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/eranamarante/go-expense-tracker-api/helper"
	"github.com/eranamarante/go-expense-tracker-api/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var expenseCollection *mongo.Collection = helper.OpenCollection(helper.Client, "expenses")

func AddExpense() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		var expense models.Expense

		if err := c.BindJSON(&expense); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		validationErr := validate.Struct(expense)
		if validationErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
			return
		}

		expense.Id = primitive.NewObjectID()
		expense.CreatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		expense.UpdatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))

		resultInsertionNumber, insertErr := expenseCollection.InsertOne(ctx, expense)
		if insertErr != nil {
			msg := fmt.Sprintf("Expense item was not created")
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}
		defer cancel()

		c.JSON(http.StatusOK, resultInsertionNumber)
	}
}

func GetAllExpenses() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		var expenses []primitive.M

		findOptions := options.Find()
		findOptions.SetSort(bson.D{primitive.E{Key: "_id", Value: -1}})

		cur, err := expenseCollection.Find(ctx, bson.D{}, findOptions)
		if err != nil {
			log.Fatal(err)
		}

		for cur.Next(ctx) {
			var result bson.M
			e := cur.Decode(&result)
			if e != nil {
				log.Fatal(e)
			}

			expenses = append(expenses, result)
		}

		if err := cur.Err(); err != nil {
			log.Fatal(err)
		}

		c.JSON(http.StatusOK, expenses)
	}
}

func GetExpense() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		id, _ := primitive.ObjectIDFromHex(c.Param("id"))
		filter := bson.M{"_id": id}

		var expense primitive.M
		err := expenseCollection.FindOne(ctx, filter).Decode(&expense)

		if err != nil {
			log.Fatal(err)
		}

		c.JSON(http.StatusOK, expense)
	}
}

func UpdateExpense() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		var expense models.Expense

		id, _ := primitive.ObjectIDFromHex(c.Param("id"))
		if err := c.BindJSON(&expense); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		expense.UpdatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))

		result, err := expenseCollection.UpdateOne(
			ctx,
			bson.M{"_id": id},
			bson.D{
				{Key: "$set", Value: bson.D{
					{Key: "description", Value: expense.Description},
					{Key: "amount", Value: expense.Amount},
					{Key: "due_date", Value: expense.DueDate},
					{Key: "is_paid", Value: expense.IsPaid},
					{Key: "updated_at", Value: expense.UpdatedAt},
				}},
			},
		)
		if err != nil {
			log.Fatal(err)
		}

		c.JSON(http.StatusOK, result)
	}
}
