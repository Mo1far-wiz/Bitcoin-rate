package models

import "bitcoin-rate/db"

type Email struct {
	ID    int64
	Email string `binding:"required"`
}

func (e *Email) Save() error {
	query := "INSERT INTO emails(email) VALUES ($1) RETURNING id;"

	err := db.DB.QueryRow(query, e.Email).Scan(&e.ID)

	return err
}

func GetAllEmails() ([]Email, error) {
	query := "SELECT * FROM emails;"
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var emails []Email
	for rows.Next() {
		var email Email
		err := rows.Scan(&email.ID, &email.Email)
		if err != nil {
			return nil, err
		}

		emails = append(emails, email)
	}

	return emails, nil
}
