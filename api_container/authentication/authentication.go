package authentication

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/julienschmidt/httprouter"
)

type Data struct {
	Active    bool     `json:"active"`
	ClientID  string   `json:"client_id"`
	Sub       string   `json:"sub"`
	Exp       int      `json:"exp"`
	Iat       int      `json:"iat"`
	Nbf       int      `json:"nbf"`
	Aud       []string `json:"aud"`
	Iss       string   `json:"iss"`
	TokenType string   `json:"token_type"`
	TokenUse  string   `json:"token_use"`
}

const tokenIntrospectURL = "https://127.0.0.1:5445/oauth2/introspect"

type Response struct {
	Nested Data `json:"nested"`
}

func Authenticate(next httprouter.Handle) httprouter.Handle {
	return func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		log.Println("Started: checking authorization of API call")

		// Extract the token from the Authorization header
		token := strings.Replace(request.Header.Get("Authorization"), "Bearer ", "", 1)

		// Make a POST request to the token introspection endpoint
		resp, err := introspectToken(token)
		if err != nil {
			log.Println(err)
			writer.WriteHeader(http.StatusUnauthorized)
			return
		}
		defer resp.Body.Close()

		// Read the response body
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Println(err)
			writer.WriteHeader(http.StatusUnauthorized)
			return
		}

		// Parse the response JSON
		var response Data
		if err := json.Unmarshal(body, &response); err != nil {
			log.Println(err)
			writer.WriteHeader(http.StatusUnauthorized)
			return
		}

		// Check if the token is active
		if !response.Active {
			writer.WriteHeader(http.StatusUnauthorized)
			return
		}
		log.Println("End: successfully authenticated API call")

		// If authentication is successful, call the next handler
		next(writer, request, params)
	}
}

func introspectToken(token string) (*http.Response, error) {
	method := "POST"
	payload := fmt.Sprintf("token=%s", token)

	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	// Create a new HTTP client with timeout settings
	client := &http.Client{
		Timeout: time.Second * 10, // Set a reasonable timeout value
	}

	// Create a request to the token introspection endpoint
	req, err := http.NewRequest(method, tokenIntrospectURL, strings.NewReader(payload))
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	// Perform the HTTP request
	return client.Do(req)
}
