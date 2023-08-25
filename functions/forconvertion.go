package goreloaded

import (
	"strconv"
	"strings"
)

// Функции для конвертации
func ToUp(str string) string {
	return strings.ToUpper(str)
}

func ToLow(str string) string {
	return strings.ToLower(str)
}

func ToCap(str string) string {
	return strings.Title(str)
}

func ToHex(str string) string {
	var num int64
	var err error
	num, err = strconv.ParseInt(str, 16, 64)
	if err != nil {
		return "'ERROR: Cannot convert to Hex!'"
	}
	numHex := strconv.Itoa(int(num))
	return numHex
}

func ToBin(str string) string {
	var num int64
	var err error
	num, err = strconv.ParseInt(str, 2, 64)
	if err != nil {
		return "'ERROR: Cannot convert to Binary!'"
	}
	numBin := strconv.Itoa(int(num))
	return numBin
}

func Letter(str string) bool {
	switch str {
	case "a", "e", "i", "o", "u", "y", "h":
		return true
	}
	return false
}
