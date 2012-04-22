//Unit Testing
package ParallelMat //same package name as source file

import (
    "testing" //import go package for testing related functionality
    )

func Test_Matmul1(t *testing.T) { 
    
    mat1="./data1.csv"
    
    NumWorkers=5
    mat1 := OpenCsv(mat1)
	mat2 := mat1
	
	//Identity Matrix
	  for i:=0;i<mat2.Columns;i++ {
                 for j:=0;j<mat2.Rows;j++ {
                        if i==j {
                        	mat2.Data[i][j]=1
                        } else{
                        	mat2.Data[i][j]=0
                        }
                      }  
        }
        
	matres := Matrix{mat1.Rows,mat2.Columns,make([][]int,mat1.Rows)}
	done := make(chan bool)
	initMatrix(&matres)//matres.initMatrix() make it this way
	rowCol := make(chan MatrixRowColPair)
	
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
	
	//To test if the result matches
	pass:=true
	 for i:=0;i<matres.Columns;i++ {
                 for j:=0;j<matres.Rows;j++ {
                        if (matres.Data[i][j]!=mat1.Data[i][j]) {
                        	t.Error("Error in matrix multiplication!")
                        	pass=false
                        	break
                        	
                        } else{
                        	mat2.Data[i][j]=0
                        }
                      }
           if !pass{
	break
	}
             
        }
	if pass{
	t.Log("Multiplication with identity matrix:Test passed.")
	}
	
}


