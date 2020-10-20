package main

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"github.com/mitchellh/mapstructure"
)

type CustomJWTClaim struct {
	Id string `json:"id"`
	jwt.StandardClaims
}

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

var JWT_SECRET []byte = []byte("dpthegrey")

func ValidateJWT(t string) (interface{}, error) {
	token, err := jwt.Parse(t, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method %v", token.Header["alg"])
		}
		return JWT_SECRET, nil
	})
	if err != nil {
		return nil, errors.New(`{ "message": "` + err.Error() + `" }`)
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		var tokenData CustomJWTClaim
		mapstructure.Decode(claims, &tokenData)
		return tokenData, nil
	} else {
		return nil, errors.New(`{ "message": "invalid token" }`)
	}
}

func RootEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	response.Write([]byte(`{ "Message": "Hello World" }`))
}

func main() {
	fmt.Println("Starting the application...")
	router := mux.NewRouter()
	router.HandleFunc("/", RootEndpoint).Methods("GET")
	router.HandleFunc("/register", RegisterEndpoint).Methods("POST")
	router.HandleFunc("/login", AuthorRetrieveAllEndpoint).Methods("POST")
	router.HandleFunc("/authors", AuthorRetrieveAllEndpoint).Methods("GET")
	router.HandleFunc("/author/{id}", AuthorRetrieveEndpoint).Methods("GET")
	router.HandleFunc("/author/{id}", AuthorDeleteEndpoint).Methods("DELETE")
	router.HandleFunc("/author/{id}", AuthorUpdateEndpoint).Methods("PUT")
	router.HandleFunc("/articles", RootEndpoint).Methods("GET")
	router.HandleFunc("/authors", ArticleRetrieveAllEndpoint).Methods("GET")
	router.HandleFunc("/article/{id}", ArticleRetrieveEndpoint).Methods("GET")
	router.HandleFunc("/article/{id}", ArticleDeleteEndpoint).Methods("DELETE")
	router.HandleFunc("/article/{id}", ArticleUpdateEndpoint).Methods("PUT")
	router.HandleFunc("/article", ArticleCreateEndpoint).Methods("POST")

	http.ListenAndServe(":12345", router)
}
