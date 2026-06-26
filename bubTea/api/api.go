package api

import (
	"encoding/json"
	"net/http"
)

type ChoicesParam struct {
	Language string
	Username string
}

type ChoicesResponse struct {
	//margin for success status 200 etc
	Code int

	Information string
}

type Error struct {
	//margin for error status
	Code int

	//error message
	Message string
}

func writeError(w http.ResponseWriter, message string, code int) {
	resp := Error{
		Code:    code,
		Message: message,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	json.NewEncoder(w).Encode(resp)

}

// wrapper for we func
var (
	RequestErrorHandler = func(w http.ResponseWriter, err error) {
		writeError(w, err.Error(), http.StatusBadRequest)
	}
	InternalErrorHandler = func(w http.ResponseWriter) {
		writeError(w, "unexpected error", http.StatusInternalServerError)
	}
)

type SubmitChoiceRequest struct {
	Username string `json:"username"`
	Choice   string `json:"choice"`
}
