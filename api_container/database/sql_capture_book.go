package database

import (
	"encoding/json"
	"log"
)

type Response struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

func CaptureBook(author, title, isbn, category string) ([]byte, error) {
	log.Println("Started: executing sql statement", author, title, isbn, category)
	statement := `
        INSERT INTO books (author, title, isbn, category)
        VALUES ($1, $2, $3, $4) RETURNING id`

	var id string
	err := db.QueryRow(statement, author, title, isbn, category).Scan(&id)
	if err != nil {
		log.Printf("Error: Failed to execute SQL statement: %v", err)
		return nil, err
	}

	response := Response{
		Success: true,
		Message: id,
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		log.Printf("Error: Error marshaling JSON response: %v", err)
		return nil, err
	}
	log.Println("End: executed sql statement successful", author, title, isbn, category)
	return jsonResponse, nil
}
