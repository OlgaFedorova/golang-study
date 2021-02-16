package main

import "fmt"

func main() {
	fn := func(num uint) uint {
		var res uint
		digitRes := uint(1)
		for num != 0 {
			var temp uint
			if num < 10 {
				temp = num
				num = 0
			} else {
				temp = num % 10
				num = (num - temp) / 10
			}

			if temp != 0 && temp%2 == 0 {
				res = res + temp*digitRes
				digitRes = digitRes * 10
			}

		}
		if res == 0 {
			return uint(100)
		}
		return res
	}

	fmt.Println(fn(727178))
	fmt.Println(fn(7717))

}
