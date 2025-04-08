package main

import (
	"fmt"
	"goreact2025/api/route"
	"goreact2025/internal/httpserver"

	"github.com/joho/godotenv"
)

func main() {
	// Загружаем переменные окружения из файла .env
	errEnv := godotenv.Load()
	if errEnv != nil {
		fmt.Println("No .env file found. Using environment variables directly.")
	}
	// Создаем новый HTTP-сервер на порту 8080
	server := httpserver.NewHTTPServer("8080")

	// Регистрируем маршруты
	route.InitRoutes(server)

	// Запускаем сервер
	err := server.Start()
	if err != nil {
		panic(err)
	}
}
