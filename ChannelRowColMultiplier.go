// this row col multiplier is completely serial and performs i/o through channels
package main

/*import (
	"fmt"
	"flag"
)*/

//import "os"

/*var sliceSize int

func init() {
	flag.IntVar(&sliceSize,"s",5,"specifies the size of slices to be multiplied")
}*/

func RowColMultiplier(rowCol <-chan MatrixRowColPair, val chan<- MatEl,done chan bool) {
for ;; {
	pair,ok := <- rowCol 
	if(!ok) {
        	done<- true
                return
        } else  {
		sum:=0
                for i:=0;i<len(pair.RowData);i++ {
                	sum += pair.RowData[i]*pair.ColData[i]
                }
                	val <- MatEl{pair.Row,pair.Col,sum}
        }

/*	select	{
		case pair,ok := <- rowCol : {
			if(!ok) {
				done1<- true
				os.Exit(1) 	
			} else	{
				sum:=0
				for i:=0;i<len(pair.RowData);i++ {
					sum += pair.RowData[i]*pair.ColData[i]
				}
				val <- MatEl{pair.Row,pair.Col,sum}
			}
		}
	}*/
}//end for
}

/*func main() {//run this to check
	flag.Parse()// must be called before flags are used
	a,b := make([]int, sliceSize),make([]int, sliceSize)
	fmt.Println("Creating slices of len ", sliceSize)
	for i:=0;i<sliceSize;i++ {//creating 2 slices containing only 1
		a[i],b[i]=1,1
	}
	rowCol := make(chan MatrixRowColPair, 1)//channels must be buffered coz the routine that reads them aren't active yet
	res := make(chan MatEl, 1)
	rowCol <- MatrixRowColPair{0,0,a,b}
	RowColMultiplier(rowCol, res)
	temp := <-res
	fmt.Println("The result is ", temp)//should be same as length if the slices are multiplied correctly
}*/
