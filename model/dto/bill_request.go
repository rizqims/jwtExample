package dto

import "apilaundry/model"

type BillRequest struct {
	Id          string             `json:"id"`
	Customer    string             `json:"customerId"`
	User        string             `json:"userId"`
	BillDetails []model.BillDetail `json:"billDetails"`
}