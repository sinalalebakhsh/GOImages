package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type book struct {
	/* Upper case is for exported to another modules */
	ID       string `json:"id"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Quantity int    `json:"quantity"`
}

var books = []book{
	{ID: "1", Title: "In Search of Lost Time", Author: "Marcel Proust", Quantity: 2},
	{ID: "2", Title: "The Great Gatsby", Author: "F. Scott Fitzgerald", Quantity: 4},
	{ID: "3", Title: "War and Peace", Author: "Leo Tolstoy", Quantity: 3},
	{ID: "4", Title: "The Little Prince", Author: "Antoine de Saint-Exupéry", Quantity: 5},
}

// This function return a json version of books variable
func getBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)
}

// This function Create new book and add it into "books" objects
func createBook(c *gin.Context) {
	var newBook book

	if err := c.BindJSON(&newBook); err != nil {
		return
	}

	books = append(books, newBook)
	c.IndentedJSON(http.StatusOK, newBook)
}

// This function uses getBookById() function
func bookById(c *gin.Context) {
	id := c.Param("id")
	// example: "/books/2"

	book, err := getBookById(id)
	// err Stand for error

	// nil is like empty glass of wine,
	// if is empty so person is sober,
	// on the contrary person is drunk!! and glass IS NOT EMPTY != False not nil
	if err != nil {
		// We have to create custom message for error arising in here
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found!"})
		return
	}

	c.IndentedJSON(http.StatusOK, book)
}

func checkoutBook(c *gin.Context) {
	id, ok := c.GetQuery("id")

	// Which means if "OK" was equal to false, so that id doesn't exist.
	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Missing id query parameter."})
		return
	}

	book, err := getBookById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found!"})
		return
	}

	if book.Quantity <= 0 {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Book not available."})
		return
	}

	book.Quantity -= 1
	c.IndentedJSON(http.StatusOK, book)
}

func returnBook(c *gin.Context) {
	id, ok := c.GetQuery("id")

	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Missing id query parameter."})
		return
	}

	book, err := getBookById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found!"})
		return
	}

	book.Quantity += 1
	c.IndentedJSON(http.StatusOK, book)
}

// This function have to helps for getting books by id
func getBookById(id string) (*book, error) {
	for i, b := range books {
		if b.ID == id {
			return &books[i], nil
		}
	}
	return nil, errors.New("book not found")
}

func main() {
	router := gin.Default()
	router.GET("/books", getBooks)

	// id := c.Param("id")
	// This uses in here
	router.GET("/books/:id", bookById)
	router.POST("/books", createBook)
	router.PATCH("/checkout", checkoutBook)
	router.PATCH("/return", returnBook)
	router.Run("localhost:8080")
}
