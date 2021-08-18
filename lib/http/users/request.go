package users

type UserRequest struct {
	UserId            int16  `json:"userId"`
	CurrencyAvailable string `json:"currencyAvailabe"`
	CurrencyRequested string `json:"currencyRequested"`
}
