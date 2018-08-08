package main

import (
	"log"
	"time"
)

// TagSearchResponse response by tag search
type TagSearchResponse struct {
	Tag          string   `json:"tag"`
	Count        int      `json:"count"`
	Articles     []string `json:"article"`
	Related_tags []string `json:"related_tags"`
}

// changeDateFormat changes the format of date from YYYYMMDD to YYYY-MM-DD and returns a string
func changeDateFormat(ipDateString string) string {

	date, err := time.Parse("20060102", ipDateString)
	if err != nil {
		log.Fatal("Error in date conversion ", err)
	}
	date2 := date.Format("2006-01-02")
	log.Print("converted  date ", date.Format("YYYY-MM-DD"))

	return date2
}

// filterAndCreateTagResponse filters the articles based on date and tags
func filterAndCreateTagResponse(articles []Article, tag string, date string) TagSearchResponse {

	var filterArticles []string
	var check bool
	check = false

	alltags := make(map[string]string)
	for _, item := range articles {

		if item.DATE == date {
			for _, t := range item.TAGS {
				if t == tag {
					check = true
					break
				}
			}
			if check {
				for _, tag := range item.TAGS {
					alltags[tag] = ""
				}
				filterArticles = append(filterArticles, item.ID)
			}
			check = false
		}
	}
	delete(alltags, tag)
	var allkeys []string
	for k := range alltags {
		allkeys = append(allkeys, k)
	}
	var tagresponse = TagSearchResponse{
		Tag:          tag,
		Count:        len(filterArticles),
		Articles:     filterArticles,
		Related_tags: allkeys,
	}

	return tagresponse
}
