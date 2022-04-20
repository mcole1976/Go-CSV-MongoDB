package main

import (
	"context"
	"encoding/csv"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
	"strconv"
	"time"
)

func main() {

	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		fmt.Println(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		fmt.Println(err)
	}
	defer client.Disconnect(ctx)

	type Food struct {
		Meal             string    `bson:"Meal,omitempty"`
		Meal_Description string    `bson:"Meal_Description,omitempty"`
		Calorie_Count    int16     `bson:"Calorie_Count,omitempty"`
		Consumption_Date time.Time `bson:"Consumption_Date,omitempty"`
	}
	collection := client.Database("ExerciseDB").Collection("Food_Diary")

	file, err := os.Open("/Users/Marcus/Documents/Code/ExerciseXML/Foodlog.csv")
	if err != nil {
		fmt.Println(err)
	}
	reader := csv.NewReader(file)
	records, _ := reader.ReadAll()
	fmt.Println(records)

	for _, line := range records {
		var i int64 = 0
		i, err1 := strconv.ParseInt(line[2], 10, 64)
		if err1 != nil {
			fmt.Println(err1)
		}
		myDateString := line[3]
		myDate, err2 := time.Parse("2006-01-02 15:04:00", myDateString)
		if err2 != nil {
			fmt.Println(err2)
		}
		//var b int16 = 0
		var b = int16(i)
		data := Food{
			Meal:             line[0],
			Meal_Description: line[1],
			Calorie_Count:    b,
			Consumption_Date: myDate,
		}
		res, insertErr := collection.InsertOne(ctx, data)
		fmt.Println(res)
		if res != nil {
			fmt.Println(res)
		}

		if insertErr != nil {
			fmt.Println(insertErr)
		}
	}

	fmt.Println(records)
}
