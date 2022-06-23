package main

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func isValidStr(runesArray []rune) bool {
	var previousIsDigit bool
	for key, char := range runesArray {
		currentIsDigit := unicode.IsDigit(char)
		if key == 0 {
			if currentIsDigit {
				return false
			}
			previousIsDigit = currentIsDigit
			continue
		}

		if currentIsDigit && previousIsDigit {
			return false
		}

		previousIsDigit = currentIsDigit
	}

	return true
}

func Unpack(s string) (string, error) {
	stringBuilder := strings.Builder{}
	runesArray := []rune(s)
	if !isValidStr(runesArray) {
		return "", ErrInvalidString
	}

	var buf string
	for key, char := range runesArray {
		currentSymbol := string(char)
		if unicode.IsDigit(char) && buf != "" && buf != "\\" {
			countRepeat, _ := strconv.Atoi(currentSymbol)
			stringBuilder.WriteString(strings.Repeat(buf, countRepeat))
			continue
		}

		if len(runesArray) > key+1 && unicode.IsDigit(runesArray[key+1]) {
			buf = string(char)
			continue
		}

		stringBuilder.WriteRune(char)
	}

	return stringBuilder.String(), nil
}
