package main

import (
	"github.com/gorilla/mux"
)

func setupRoutes(r *mux.Router) {
	r.HandleFunc("/getdutchwords", getDutchWords).Methods("GET")
	r.HandleFunc("/getenglishwords", getEnglishWords).Methods("GET")
	r.HandleFunc("/translation/{word}", getTranslation).Methods("GET")
	r.HandleFunc("/search/{query}", searchWords).Methods("GET")
	r.HandleFunc("/randomdutchword", getRandomDutchWord).Methods("GET")
	r.HandleFunc("/randomenglishword", getRandomEnglishWord).Methods("GET")
}

type Word struct {
	ID       int    `json:"id"`
	Text     string `json:"text"`
	Language string `json:"language"`
}

func NewWord(id int, text, language string) Word {
	return Word{
		ID:       id,
		Text:     text,
		Language: language,
	}
}
