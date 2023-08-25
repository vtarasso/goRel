package goreloaded

// Функция расстановки знаков препинания
func Punctuation(arr []string) []string {
	marks := []string{",", ".", "!", "?", ":", ";"}
	for i := 1; i < len(arr); i++ {
		word := arr[i]
		for _, pointing := range marks {
			if i != 0 && (string(word[0]) == pointing) && (string(word[len(word)-1]) != pointing) {
				arr[i-1] = arr[i-1] + pointing
				arr[i] = word[1:]
				i--
			}
		}
		for i := 1; i < len(arr); i++ {
			word := arr[i]
			for _, pointing := range marks {
				if (string(word[0]) == pointing) && (string(word[len(word)-1]) == pointing) && (arr[i] != arr[len(arr)-1]) {
					arr[i-1] = arr[i-1] + word
					arr = append(arr[:i], arr[i+1:]...)
					i--
				}
			}
		}
		for i, word := range arr {
			for _, pointing := range marks {
				if (string(word[0]) == pointing) && (arr[len(arr)-1] == arr[i]) {
					arr[i-1] = arr[i-1] + word
					arr = arr[:len(arr)-1]
				}
			}
		}
		num1 := 0
		for i := 0; i < len(arr); i++ {
			word := arr[i]
			if string(word[0]) == string(rune('\'')) {
				if num1%2 != 0 {
					if i != 0 {
						arr[i-1] = arr[i-1] + string(rune('\''))
					}
					if len(word) > 1 {
						arr[i] = arr[i][1:]
					} else {
						arr = append(arr[:i], arr[i+1:]...)
					}
				} else if num1%2 == 0 && len(word) == 1 {
					if i != len(arr)-1 {
						arr[i+1] = string(rune('\'')) + arr[i+1]
						arr = append(arr[:i], arr[i+1:]...)
						i++
					}
				}
				num1++
			}
			if len(word) > 1 && string(word[len(word)-1]) == string(rune('\'')) {
				if num1%2 == 0 {
					arr[i+1] = string(rune('\'')) + arr[i+1]
					arr[i] = arr[i][:len(arr[i])-1]
					i++
				}
				num1++
			}
		}
	}
	num2 := 0
	for i := 0; i < len(arr); i++ {
		word := arr[i]
		if string(word[0]) == string(rune('"')) {
			if num2%2 != 0 {
				if i != 0 {
					arr[i-1] = arr[i-1] + string(rune('"'))
				}
				if len(word) > 1 {
					arr[i] = arr[i][1:]
				} else {
					arr = append(arr[:i], arr[i+1:]...)
				}
			} else if num2%2 == 0 && len(word) == 1 {
				if i != len(arr)-1 {
					arr[i+1] = string(rune('"')) + arr[i+1]
					arr = append(arr[:i], arr[i+1:]...)
					i++
				}
			}
			num2++
		}
		if len(word) > 1 && string(word[len(word)-1]) == string(rune('"')) {
			if num2%2 == 0 {
				arr[i+1] = string(rune('"')) + arr[i+1]
				arr[i] = arr[i][:len(arr[i])-1]
				i++
			}
			num2++
		}
	}
	return arr
}
