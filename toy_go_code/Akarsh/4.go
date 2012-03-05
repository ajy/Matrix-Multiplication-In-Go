package main

import (
	"fmt"
)

func send(ch chan int) {
	for i:=0;;i++	{ch<-i}
}

func receive(ch chan int){
	for {fmt.Println(<-ch)}
}

func main(){
	ch1:=make(chan int)
	go send(ch1)
	fmt.Println(<-ch1)
	go receive(ch1)
}

