package service

import (
	"errors"
	"strings"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

// convert определяет морзе или обычный текст - конвертирует в противоположное
func Convert(input string) (string, error) {
	// удалим пробелы и переносы строк по краям
	trimmed := strings.TrimSpace(input)

	if trimmed == "" {
		return "", errors.New("входная строка пуста")
	}

	// если строка содердит только точки и тире, пробелы и слэши - это морзе
	if isMorse(trimmed) {
		// конвертируем Морзе в текст
		return morse.ToText(trimmed), nil
	}

	// иначе конвертируем текст в Морзе
	return morse.ToMorse(trimmed), nil
}

// isMorse проверяет, является ли строка кодом Морзе
func isMorse(s string) bool {
	for _, r := range s {
		if r != '.' && r != '-' && r != ' ' && r != '/' && r != '\n' && r != '\r' {
			return false
		}
	}
	return true
}
