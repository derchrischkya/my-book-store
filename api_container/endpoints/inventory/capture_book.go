package inventory

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/julienschmidt/httprouter"
)

type CaptureBookRequestBody struct {
	Title  string `json:"title"`
	Author string `json:"author"`
}

type CaptureBookResponseBody struct {
	ID uuid.UUID `json:"id"`
}

func CaptureBook() httprouter.Handle {
	return func(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
		log.Println("Started: running endpoint to capure new book")
		// NON-PROD ready --> request.Body can be 10tb and crash server
		// PROD - Use diffrent reader with size limitation functionality
		requestBodyBytes, err := io.ReadAll(request.Body)
		if err != nil {
			// NON-PROD ready --> is return server error to customer!!!
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}

		var requestBody CaptureBookRequestBody
		err = json.Unmarshal(requestBodyBytes, &requestBody)

		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}

		id := uuid.New()
		log.Println("received capture-book", id, requestBody)

		// Logic...MinIO :)

		responseBody := CaptureBookResponseBody{
			ID: id,
		}

		responseBodyBytes, err := json.Marshal(responseBody)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}

		writer.WriteHeader(http.StatusOK)
		writer.Header().Set("Content-Type", "application/json")
		writer.Write(responseBodyBytes)
	}
}
