package app

import (
	"goreact2025/internal/response"
	"net/http"
)

func GetStart(w http.ResponseWriter, r *http.Request) {
	response.Success(w, "Hello, World!")
}
