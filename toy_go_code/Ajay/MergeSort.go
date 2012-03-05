package main

import (
	"fmt"
	"os"
	"flag"
)

func Mergesort(a []int) {
	if len(a)>1 {
		/*divide and sort*/
		b := make([]int, len(a)/2)
		c := make([]int, len(a)-len(a)/2)
		copy(b, a[:int(len(a)/2)])
		copy(c, a[int(len(a)/2):])
		Mergesort(b)
		Mergesort(c)
		/*merge*/ 		
		i,j,k := 0,0,0
		for i<len(b)&&j<len(c) {
			if b[i]<=c[j] {				
				a[k] = b[i]
				i++
			}else if c[j]<b[i] {
				a[k] = c[j]
				j++
			}
			k++
		}
		if(i==len(b)) {
			copy(a[k:len(a)],c[j:len(c)])
		} else {
			copy(a[k:len(a)],b[i:len(b)])
		}
	}	
}

var 

func main() {
	a := []int{6,5,4,3,2,1}
	fmt.Println("Starting merge sort, the array is ", a)
	Mergesort(a)
	fmt.Println("Finished merge sort, the array is ", a)
}

