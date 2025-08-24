package server

import (
	"log"
	"net/http"
	"time"
)

// Server структура для HTTP сервера
type Server struct {
	logger *log.Logger
	server *http.Server
}

// New создает и настраивает новый HTTP сервер
func New(logger *log.Logger) *Server {
	// Создаем роутер
	router := createRouter()

	// Создаем и настраиваем HTTP сервер
	httpServer := &http.Server{
		Addr:         ":8080",          // Порт 8080
		Handler:      router,           // Наш роутер
		ErrorLog:     logger,           // Логгер для ошибок сервера
		ReadTimeout:  5 * time.Second,  // Таймаут чтения - 5 секунд
		WriteTimeout: 10 * time.Second, // Таймаут записи - 10 секунд
		IdleTimeout:  15 * time.Second, // Таймаут простоя - 15 секунд
	}

	return &Server{
		logger: logger,
		server: httpServer,
	}
}

// createRouter создает HTTP роутер и регистрирует хендлеры
func createRouter() *http.ServeMux {
	router := http.NewServeMux()

	// Здесь будут регистрироваться хендлеры

	return router
}

// Start запускает HTTP сервер
func (s *Server) Start() error {
	s.logger.Printf("Starting server on %s", s.server.Addr)
	return s.server.ListenAndServe()
}

// Shutdown gracefully останавливает сервер
func (s *Server) Shutdown() error {
	s.logger.Println("Shutting down server...")

	return nil
}
