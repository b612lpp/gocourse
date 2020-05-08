package main

import (
	"fmt"
	"gocourse/homework/homework-04/help/calculator"
	"io/ioutil"
)

func main() {
	input := ""
	for {
		fmt.Print("> ")
		if _, err := fmt.Scanln(&input); err != nil {
			fmt.Println(err)
			continue
		}

		if input == "exit" {
			break
		}

		if input == "help" {

			data, err := ioutil.ReadFile("help.txt")
			if err != nil {
				fmt.Println("File reading error", err)
				return
			}
			fmt.Println("Here is About calculator", string(data))
		}

		if res, err := calculator.Calculate(input); err == nil {
			fmt.Printf("Результат: %v\n", res)
		} else {
			fmt.Println("Не удалось произвести вычисление")
		}
	}
}
