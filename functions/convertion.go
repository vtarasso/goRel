package goreloaded

import (
	"strconv"
)

// Функция конвертации
func Convert(arr []string) []string {
	for i := 0; i < len(arr); i++ {
		word := arr[i]
		switch word {
		case "(up)":
			if i != 0 {
				arr[i-1] = ToUp(arr[i-1])
			}
			arr = append(arr[:i], arr[i+1:]...)
			i--
		case "(low)":
			if i != 0 {
				arr[i-1] = ToLow(arr[i-1])
			}
			arr = append(arr[:i], arr[i+1:]...)
			i--
		case "(cap)":
			cap := arr[i-1]
			if cap[len(cap)-2:] == "'t" {
				arr[i-1] = ToLow(arr[i-1])
				arr = append(arr[:i], arr[i+1:]...)
				if arr[i-1] == ToLow(arr[i-1]) {
					arr[i-1] = ToCap(arr[i-1][0:1]) + arr[i-1][1:]
				}
			} else if i != 0 {
				arr[i-1] = ToCap(arr[i-1])
			}
			arr = append(arr[:i], arr[i+1:]...)
			i--
		case "(hex)":
			if i != 0 {
				arr[i-1] = ToHex(arr[i-1])
			}
			arr = append(arr[:i], arr[i+1:]...)
			i--
		case "(bin)":
			if i != 0 {
				arr[i-1] = ToBin(arr[i-1])
			}
			arr = append(arr[:i], arr[i+1:]...)
			i--
		// Проверка после запятой, отработка не только на однозначные числа
		case "(up,":
			b := arr[i+1][:len(arr[i+1])-1]
			number, err := strconv.Atoi(b)
			if err != nil || number <= 0 {
				return nil
			}
			if number >= i {
				for a := 0; a < i-1; a++ {
					arr[a] = ToUp(arr[a])
				}
			} else {
				for j := 1; j <= number; j++ {
					arr[i-j] = ToUp(arr[i-j])
				}
			}
			arr = append(arr[:i], arr[i+2:]...)
			i--
		case "(low,":
			b := arr[i+1][:len(arr[i+1])-1]
			number, err := strconv.Atoi(b)
			if err != nil || number <= 0 {
				return nil
			}
			if number >= i {
				for a := 0; a < i-1; a++ {
					arr[a] = ToLow(arr[a])
				}
			} else {
				for j := 1; j <= number; j++ {
					arr[i-j] = ToLow(arr[i-j])
				}
			}
			arr = append(arr[:i], arr[i+2:]...)
			i--
		case "(cap,":
			b := arr[i+1][:len(arr[i+1])-1]
			number, err := strconv.Atoi(b)
			if err != nil || number <= 0 {
				return nil
			}
			if number >= i {
				for a := 0; a < i-1; a++ {
					arr[a] = ToCap(arr[a])
				}
			} else {
				for j := 1; j <= number; j++ {
					arr[i-j] = ToCap(arr[i-j])
				}
			}
			arr = append(arr[:i], arr[i+2:]...)
			i--
		}
		// Проверка артиклей
		for i, articles := range arr {
			if i < len(arr)-1 && (articles == "a" || articles == "A") {
				if Letter(string(arr[i+1][0])) {
					arr[i] += "n"
				}
			}
		}
	}
	return arr
}
