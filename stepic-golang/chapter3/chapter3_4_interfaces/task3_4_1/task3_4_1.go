package main

import (
	"errors"
	"fmt" // пакет используется для проверки ответа, не удаляйте его
)

func main() {
	value1, value2, operation := readTask() // исходные данные получаются с помощью этой функции

	checkNum := func(i interface{}) (float64, error) {
		switch v := i.(type) {
		case float64:
			return v, nil
		default:
			return 0, errors.New(fmt.Sprint("value=", v) + fmt.Sprintf(": %T", v))
		}
	}

	checkOper := func(i interface{}) (string, error) {
		if v, ok := i.(string); ok {
			if v == "+" || v == "-" || v == "*" || v == "/" {
				return v, nil
			}
		}

		return "", errors.New("неизвестная операция")
	}

	if v1, err := checkNum(value1); err != nil {
		fmt.Println(err)
	} else {
		if v2, err := checkNum(value2); err != nil {
			fmt.Println(err)
		} else {
			if oper, err := checkOper(operation); err != nil {
				fmt.Println(err)
			} else {
				if oper == "+" {
					fmt.Printf("%.4f", v1+v2)
				} else if oper == "-" {
					fmt.Printf("%.4f", v1-v2)
				} else if oper == "*" {
					fmt.Printf("%.4f", v1*v2)
				} else if oper == "/" {
					fmt.Printf("%.4f", v1/v2)
				}
			}
		}
	}
}

func readTask() (value1, value2, operation interface{}) {
	return 2.07, 3.06789, "/"
}
