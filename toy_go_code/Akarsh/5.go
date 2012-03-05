package main

import	(
	"fmt"
)
var a = [4]int{1,2,3,4}
func main()	{
	b := []int{2,3,4,5}
	c:= []int{3,4,5,6}
	d := []int{2,6,4,5}
	ch :=make(chan int)
	//ch2 := make(chan int)
	//ch3 := make(chan int)
	//ch4 := make(chan int)
	go test(b,ch,1)
	go test(c,ch,2)
	go test(d,ch,3)
	go test(b,ch,4)
	for i:=0;i<4;i++	{
		select	{
			case a:=<-ch:	fmt.Println("Returned From ",a)
		}
	}
}

func test(list []int,ch chan int,n int){
	var b = list[0:]
	for i,v := range b	{
		b[i] = v * a[i]
	}
	fmt.Println(b)
	ch<-n
}

