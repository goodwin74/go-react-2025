package route

import (
	app "goreact2025/api/controllers"
	"goreact2025/internal/httpserver"
	"goreact2025/internal/response"
	"net/http"
)

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	// Возвращаем успешный ответ
	response.Success(w, []string{"This", "is", "the", "about", "page"})
}

func errorHandler(w http.ResponseWriter, r *http.Request) {
	// Возвращаем ошибку
	response.Error(w, http.StatusBadRequest, map[string]string{"error": "Invalid request"})
}
func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	// Возвращаем ошибку 404
	response.Error(w, http.StatusNotFound, map[string]string{"error": "404 Not Found"})
}

func InitRoutes(s *httpserver.HTTPServer) {
	s.AddReactRouter()
	s.AddRoute("/api/start", app.GetStart)
	s.AddRoute("/api/error", errorHandler)

	// Добавляем обработчик для неизвестных маршрутов
	s.SetNotFoundHandler(notFoundHandler)
}
