package models

import "bitcoin-rate/db"

type Email struct {
	Email string
}

func (e *Email) Save() error {
	result := db.DB.Create(e)

	return result.Error
}

func GetAllEmails() ([]Email, error) {
	var emails []Email
	result := db.DB.Select("email").Find(&emails)

	return emails, result.Error
}
