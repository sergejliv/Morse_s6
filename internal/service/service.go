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
	// Очищаем строку от пробелов по краям
	trimmed := strings.TrimSpace(s)
	if trimmed == "" {
		return false
	}

	// Если есть ЛЮБЫЕ буквы (русские, английские) - это текст, а не Morse
	for _, char := range trimmed {
		if unicode.IsLetter(char) {
			return false
		}
	}

	// Если есть цифры - это тоже не Morse (в контексте этого задания)
	for _, char := range trimmed {
		if char >= '0' && char <= '9' {
			return false
		}
	}

	// Проверяем, что строка содержит ТОЛЬКО допустимые символы Морзе
	for _, char := range trimmed {
		// Допустимые символы: точка, тире
		if char != '.' && char != '-' {
			return false
		}
	}

	// Дополнительная проверка: должна быть хотя бы одна точка или тире
	hasDotsOrDashes := strings.Contains(trimmed, ".") || strings.Contains(trimmed, "-")
	if !hasDotsOrDashes {
		return false
	}

	return true
}
