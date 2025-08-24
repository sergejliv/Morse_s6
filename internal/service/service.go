package service

import (
	"errors"
	"strings"

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
func isMorseCode(s string) bool {
	// Очищаем строку от пробелов по краям
	trimmed := strings.TrimSpace(s)
	if trimmed == "" {
		return false
	}

	// Проверяем, что строка содержит только допустимые символы Морзе
	for _, char := range trimmed {
		// Допустимые символы: точка, тире, пробел, слэш
		if char != '.' && char != '-' && char != ' ' && char != '/' {
			return false
		}
	}

	return true
}
