package goreloaded

import "strings"

// Функция split - разбивает строку на отдельные элементы(слова), по пробелам
func Split(str string) []string {
	var text []string
	text = strings.Split(str, " ")
	var arr []string
	for _, res := range text {
		if len(res) != 0 {
			arr = append(arr, res)
		}
	}
	return arr
}
