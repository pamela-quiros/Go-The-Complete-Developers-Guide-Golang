package main

import "fmt"

func main() {
	numbers := []int{12, 7, 5, 4, 9, 8}

	for _, number := range numbers {

		if number%2 == 0 {
			fmt.Println(number, "is even")
		} else {
			fmt.Println(number, "is odd")
		}
	}
}
