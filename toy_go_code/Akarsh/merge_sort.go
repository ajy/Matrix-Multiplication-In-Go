package main

import (
//	"fmt"
	"rand"
)

func main() {
	a  := make([]int32,30000)
	for i:= 0;i<30000; i++ {
		a[i] = rand.Int31()
	}
	done := make(chan int)
        go merge_sort(a,done)
	<-done
}

func merge_sort(list []int32,done chan int) {
	if len(list) == 1 {
		done <- 1
		return
	}
	d := make(chan int,2)
	m := int(len(list)/2)
	go merge_sort(list[:m],d)
	go merge_sort(list[m:],d)
	<-d
	<-d
	a := make(chan []int32,1)
	go merge(list[:m],list[m:],a)
	cpy := <-a
	for i := 0; i < len(cpy);i++ {
		list[i] = cpy[i]
	}
	done <- 1
}

func merge(fh []int32,sh []int32,ch chan []int32) {
	ret := make([]int32,len(fh)+len(sh),len(fh)+len(sh))
	var fp,sp  int = 0,0
	var fs, ss int = len(fh),len(sh)
	var list_ptr int32 = 0
	for fp < fs || sp < ss {
		if fp < fs && sp < ss {
			if fh[fp] <= sh[sp] {
				ret[list_ptr] = fh[fp]
				fp += 1
				list_ptr += 1
			} else {
				ret[list_ptr] = sh[sp]
				sp += 1
				list_ptr += 1
			}
		} else if fp == (fs){
			for sp < ss {
				ret[list_ptr] = sh[sp]
				sp += 1
				list_ptr += 1
			}
			break
		} else if sp == (ss){
                        for fp < fs {
                                ret[list_ptr] = fh[fp]
                                fp += 1
                                list_ptr += 1
                        }
			break
                }

	}
	ch <- ret
}
