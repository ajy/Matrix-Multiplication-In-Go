package main

import (
	"fmt"
)

func main() {
	a := make([][]int,10)
	a[0] = []int{1,2,3,4,5}
	a[1] = []int{2,3,4,5,6,7,8,9,0}
	fmt.Println(a[0][:]) 
}

