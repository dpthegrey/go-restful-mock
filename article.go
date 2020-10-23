package main

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	"gopkg.in/go-playground/validator.v9"
)

type Article struct {
	Id      string `json:"Id,omitempty" validate:"omitempty,uuid"`
	Author  string `json:"author,omitempty" validate:"isdefault"`
	Title   string `json:"title,omitempty" validate:"required"`
	Content string `json:"content,omitempty" validate:"required"`
}

func ArticleRetrieveAllEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	json.NewEncoder(response).Encode(articles)
}

func ArticleRetrieveEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	params := mux.Vars(request)
	for _, article := range articles {
		if article.Id == params["id"] {
			json.NewEncoder(response).Encode(article)
			return
		}
	}
	json.NewEncoder(response).Encode(Article{})
}

func ArticleCreateEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	var article Article
	json.NewDecoder(request.Body).Decode(&article)
	token := context.Get(request, "decoded").(CustomJWTClaim)

	validate := validator.New()
	err := validate.Struct(article)
	if err != nil {
		response.WriteHeader(500)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	article.Id = uuid.Must(uuid.NewRandom()).String()
	article.Author = token.Id
	articles = append(articles, article)
	json.NewEncoder(response).Encode(articles)
}

func ArticleDeleteEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	params := mux.Vars(request)
	token := context.Get(request, "decoded").(CustomJWTClaim)
	for index, article := range articles {
		if article.Id == params["id"] && article.Author == token.Id {
			articles = append(articles[:index], articles[index+1:]...)
			json.NewEncoder(response).Encode(article)
			return
		}
	}
	json.NewEncoder(response).Encode(Article{})
}

func ArticleUpdateEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	params := mux.Vars(request)
	var changes Article
	json.NewDecoder(request.Body).Decode(&changes)
	token := context.Get(request, "decoded").(CustomJWTClaim)
	for index, article := range articles {
		if article.Id == params["id"] && article.Author == token.Id {
			if changes.Title != "" {
				article.Title = changes.Title
			}
			if changes.Content != "" {
				article.Content = changes.Content
			}
			articles[index] = article
			json.NewEncoder(response).Encode(articles)
			return
		}
	}
	json.NewEncoder(response).Encode(Article{})
}
