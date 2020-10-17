package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type Author struct {
	Id        string `json:"id,omitempty"`
	Firstname string `json:"firstname,omitempty"`
	Lastname  string `json:"lastname,omitempty"`
	Username  string `json:"username,omitempty"`
	Password  string `json:"password,omitempty"`
}

func AuthorRetrieveAllEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	json.NewEncoder(response).Encode(authors)
}

func AuthorRetrieveEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	params := mux.Vars(request)
	for _, author := range authors {
		if author.Id == params["id"] {
			json.NewEncoder(response).Encode(author)
			return
		}
	}
	json.NewEncoder(response).Encode(Author{})
}

func AuthorDeleteEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	params = mux.Vars(request)
	for index, author := range authors {
		if author.id == params["id"] {
			authors = append(authors[:index], authors[index+1:]...)
			json.NewEncoder(response).Encode(authors)
			return
		}
	}
	json.NewEncoder(response).Encode(Author{})
}

func AuthorUpdateEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	params = mux.Vars(request)
	var changes Author
	json.NewDecoder(request.Body).Decode(&changes)
	for index, author := range authors {
		if author.Id == params["id"] {
			if changes.Firstname != "" {
				author.Firstname = changes.Firstname
			}
			if changes.Lastname != "" {
				author.Lastname = changes.Lastname
			}
			if changes.Username != "" {
				author.Username = changes.Username
			}
			if changes.Password != "" {
				author.Password = changes.Password
			}
			authors[index] = author
			json.NewEncoder(response).Encode(authors)
			return
		}
	}
	json.NewEncoder(response).Encode(Author{})
}
