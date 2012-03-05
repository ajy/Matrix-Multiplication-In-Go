package main

import (
	"fmt"
	"flag"
	"os"
)

var BlockSize int
var sliceSize int

func init() {
	flag.IntVar(&BlockSize,"b",5,"specifies the size of blocks that will be run concurrently")
	flag.IntVar(&sliceSize,"s",5,"specifies the size of slices to be multiplied")
}

func RowColMultiplier(a,b []int) (n int, err os.Error) {
	if(len(a) != len(b)) {
		return 0, os.NewError("Slices not of same size")
	}
	val := make(chan int)
	aLen:=len(a)
	sum:=0
	j := 0
	for i := aLen/BlockSize; i>0; i-- {//breaks the slice into blocks that can be evaluated concurrently
		go rowColMultiplierDrone(a[j:j+BlockSize], b[j:j+BlockSize], val)
		j += BlockSize
	}
	if((aLen%BlockSize) != 0) {//incase it can't be broken perfectly into blocks
		go rowColMultiplierDrone(a[j:], b[j:], val)
		j+=BlockSize
	}
	j/=BlockSize//j now counts the number of goroutines launched
	for ;; {
		sum += <-val
		j--
		if (j == 0) {break}
	}
	return sum, nil
}

func rowColMultiplierDrone(a,b []int, val chan int) {
	sum:=0
	for i:=0;i<len(a);i++ {
		sum += a[i]*b[i]
	}
	val <- sum
}

func main() {//run this to check
	flag.Parse()// must be called before flags are used
	a,b := make([]int, sliceSize),make([]int, sliceSize)
	fmt.Println("Creating slices of len", sliceSize)
	for i:=0;i<sliceSize;i++ {//creating 2 slices containing only 1
		a[i],b[i]=1,1
	}
	res,_ := RowColMultiplier(a,b)
	fmt.Println("The result is", res)//should be same as length if the slices are multiplied correctly
}
