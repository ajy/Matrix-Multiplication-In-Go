package ParallelMat

import (
	. "../comm"
)

var NumWorkers int = 5

func Mul(mat1, mat2, matres *Matrix) {
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
