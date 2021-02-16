package main

import "fmt"

func main() {
	workArray := [10]uint8{}
	for i := 0; i < len(workArray); i++ {
		fmt.Scan(&workArray[i])
	}

	var idx1, idx2 int
	for i := 1; i <= 3; i++ {
		fmt.Scan(&idx1)
		fmt.Scan(&idx2)
		tmp := workArray[idx1]
		workArray[idx1] = workArray[idx2]
		workArray[idx2] = tmp
	}
	for _, elem := range workArray {
		fmt.Print(elem, " ")
	}
}
