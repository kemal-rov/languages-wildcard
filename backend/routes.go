package main

import (
	"encoding/json"
	"net/http"

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
	ID              int    `json:"id"`
	Word            string `json:"word"`
	WordType        string `json:"word_type"`
	WordDescription string `json:"word_description"`
	WordExample     string `json:"word_example"`
	WordTranslation string `json:"word_translation"`
}

func getDutchWords(w http.ResponseWriter, r *http.Request) {
	words := []Word{}

	rows, err := db.Query("SELECT * FROM languages_words WHERE language = 'NL'")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var word Word
		if err := rows.Scan(&word.ID, &word.Word, &word.WordType, &word.WordDescription, &word.WordExample, &word.WordTranslation); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		words = append(words, word)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(words)
}

func getEnglishWords(w http.ResponseWriter, r *http.Request) {
	words := []Word{}

	rows, err := db.Query("SELECT * FROM languages_words WHERE language = 'EN'")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var word Word
		if err := rows.Scan(&word.ID, &word.Word, &word.WordType, &word.WordDescription, &word.WordExample, &word.WordTranslation); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		words = append(words, word)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(words)
}
