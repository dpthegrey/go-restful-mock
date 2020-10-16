package main

type Article struct {
	Id      string `json:"Id,omitempty"`
	Author  string `json:"author,omitempty"`
	Title   string `json:"title,omitempty"`
	Content string `json:"content,omitempty"`
}
