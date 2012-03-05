// this row col multiplier is completely serial and performs i/o through channels
package main

import (
	"fmt"
	"flag"
)

var sliceSize int

func init() {
	flag.IntVar(&sliceSize,"s",5,"specifies the size of slices to be multiplied")
}

func RowColMultiplier(row,col chan []int, val chan int) {
	a := <- row
	b := <- col
	sum:=0
	for i:=0;i<len(a);i++ {
		sum += a[i]*b[i]
	}
	val <- sum
}

func main() {//run this to check
	flag.Parse()// must be called before flags are used
	a,b := make([]int, sliceSize),make([]int, sliceSize)
	fmt.Println("Creating slices of len ", sliceSize)
	for i:=0;i<sliceSize;i++ {//creating 2 slices containing only 1
		a[i],b[i]=1,1
	}
	row,col := make(chan []int, 1),make(chan []int, 1)
	res := make(chan int, 1)
	row <- a
	col <- b
	RowColMultiplier(row, col, res)
	fmt.Println("The result is", <-res)//should be same as length if the slices are multiplied correctly
}
