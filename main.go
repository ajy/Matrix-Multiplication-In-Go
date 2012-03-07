package main
import (
	"fmt"
	"flag"
)
var mat1 string
var mat2 string

func init() {
	flag.StringVar(&mat1,"mat1","/home/akarsh/Documents/go_code/data.csv","Path to the CSV data file.")
	flag.StringVar(&mat2,"mat2","/home/akarsh/Documents/go_code/data.csv","Path to the CSV data file.")
}



func main() {
        flag.Parse()
        mat1 := OpenCsv(mat1)
	mat2 := OpenCsv(mat2)
	matres := Matrix{mat1.Rows,mat2.Columns,make([][]int,mat1.Rows)}
	done := make(chan bool)
	initMatrix(&matres)//matres.initMatrix() make it this way
	rowCol := make(chan MatrixRowColPair)
	val := make(chan MatEl)

	for i:=0;i<5;i++ {
		go ReadResult(&matres,val,done)
	}


go func() {
        for i:=0;i<mat1.Rows;i++ {
                row1 := mat1.GetRow(i)
                for j:=0;j<mat2.Columns;j++ {
                        col1 := mat2.GetCol(j)
                        matobj := MatrixRowColPair{i , j,row1 , col1}
                        rowCol <- matobj
                        }
        }
	close(rowCol)
}();

	for i := 0;i < 5;i++ {
                go RowColMultiplier(rowCol, val,done)
//                go ReadResult(&matres,val,done)
	}
	<-done
	
	fmt.Println(matres.Data)	
}