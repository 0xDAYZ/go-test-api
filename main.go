package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
	// "errors"
)

type book struct {
	ID			string	`json: "id"`  // Serialize Public field ID to json 
	Title   	string	`json: "title"`
	Author  	string	`json: "author"`
	Quantity	int		`json: "quantity"`
}

var books = []book{
	{ID: "1", Title: "In Search of Lost Time", Author: "Marcel Proust", Quantity: 2},
	{ID: "2", Title: "The Great Gatsby", Author: "F. Scott Fitzgerald", Quantity: 5},
	{ID: "3", Title: "War and Peace", Author: "Leo Tolstoy", Quantity: 6},
}

func getBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)
}

func main() {
	router := gin.Default() // Handles routes to different pages in a website
	router.GET("/books", getBooks) // GET endpoint
	router.Run("localhost:8081")
}