package use_module

import (
	"fmt"
	"strings"
)

// func multiply(a int, b int) int {
// 	return a * b
// }

// func lenAndUpper(name string) (int, string) {
// 	return len(name), strings.ToUpper(name)
// }

func lenAndUpper(name string) (lenght int, uppercase string) {
	defer fmt.Println("I'm done") // 함수 호출 끝나면 실행
	lenght = len(name)
	uppercase = strings.ToUpper(name)
	return
}

// func repeatMe(words ...string) {
// 	for i := 0; i < len(words); i++ {
// 		fmt.Println(words[i])
// 	}
// 	fmt.Println(words)
// }

func main1() {
	totalLenght, up := lenAndUpper("hyojun")
	fmt.Println(totalLenght, up)
	// repeatMe("hyo", "jun", "man")
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
