package model

type SuccessResponse struct {
	StatusCode int `json:"status_code"`
	Message string `json:"message"`
}