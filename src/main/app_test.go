package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestChangeDateFormat(t *testing.T) {
	date := changeDateFormat("20180515")
	if date != "2018-05-15" {
		t.Errorf("Sum was incorrect, got: %s, want: %s.", date, "2018-05-15")
	}
}

func TestFilterAndCreateTagResponse(t *testing.T) {
	searchTag := "health"
	searchDate := "2016-09-22"
	response := filterAndCreateTagResponse(getArticles(), searchTag, searchDate)
	if response.Tag != searchTag {
		t.Errorf("Tag was incorrect, got: %s, want: %s.", response.Tag, searchTag)
	}
	if response.Count != 2 {
		t.Errorf("Article count returned was incorrect, got: %d, want: %d.", response.Count, 2)
	}
}

func TestCreateArticle(t *testing.T) {
	var jsonStr = []byte(`
  {
    "id": "1",
     "title": "latest science",
     "date" : "2016-09-22",
     "body" : "some body text",
     "tags" : ["health", "science"]
  }`)
	req, err := http.NewRequest("POST", "/articles", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Error(err)
	}

	w := httptest.NewRecorder()
	CreateArticle(w, req)
	if w.Code != 201 {
		t.Errorf("In-correct status , got: %d, want: %d.", w.Code, 201)
	}
}

func TestGetArticle(t *testing.T) {
	req, err := http.NewRequest("GET", "/articles/1", nil)
	if err != nil {
		t.Error(err)
	}

	w := httptest.NewRecorder()
	GetArticle(w, req)
	if w.Code != 404 {
		t.Errorf("In-correct status , got: %d, want: %d.", w.Code, 404)
	}
}

func getArticles() []Article {
	var articles []Article
	var article1 = Article{
		ID:    "1",
		TITLE: "latest science",
		DATE:  "2016-09-22",
		BODY:  "Text body",
		TAGS:  []string{"health", "horror", "science"},
	}
	var article2 = Article{
		ID:    "2",
		TITLE: "latest fitness",
		DATE:  "2016-09-22",
		BODY:  "Text body",
		TAGS:  []string{"health", "fitness", "exercise"},
	}
	var article3 = Article{
		ID:    "3",
		TITLE: "testing 3",
		DATE:  "2016-09-23",
		BODY:  "Text body",
		TAGS:  []string{"health", "fitness", "exercise"},
	}
	var article4 = Article{
		ID:    "4",
		TITLE: "testing 4",
		DATE:  "2016-09-22",
		BODY:  "Text body",
		TAGS:  []string{"horror", "test"},
	}
	articles = append(articles, article1)
	articles = append(articles, article2)
	articles = append(articles, article3)
	articles = append(articles, article4)

	return articles
}
