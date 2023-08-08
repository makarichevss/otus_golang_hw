package hw02unpackstring

import (
	"errors"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(str string) (string, error) {
	var builder strings.Builder

	// Кидает ошибку, если строка начинается с цифры или содержит числа больше 9
	if len(str) > 0 && (unicode.IsDigit(rune(str[0])) || hasConsecutiveDigits(str)) {
		return "", ErrInvalidString
	}

	for i := 0; i < len(str); {
		r := rune(str[i])
		i++

		// Если следующая руна является цифрой, она используется как множитель
		if i < len(str) && unicode.IsDigit(rune(str[i])) {
			// Преобразовать байт цифры в int
			count := int(str[i] - '0')
			builder.WriteString(strings.Repeat(string(r), count))
			// пропустить цифру
			i++
		} else {
			builder.WriteRune(r)
		}
	}

	return builder.String(), nil
}

// Проверяет, есть ли в строке две или более подряд идущие цифры
func hasConsecutiveDigits(s string) bool {
	for i := 0; i < len(s)-1; i++ {
		if unicode.IsDigit(rune(s[i])) && unicode.IsDigit(rune(s[i+1])) {
			return true
		}
	}

	return false
}
