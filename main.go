package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Book struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	PublisherId int    `json:"publisherid"`
	AuthorId    int    `json:"authorid"`
	PublishYear string `json:"publishyear"`
	Price       string `json:"price"`
	Count       int    `json:"count"`
}

type Publisher struct {
	Id   int    `json:"id"`
	Name string `json:"publishername, name"`
}

type Reader struct {
	Id        int    `json:"id"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Surname   string `json:"surname"`
	Address   string `json:"address"`
	Number    string `json:"number"`
	Signature string `json:"signature"`
}

type Author struct {
	Id        int    `json:"id"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Surname   string `json:"surname"`
}

type Gived struct {
	Id        int    `json:"id"`
	ReaderId  int    `json:"readerid"`
	BookId    int    `json:"bookid"`
	Date      string `json:"givedate"`
	Signature string `json:"signature"`
}

var (
	books = []Book{
		{Id: 0, Title: "First book", PublisherId: 0, AuthorId: 0, PublishYear: "1999", Price: "255", Count: 3},
		{Id: 1, Title: "Second book", PublisherId: 1, AuthorId: 1, PublishYear: "2000", Price: "147", Count: 7},
	}

	publishers = []Publisher{
		{Id: 0, Name: "BooksWorld"},
		{Id: 1, Name: "Read Shop"},
	}

	readers = []Reader{
		{Id: 0, FirstName: "Vlad", LastName: "Pronin", Surname: "Aleekseevich", Address: "street Pyshkina dom Kolotyshkina", Number: "88005553535", Signature: "vp"},
		{Id: 1, FirstName: "Maksim", LastName: "Loxin", Surname: "Aleekseevich", Address: "street Ylica dom Mycorka", Number: "88887775544", Signature: "blo"},
	}

	authors = []Author{
		{Id: 0, FirstName: "Oleg", LastName: "Zybkov", Surname: "Andreevich"},
		{Id: 1, FirstName: "Kiril", LastName: "Popov", Surname: "Yagodipopoidavil"},
	}

	gived = []Gived{
		{Id: 0, ReaderId: 0, BookId: 0, Date: "25.01.25"},
		{Id: 1, ReaderId: 1, BookId: 1, Date: "13.11.23"},
	}
)

func main() {
	// новый роутер gin
	router := gin.Default()

	// определение маршрутов
	router.GET("/books", getBooks)
	router.GET("/book/:id", getBookById)

	// запуск сервера с портом 8080
	router.Run(":8080")
}

func getBooks(context *gin.Context) {
	context.JSON(http.StatusOK, books)
}

func getBookById(context *gin.Context) {
	id := context.Param("id")
	for _, book := range books {
		if id == fmt.Sprintf("%d", book.Id) {
			context.JSON(http.StatusOK, book)
			return
		}
	}
	context.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
}
