package inventory

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/derchrischkya/libary/database"
	"github.com/julienschmidt/httprouter"
)

type CaptureBookRequestBody struct {
	Title    string `json:"title"`
	Author   string `json:"author"`
	Category string `json:"category"`
	Isbn     string `json:"isbn"`
}

type CaptureBookResponseBody struct {
	ID string `json:"id"`
}

type Response struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

func CaptureBook() httprouter.Handle {
	return func(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
		log.Println("Started: running endpoint to capture a new book")

		// Read and parse the request body
		var requestBody CaptureBookRequestBody
		err := json.NewDecoder(request.Body).Decode(&requestBody)
		if err != nil {
			http.Error(writer, "Invalid JSON request body", http.StatusBadRequest)
			return
		}

		// Execute the prepared SQL statement with the book details
		responseBytes, err := database.CaptureBook(requestBody.Author, requestBody.Title, requestBody.Isbn, requestBody.Category)
		if err != nil {
			http.Error(writer, "Failed to insert book into the database", http.StatusInternalServerError)
			log.Printf("Failed to insert book into the database: %v", err)
			return
		}

		// Unmarshal JSON response into a struct
		var insertResponse Response
		err = json.Unmarshal(responseBytes, &insertResponse)
		if err != nil {
			http.Error(writer, "Failed to insert book into the database", http.StatusInternalServerError)
			fmt.Println("Error unmarshaling JSON response:", err)
			return
		}

		// Prepare the response with the captured book ID
		responseBody := CaptureBookResponseBody{
			ID: insertResponse.Message,
		}

		// Marshal the response body to JSON
		responseBodyBytes, err := json.Marshal(responseBody)
		if err != nil {
			http.Error(writer, "Failed to insert book into the database", http.StatusInternalServerError)
			log.Printf("Failed to marshal response: %v", err)
			return
		}

		log.Println("End: finished endpoint to capture a new book", insertResponse.Message)
		// Write the response
		writer.WriteHeader(http.StatusOK)
		writer.Header().Set("Content-Type", "application/json")
		writer.Write(responseBodyBytes)
	}
}
