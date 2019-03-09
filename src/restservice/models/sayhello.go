package models
import (
	u "restservice/utils"
)

// STRUCT to store account information
type Account struct {
	Name     string `json:"name"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
}

func (account *Account) SayHi() map[string]interface{}{
	response := u.Message(true, "success")
	response["data"] = account
	return response
}