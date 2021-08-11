package models

import "time"

type RequestData struct {
	URL string `json:"url"`
}

type URLData struct {
	Alias       string    `json:"alias"`
	OriginalURL string    `json:"originalURL"`
	CreatedOn   time.Time `json:"createdOn"`
	ShortURL    string    `json:"shortURL"`
}

type Response struct {
	Result interface{} `json:"result"`
	Error  string      `json:"error"`
}
