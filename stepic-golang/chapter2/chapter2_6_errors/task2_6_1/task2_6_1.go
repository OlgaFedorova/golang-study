package main

import (
	"errors"
	"fmt"
)

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	res, err := divide(a, b)
	if err != nil {
		fmt.Println("ошибка")
	} else {
		fmt.Println(res)
	}
}

func divide(a int, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("Try input others numbers")
	} else {
		return a / b, nil
	}
}
