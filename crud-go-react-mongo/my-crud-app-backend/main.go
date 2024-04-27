package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)


type User struct {
	// ID string `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
	// Password string `json:"password"`
}

func connectToMongoDB()(*mongo.Client, error) {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return nil, err
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		return nil, err 
	}

	return client, nil
}

// func getUsers (w http.ResponseWriter, r *http.Request){
// 	fmt.Fprint(w, "Get all users",)
// }
// func getUser(w http.ResponseWriter, r *http.Request){
// 	fmt.Fprint(w, "Get a user")
// }

// func updateUser(w http.ResponseWriter, r *http.Request){
// 	fmt.Fprint(w, "Update a user")
// }
// func deleteUser(w http.ResponseWriter, r *http.Request){
// 	fmt.Fprint(w, "Delete a user")
// }

func main(){

	// client, err := connectToMongoDB()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	
	// fmt.Println("Connected to MongoDB!")
	// defer client.Disconnect(context.Background())

	// http.HandleFunc("/users", getUsers)
	// http.HandleFunc("/users/", getUser)
	http.HandleFunc("/user/create", createUser)
	// http.HandleFunc("/user/update/", updateUser)
	// http.HandleFunc("/user/delete/", deleteUser)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "Hello, world!")
	})

	http.ListenAndServe(":8080", nil)
}

func createUser(w http.ResponseWriter, r *http.Request){
	client, err := connectToMongoDB()
	if err != nil {
		log.Fatal(err)
	}
	
	fmt.Println("Connected to MongoDB!")
	defer client.Disconnect(context.Background())

	collection := client.Database("demo").Collection("users")
	newUSer := User{Name: "Oscar", Email: "oscar@mail.com"}
	collection.InsertOne(context.Background(), newUSer)
	
	http.Redirect(w, r, "/", 200)
}