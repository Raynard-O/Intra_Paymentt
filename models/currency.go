package models

import "github.com/jinzhu/gorm"

type Request struct {
	gorm.DB
	ID 				  string	`json:"id"`
	UserId            int16  	`json:"userId"`
	CurrencyAvailable string 	`json:"currencyAvailable"`
	CurrencyRequested string 	`json:"currencyRequested"`
	CreatedOn 		  string 	`json:"createdOn"`
}


