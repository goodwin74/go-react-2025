package test

import (
	"goreact2025/internal/response"
	"net/http"
)

func TestHandlerController(w http.ResponseWriter, r *http.Request) {
	response.Success(w, "Hello, World!")
}
