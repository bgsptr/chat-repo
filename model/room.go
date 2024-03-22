package model

import "groupservice/model"

type Room struct {
	ID     string    `json:"id"`
	Name   string `json:"name"`
	Clients []*model.Client `json:"clients"`
}