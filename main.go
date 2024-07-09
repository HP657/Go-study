package main

import "fmt"

// func multiply(a int, b int) int {
// 	return a * b
// }

// func lenAndUpper(name string) (int, string) {
// 	return len(name), strings.ToUpper(name)
// }

func repeatMe(words ...string) {
	for i := 0; i < len(words); i++ {
		fmt.Println(words[i])
	}
	fmt.Println(words)
}

func main() {
	repeatMe("hyo", "jun", "man")
	// totalLengtht, upperName := lenAndUpper("HP657")
	// fmt.Println(totalLengtht, upperName)
	// fmt.Println(multiply(2, 4))
	// name := "hyojun"
	// p := false
	// var name string = "hyojun"
	// name = "leehyojun"
	// fmt.Println(name)
	// fmt.Println(!p)
}
