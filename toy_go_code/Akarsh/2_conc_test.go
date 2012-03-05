package main

import (
	"fmt"
)

func main() {
	list := []int{1,145,446,34,4543,45,45,45,32,456,56,6,767,76,89,766,45,34,545,53,534,65,645,754,756,645,11,4123,4132,14345,12345,53541,1,134,3634664,3416888,4876568,6846491,11,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1}	
	ch := make(chan bool)
	go sort(list[:20],ch)
	go sort(list[20:],ch)
	<-ch
	<-ch
}

func sort(a []int,ch chan bool)	{
	var temp int
	for i:=len(a)-1 ; i>0 ; i-- {
		for j:=0 ; j<i ; j++ {
			if a[j] > a[j+1] {
				temp = a[j]
				a[j] = a[j+1]
				a[j+1] = temp
			}
		}
	}
	fmt.Println(a)
	ch<-true	
}
