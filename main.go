package main

import (
	"fmt"
	"os"
	"strings"

	funcs "goreloaded/functions"
)

func main() {
	if len(os.Args) == 3 {
		// Приходящие аргументы(sample/result)
		input := os.Args[1]
		output := os.Args[2]

		if input[len(input)-4:] == ".txt" && output[len(output)-4:] == ".txt" {
			// Чтение файла(input)
			file, err := os.ReadFile(input)
			if err != nil {
				fmt.Println("ERROR: Cannot READ this file!")
				os.Exit(1)
			}
			res := ""
			// Заходит сюда если файл имеет содержимое || записывает пустотой стринг от 'res' в файл
			if len(file) != 0 {
				text := funcs.Split(string(file))
				converting := funcs.Convert(text)
				if converting == nil {
					fmt.Println("ERROR: Cannot convert this!")
				}
				pointing := funcs.Punctuation(converting)
				// Join - объединяет элементы(слова) обратно, для создания единой строки.
				res = strings.Join((pointing), " ")
			}
			// Записываем полученные данные в наш "result.txt"
			err = os.WriteFile(output, []byte(res), 0o644)
			if err != nil {
				fmt.Println("ERROR: Cannot WRITE this file!")
				return
			}

		} else {
			fmt.Println("ERROR: Please write correct files format -> Ex: go run . example.txt example.txt")
		}

	} else {
		fmt.Println("ERROR: Incorrect number of arguments!")
	}
}
