package response

import (
    "encoding/json"
    "net/http"
)

// Структура для формирования JSON-ответа.
type Response struct {
    Status string      `json:"status"`
    Result interface{} `json:"result"`
}

// Возвращает успешный JSON-ответ.
func Success(w http.ResponseWriter, result interface{}) {
    response := Response{
        Status: "success",
        Result: result,
    }
    sendResponse(w, http.StatusOK, response)
}

// Возвращает JSON-ответ с ошибкой.
func Error(w http.ResponseWriter, statusCode int, result interface{}) {
    response := Response{
        Status: "error",
        Result: result,
    }
    sendResponse(w, statusCode, response)
}

// Отправляет JSON-ответ.
func sendResponse(w http.ResponseWriter, statusCode int, response Response) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(statusCode)

    if err := json.NewEncoder(w).Encode(response); err != nil {
        http.Error(w, "Failed to encode response", http.StatusInternalServerError)
    }
}