package main
import (
//	"fmt"
	"flag"
)
var mat1 string
var mat2 string
var NumWorkers int

func init() {
	flag.StringVar(&mat1,"mat1","./data1.csv","Path to the CSV data file.")
	flag.StringVar(&mat2,"mat2","./data2.csv","Path to the CSV data file.")
	flag.IntVar(&NumWorkers,"workers",5,"number of goroutines doing the work")
}


func main() {
    flag.Parse()
    mat1 := OpenCsv(mat1)
	mat2 := OpenCsv(mat2)
	matres := Matrix{mat1.Rows,mat2.Columns,make([][]int,mat1.Rows)}
	done := make(chan bool)
	initMatrix(&matres)//matres.initMatrix() make it this way
	rowCol := make(chan MatrixRowColPair)
	
	go func() {
        for i:=0;i<mat2.Columns;i++ {
                col1 := mat2.GetCol(i)//was row1 := mat1.GetRow(i)
                for j:=0;j<mat1.Rows;j++ {
                        row1 := mat1.GetRow(j)//was col1 := mat2.GetCol(j)
                        matobj := MatrixRowColPair{i , j,row1 , col1}
                        rowCol <- matobj
                        }
        }
		close(rowCol)
	}();

	for i := 0;i < NumWorkers;i++ {
                go RowColMultiplier(&matres, rowCol,done)
	}
	<-done
	
//	fmt.Println(matres.Data)	
}
