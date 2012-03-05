package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 100; i++ {
		go test(i)
	}
	fmt.Println("Finished")
}

func test(a int) {
	fmt.Println(a)
}

