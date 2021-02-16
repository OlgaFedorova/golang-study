package main

import "fmt"

func main() {
	var x1, x2 int
	fmt.Scan(&x1, &x2)
	test(&x1, &x2)
}

func test(x1 *int, x2 *int) {
	tmp := *x1
	*x1 = *x2
	*x2 = tmp
	fmt.Println(*x1, *x2)
}
