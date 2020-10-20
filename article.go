package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	uuid "github.com/satori/go.uuid"
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
	validate := validator.New()
	err := validate.Struct(article)
	if err != nil {
		response.WriteHeader(500)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	article.Id = uuid.Must(uuid.NewV4()).String()
	article.Author = "dp"
	articles = append(articles, article)
	json.NewEncoder(response).Encode(articles)
}

func ArticleDeleteEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	params := mux.Vars(request)
	for index, article := range articles {
		if article.Id == params["id"] {
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
	for index, article := range articles {
		if article.Id == params["id"] {
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
