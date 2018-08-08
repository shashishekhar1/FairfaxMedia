package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Article Article structure
type Article struct {
	ID    string   `json:"id,omitempty"`
	TITLE string   `json:"title,omitempty"`
	DATE  string   `json:"date,omitempty"`
	BODY  string   `json:"body,omitempty"`
	TAGS  []string `json:"tags,omitempty"`
}

// articles acting as a in memory storage
var articles []Article

// CreateArticle Takes a  json request and creates an article
func CreateArticle(w http.ResponseWriter, r *http.Request) {
	var article Article
	err := json.NewDecoder(r.Body).Decode(&article)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Invalid content.")
		return
	}

	articles = append(articles, article)
	log.Print("New Article created with id: ", article.ID)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(article)
}

// GetArticle Get an article using article id
func GetArticle(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	log.Print("Get article id ", params["id"])

	for _, item := range articles {
		if item.ID == params["id"] {
			log.Print("Article found with: ", params["id"])
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusFound)
			json.NewEncoder(w).Encode(item)
			return
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(&Article{})
}

// GetDateBasedTagName Returns a list of articles with same tag name and day
func GetDateBasedTagName(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	var tag = params["tagName"]
	var date = changeDateFormat(params["date"])

	log.Print("Get articles based on tag ", tag, " and date ", date)
	var tagresponse = filterAndCreateTagResponse(articles, tag, date)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tagresponse)
}
