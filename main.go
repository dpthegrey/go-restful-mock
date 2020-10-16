package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

var authors []Author = []Author{
	Author{
		Id:        "author-1",
		Firstname: "D",
		Lastname:  "P",
		Username:  "dp",
		Password:  "pass",
	},
	Author{
		Id:        "author-2",
		Firstname: "Maria",
		Lastname:  "Raboy",
		Username:  "mraboy",
		Password:  "abc123",
	},
}

var articles []Article = []Article{
	Article{
		Id:      "article-1",
		Author:  "author-1",
		Title:   "Blog Post 1",
		Content: "This is an example blog article",
	},
}

func RootEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	response.Write([]byte(`{ "Message": "Hello World" }`))
}

func main() {
	fmt.Println("Starting the application...")
	router := mux.NewRouter()
	router.HandleFunc("/", RootEndpoint).Methods("GET")
	http.ListenAndServe(":12345", router)
}
