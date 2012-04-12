//Benchmarking
package main //same package name as source file

import (
    "testing" //import go package for testing related functionality
    )

func Benchmark_Matmul1(b *testing.B) { 
    b.StopTimer() //Stopping the performance timer temporarily for reading data
    mat1="./data1.csv"
    mat2="./data2.csv"
    NumWorkers=5
    	mat1 := OpenCsv(mat1)
	mat2 := OpenCsv(mat2)
	
    
	matres := Matrix{mat1.Rows,mat2.Columns,make([][]int,mat1.Rows)}
	done := make(chan bool)
	initMatrix(&matres)//matres.initMatrix() make it this way
	rowCol := make(chan MatrixRowColPair)
	b.StartTimer() //Restart the timer
	
	
	go func() {
        for i:=0;i<mat2.Columns;i++ {
                col1 := mat2.GetCol(i)//was row1 := mat1.GetRow(i)
                for j:=0;j<mat1.Rows;j++ {
                        row1 := mat1.GetRow(j)//was col1 := mat2.GetCol(j)
                        matobj := MatrixRowColPair{j, i,row1 , col1}
                        rowCol <- matobj
                        }
        }
		close(rowCol)
	}();

	for i := 0;i < NumWorkers;i++ {
                go RowColMultiplier(&matres, rowCol,done)
	}
	<-done
	
	
	
	
}


