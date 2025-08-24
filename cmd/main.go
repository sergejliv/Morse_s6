package main

import (
	"log"
	"os"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/server"
)

func main() {
	// Создаем логгер с префиксом и настройками форматирования
	logger := log.New(os.Stdout, "APP: ", log.LstdFlags|log.Lshortfile)

	// Создаем сервер с помощью функции из пакета server
	server := server.New(logger)

	// Запускаем сервер
	logger.Println("Starting server on port 8080...")
	if err := server.Start(); err != nil {
		// Если возникла ошибка при запуске сервера, выводим её на уровне Fatal
		logger.Fatalf("Failed to start server: %v", err)
	}
}
