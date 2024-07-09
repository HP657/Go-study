package use_module

import (
	"fmt"

	"github.com/HP657/go-study/use_module/something"
)

func main() {
	fmt.Println("hi")
	something.SayHello()
	something.SayBye()
}
