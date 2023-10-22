package main

import "fmt"

func printLogF(s string) {
	fmt.Printf(s)
}

func printLog(s string) {
	fmt.Println(s)
}

func main() {
	values := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for _, num := range values {

		if (num % 2) == 0 {
			fmt.Printf("\n\n Number %v is EVEN", num)
		} else {
			fmt.Printf("\n\n Number %d is ODD", num)
		}
	}
}
