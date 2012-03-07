// this row col multiplier is completely serial and performs i/o through channels
package main

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

	}
}

