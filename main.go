package main

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	// Connect to MongoDB
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://challengeUser:WUMglwNBaydH8Yvu@cluster0.jugvyba.mongodb.net/?retryWrites=true&w=majority"))
	if err != nil {
		fmt.Println(err)
		return
	}
	err = client.Connect(context.TODO())
	if err != nil {
		fmt.Println(err)
		return
	}

	// Get a handle to the "users" collection
	collection := client.Database("mydb").Collection("users")

	// Find all users with age >= 25 and sort by name in ascending order
	filter := bson.M{"age": bson.M{"$gte": 25}}
	projection := bson.M{"name": 1}
	sort := bson.M{"name": 1}
	cursor, err := collection.Find(context.TODO(), filter, options.Find().SetProjection(projection).SetSort(sort))
	if err != nil {
		fmt.Println(err)
		return
	}
	defer cursor.Close(context.TODO())

	// Iterate through the documents and print them
	var users []interface{}
	if err = cursor.All(context.TODO(), &users); err != nil {
		fmt.Println(err)
		return
	}
	for _, user := range users {
		fmt.Println(user)
	}
}
