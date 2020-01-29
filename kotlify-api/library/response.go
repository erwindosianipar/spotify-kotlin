package library

import (
	"encoding/json"
	"net/http"
)

// HandleSuccess is for handling output json success
func HandleSuccess(resp http.ResponseWriter, data interface{}) {
	jsonData, err := json.Marshal(data)

	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write([]byte("Ooops, Something Went Wrong. Check Your Request Body."))
	}

	resp.Header().Set("Content-Type", "application/json")
	resp.Write(jsonData)
}

// HandleError is for handling output json success
func HandleError(resp http.ResponseWriter, message string) {
	jsonData, err := json.Marshal(message)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write([]byte("Ooops, Something Went Wrong. Check Your Request Body."))
	}

	resp.Header().Set("Content-Type", "application/json")
	resp.Write(jsonData)
}
