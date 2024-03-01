package main

import "fmt"

func avg(results ...int) int {
	var sum int
	for _, result := range results {
		sum = sum + result
	}

	return sum / len(results)
}

func main() {
	// Varidic function

	result := avg(7, 8, 9, 10, 10, 9)

	fmt.Println(result)
}
