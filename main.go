package main

import (
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

// func postBook(w http.ResponseWriter,r *http.Request){
// 	var b Book
// 	err := json.NewDecoder(r.Body).Decode(&b)
// 	if err!=nil{
// 		log.Fatal(err)
// 	}
// 	books = append(books, b)
// 	json.NewEncoder(w).Encode(b)
// }

// func getBooks(w http.ResponseWriter, r *http.Request) {
//     // w.Header().Set("Content-Type", "application/json")
//     if err := json.NewEncoder(w).Encode(books); err != nil {
//         http.Error(w, err.Error(), http.StatusInternalServerError)
//         return
//     }
// }


func getBooks(c *gin.Context){
	c.IndentedJSON(http.StatusOK,books)
}

func addBook(c *gin.Context){
	var newBook Book
	if err := c.BindJSON(&newBook);err!=nil
}

func main(){
	router := gin.Default()
	router.GET("/books",getBooks)
	router.Run("localhost:8080")
}