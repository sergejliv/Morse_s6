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
func isMorseCode(s string) bool {
	//trimmed := strings.TrimSpace(s)
	if s == "" {
		return false
	}

	// Если есть ЛЮБЫЕ буквы (русские или английские) - это текст, а не Morse
	if strings.ContainsFunc(s, unicode.IsLetter) {
		return false
	}

	// Проверяем, что строка содержит только допустимые символы Морзе
	// Допустимые символы: точка, тире, пробел, слэш
	hasInvalidChars := strings.ContainsFunc(s, func(r rune) bool {
		return r != '.' && r != '-'
	})

	if hasInvalidChars {
		return false
	}

	// Дополнительная проверка: должна быть хотя бы одна точка или тире
	return strings.Contains(s, ".") || strings.Contains(s, "-")
}
