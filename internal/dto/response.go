package dto

type Response struct {
	Status string `json:"status"`
	StatusCode int `json:"status_code"`
	Data interface{} `json:"data"`
}