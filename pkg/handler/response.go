package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"payments/pkg/appError"
)

const (
	headerContentType = "Content-Type"
	contentTypeJSON   = "application/json"
)

func JSON(w http.ResponseWriter, r *http.Request, statusCode int, body interface{}) {
	w.Header().Set(headerContentType, contentTypeJSON)
	w.WriteHeader(statusCode)

	log.Println("response: ", body)

	if err := json.NewEncoder(w).Encode(body); err != nil {
		log.Printf("failed to write response: %v\n", err)
	}

	if isSuccess(statusCode) {
		log.Println("http.request.success")
	} else {
		log.Println("http.request.error")
	}
}

func JSONWithError(w http.ResponseWriter, r *http.Request, err *appError.Error) {
	JSON(w, r, err.Code, err)
}

func isSuccess(statusCode int) bool {
	if statusCode >= 200 && statusCode < 400 {
		return true
	}

	return false
}
