package ParallelMatMul

import (
	"fmt"
	. "../comm"
)

var NumWorkers int = 5

func ParMatMul(mat1, mat2, matres *Matrix) {
	done := make(chan bool)
	rowCol := make(chan MatrixRowColPair)

	go func() {
		for i := 0; i < mat2.Columns; i++ {
			col1 := mat2.GetCol(i) //was row1 := mat1.GetRow(i)
			for j := 0; j < mat1.Rows; j++ {
				row1 := mat1.GetRow(j) //was col1 := mat2.GetCol(j)
				matobj := MatrixRowColPair{j, i, row1, col1}
				rowCol <- matobj
			}
		}
		close(rowCol)
	}()

	for i := 0; i < NumWorkers; i++ {
		go RowColMultiplier(matres, rowCol, done)
	}
	for i := 0; i < NumWorkers; i++ {
		<-done
	}
}


func SeqMatMult(A [][]int, B [][]int, C [][]int) {
	m := len(A) //determine size of the matrix
	n := len(B[0])
	p := len(B)

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			C[i][j] = 0.0
			for k := 0; k < p; k++ {
				C[i][j] += A[i][k] * B[k][j]
			}
		}
	}
}

func CheckResults(C, C1 [][]int) bool {
	m := len(C)  //determines the size of the matrix from the matrix rather than using variables passed 
	n := len(C1) //as arguments
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if C[i][j] != C1[i][j] {
				fmt.Printf("C is\n")
				for i := 0; i < m; i++ {
					for j := 0; j < n; j++ {
						fmt.Printf("%v ", C[i][j])
					}
					fmt.Printf("\n")
				}
				fmt.Printf("C1 is\n")
				for i := 0; i < m; i++ {
					for j := 0; j < n; j++ {
						fmt.Printf("%v ", C1[i][j])
					}
					fmt.Printf("\n")
				}
				return false //return false if the matrix multiplication was not valid
			}
		}
	}
	return true //returning true on successfull validation
}
