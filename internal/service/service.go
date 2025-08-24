package service

import (
	"errors"
	"strings"
	"unicode"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

// Convert автоматически определяет тип входной строки и вызывает соответствующую функцию
func Convert(input string) (string, error) {
	if input == "" {
		return "", errors.New("пустая входная строка")
	}

	// Определяем, является ли строка кодом Морзе
	if isMorseCode(input) {
		// Вызываем функцию конвертации Морзе в текст
		return morse.ToText(input), nil
	} else {
		// Вызываем функцию конвертации текста в Морзе
		return morse.ToMorse(input), nil
	}
}

// isMorseCode определяет, является ли строка кодом Морзе
// service/converter.go
func isMorseCode(s string) bool {
	trimmed := strings.TrimSpace(s)
	if trimmed == "" {
		return false
	}

	// Проверить на наличие букв - если найдены, это текст
	for _, char := range trimmed {
		if unicode.IsLetter(char) {
			return false
		}
	}

	// Проверить, содержит ли только допустимые символы Морзе
	validChars := ". -/"
	for _, char := range trimmed {
		if !strings.ContainsRune(validChars, char) && !unicode.IsSpace(char) {
			return false
		}
	}

	// Должен содержать хотя бы одну точку или тире
	hasMorseChars := strings.Contains(trimmed, ".") || strings.Contains(trimmed, "-")
	if !hasMorseChars {
		return false
	}

	return true
}
