package models

type Response struct {
	Result interface{} `json:"result"`
	Error  string      `json:"error"`
}
