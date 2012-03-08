package main

import (
	"fmt"
	"strconv"
	"os"
	"io"
	"csv"
)

func OpenCsv(s string) (mat *Matrix){
	f,err := os.Open(s)
	if err != nil {
		fmt.Println("Could not Open the CSV File")
		return
	}
	read := csv.NewReader(io.Reader(f))
	data,err := read.ReadAll()
	mat = &Matrix{len(data),len(data[0]),make([][]int,len(data))}
	initMatrix(mat)
	if err != nil {
		fmt.Println("Failed to read from the CSV File(Maybe the file does not comply to the CSV standard defined in RFC 4180)")
	}
	for i:=0;i < len(data);i++ {
		for j:=0;j<len(data[i]);j++ {
			ret,_ := strconv.Atoi(data[i][j])
			mat.Data[i][j] = ret
		}
	}
	return
}

func initMatrix(mat *Matrix) {
	for i := 0;i < mat.Rows;i++ {
                mat.Data[i] = make([]int,mat.Columns)
        }

}

func (mat * Matrix) GetCol(colNum int)  []int {
	col := make([]int,mat.Rows)
	for i := 0; i < mat.Rows ; i++ {
			col[i] = mat.Data[i][colNum]
	}
	return col
}

func (mat * Matrix) GetRow(rowNum int)  []int {
        return mat.Data[rowNum]
}

func ReadResult(mat *Matrix,val <-chan MatEl,done chan  bool) {
	for ;; {
		a,ok:=<-val
	        if(!ok){
				done <- true
	        	return
	        } else {
	        	mat.Data[a.Row][a.Column] = a.Element
		}
	}
}
