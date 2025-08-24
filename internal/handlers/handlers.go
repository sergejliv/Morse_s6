package handlers

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

// RootHandler обрабатывает корневой эндпоинт /

func RootHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	http.ServeFile(w, r, "/index.html")

}

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Получаем файл из формы
	file, header, err := r.FormFile("myFile")
	if err != nil {
		http.Error(w, fmt.Sprintf("Error retrieving file: %v", err), http.StatusInternalServerError)
		return
	}
	defer file.Close()

	// Читаем данные из файла
	fileData, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error reading file: %v", err), http.StatusInternalServerError)
		return
	}

	// Передаем данные в функцию автоопределения из пакета service
	convertedData, err := service.Convert(string(fileData))
	if err != nil {
		http.Error(w, fmt.Sprintf("Error converting data: %v", err), http.StatusInternalServerError)
		return
	}

	// Создаем локальный файл для результата
	fileName := generateFileName(header.Filename)
	err = os.WriteFile(fileName, []byte(convertedData), 0644)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error writing result file: %v", err), http.StatusInternalServerError)
		return
	}

	// Возвращаем результат конвертации
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "File converted successfully!\n")
	fmt.Fprintf(w, "Original file: %s\n", header.Filename)
	fmt.Fprintf(w, "Result saved to: %s\n", fileName)
	fmt.Fprintf(w, "\nConverted content:\n%s", convertedData)
}

// generateFileName генерирует уникальное имя файла на основе времени и оригинального расширения
func generateFileName(originalName string) string {
	// Получаем расширение оригинального файла
	ext := filepath.Ext(originalName)
	if ext == "" {
		ext = ".txt" // расширение по умолчанию
	}

	// Генерируем уникальное имя на основе времени
	timestamp := time.Now().UTC().Format("20060102_150405")
	return fmt.Sprintf("converted_%s%s", timestamp, ext)
}
