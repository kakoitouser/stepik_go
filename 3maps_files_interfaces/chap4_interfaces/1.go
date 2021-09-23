package main

import (
	//"encoding/json" // пакет используется для проверки ответа, не удаляйте его
	"fmt" // пакет используется для проверки ответа, не удаляйте его
	//"os"            // пакет используется для проверки ответа, не удаляйте его
)

func main() {
	value1, value2, operation := readTask() // исходные данные получаются с помощью этой функции
	switch value1.(type) {
	case float64:
		switch value2.(type) {
		case float64:
			switch operation.(type) {
			case string:
				switch operation {
				case "+":
					fmt.Printf("%.4f", (value1.(float64) + value2.(float64)))
				case "-":
					fmt.Printf("%.4f", (value1.(float64) - value2.(float64)))
				case "*":
					fmt.Printf("%.4f", (value1.(float64) * value2.(float64)))
				case "/":
					fmt.Printf("%.4f", (value1.(float64) / value2.(float64)))
				default:
					err := "неизвестная операция"
					fmt.Println(err)
				}
			default:
				err := "неизвестная операция"
				fmt.Println(err)
			}
		default:
			errStr := fmt.Sprintf("value=%v: %T", value2, value2)
			fmt.Println(errStr)
		}
	default:
		errStr := fmt.Sprintf("value=%v: %T", value1, value1)
		fmt.Println(errStr)
	} // все полученные значения имеют тип пустого интерфейса
}
