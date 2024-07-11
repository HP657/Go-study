package main

import "fmt"

func supperAdd(numbers ...int) int {
	sum := 0
	for i := 0; i < len(numbers); i++ {
		fmt.Println(numbers[i])
		sum += numbers[i]
	}
	// for index, numbers := range numbers {
	// 	fmt.Println(index, numbers)
	// 	sum += numbers
	// }
	return sum
}

func main() {
	total := supperAdd(1, 2, 3, 4, 5, 6, 7, 8, 9)
	fmt.Print(total)
}
