package main

import (
	"errors"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Book struct{
	ID 			string 	`json:"id"`
	Title 		string	`json:"title"`
	Author 		string	`json:"author"`
	Quantity 	int		`json:"quantity"`
}

var books = []Book{
	{ID:"1",Title:"ABCD",Author: "Nitish",Quantity: 2},
}


func getBooks(c *gin.Context){
	c.IndentedJSON(http.StatusOK,books)
}

func bookById(c *gin.Context){
	id := c.Param("id")
	book,err := getBookById(id)
	if err!=nil{
		log.Fatal(err)
	}
	c.IndentedJSON(http.StatusOK,book)
}

func getBookById(id string)(*Book,error){
	for i,b:= range books{
		if b.ID == id{
			return &books[i],nil
		}
	}
	return nil,errors.New("Book not found")
}

func addBook(c *gin.Context){
	var newBook Book
	if err:=c.BindJSON(&newBook); err!=nil{
		return
	}
	books = append(books,newBook)
	c.IndentedJSON(http.StatusCreated,newBook)
}

func main(){
	router := gin.Default()
	router.GET("/books",getBooks)
	router.POST("/books",addBook )
	router.GET("/books/:id",bookById)
	router.Run("localhost:8080")
}