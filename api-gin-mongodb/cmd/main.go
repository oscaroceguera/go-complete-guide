package main

import (
	"log"

	. "books/database"
	. "books/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main(){
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	defer Disconnect()
	router := gin.Default()
	router.POST("/books", CreateBook)
	router.GET("/books", ListBooks)
	router.Run("localhost:8080")
}