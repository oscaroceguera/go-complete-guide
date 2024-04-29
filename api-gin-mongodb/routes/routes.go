package routes

import (
	. "books/database"
	. "books/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListBooks(c *gin.Context){
	listBooks := List_Books()
	if listBooks == nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "List is empty"})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "Data fetched successfully", "data": listBooks})
}

func FindBook(c *gin.Context){
	name := c.Param("name")
	book := Find_Book(name)

	if book == nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "List is empty"})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "Data fetched successfully", "data": book})
}

func CreateBook(c *gin.Context){
	var book Book
	err := c.BindJSON(&book)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Bad request"})
		return 
	}

	id := Create_Book(book)
	c.IndentedJSON(http.StatusCreated, gin.H{"message": "Book Created Succesful", "data": id})
}