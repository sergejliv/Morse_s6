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

func isMorseCode(s string) bool {
	trimmed := strings.TrimSpace(s)
	if trimmed == "" {
		return false
	}

	// Если есть любые буквы - это текст
	if strings.ContainsFunc(trimmed, unicode.IsLetter) {
		return false
	}

	// Проверяем, что нет запрещенных символов
	hasInvalidChars := strings.ContainsFunc(trimmed, func(r rune) bool {
		return r != '.' && r != '-' && r != ' ' && r != '/' && !unicode.IsSpace(r)
	})

	if hasInvalidChars {
		return false
	}

	// Должна быть хотя бы одна точка или тире
	return strings.Contains(trimmed, ".") || strings.Contains(trimmed, "-")
}
