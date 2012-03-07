// this row col multiplier is completely serial and performs i/o through channels
package main

func RowColMultiplier(mat *Matrix, rowCol <-chan MatrixRowColPair,done chan bool) {
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
           	mat.Data[pair.Row][pair.Col] = sum
	    }
	}
}

