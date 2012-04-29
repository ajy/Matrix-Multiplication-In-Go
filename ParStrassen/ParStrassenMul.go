package ParStrassen

import (
	. "../comm"	
)

var GRAIN int /* size of product(of dimensions) below which matmultleaf is used */

func matmultleaf(mf, ml, nf, nl, pf, pl int, A, B, C [][]int) {
	/* 
		subroutine that uses the simple triple loop to multiply 
		a submatrix from A with a submatrix from B and store the 
		result in a submatrix of C. 
	*/
	// mf, ml; /* first and last+1 i index */ 
	// nf, nl; /* first and last+1 j index */ 
	// pf, pl; /* first and last+1 k index */ 
	for i := mf; i < ml; i++ {
		for j := nf; j < nl; j++ {
			for k := pf; k < pl; k++ {
				C[i][j] += A[i][k] * B[k][j]
			}
		}
	}
}

func copyQtrMatrix(X [][]int, m int, Y [][]int, mf, nf int) {
	for i := 0; i < m; i++ {
		X[i] = Y[mf+i][nf:]
	}
}

func AddMats(T [][]int, m, n int, X [][]int, Y [][]int) {
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			T[i][j] = X[i][j] + Y[i][j]
		}
	}
}

func SubMats(T [][]int, m, n int, X [][]int, Y [][]int) {
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			T[i][j] = X[i][j] - Y[i][j]
		}
	}
}

func allocate2DArray(m, n int) [][]int {
	temp := make([][]int, m)
	for i := 0; i < len(temp); i++ {
		temp[i] = make([]int, n)
	}
	return temp
}

func strassenMMult(mf, ml, nf, nl, pf, pl int, A, B, C [][]int) {
	if (ml-mf)*(nl-nf)*(pl-pf) < GRAIN {
		matmultleaf(mf, ml, nf, nl, pf, pl, A, B, C)
	} else {
		m2 := (ml - mf) / 2
		n2 := (nl - nf) / 2
		p2 := (pl - pf) / 2

		A11 := make([][]int, m2)
		A12 := make([][]int, m2)
		A21 := make([][]int, m2)
		A22 := make([][]int, m2)

		B11 := make([][]int, p2)
		B12 := make([][]int, p2)
		B21 := make([][]int, p2)
		B22 := make([][]int, p2)

		C11 := make([][]int, m2)
		C12 := make([][]int, m2)
		C21 := make([][]int, m2)
		C22 := make([][]int, m2)

		copyQtrMatrix(A11, m2, A, mf, pf)
		copyQtrMatrix(A12, m2, A, mf, p2)
		copyQtrMatrix(A21, m2, A, m2, pf)
		copyQtrMatrix(A22, m2, A, m2, p2)

		copyQtrMatrix(B11, p2, B, pf, nf)
		copyQtrMatrix(B12, p2, B, pf, n2)
		copyQtrMatrix(B21, p2, B, p2, nf)
		copyQtrMatrix(B22, p2, B, p2, n2)

		copyQtrMatrix(C11, m2, C, mf, nf)
		copyQtrMatrix(C12, m2, C, mf, n2)
		copyQtrMatrix(C21, m2, C, m2, nf)
		copyQtrMatrix(C22, m2, C, m2, n2)

		done := make(chan int)
		//M1,M2,M3,M4,M5,M6,M7 [][]int
		M1 := allocate2DArray(m2, n2)
		M2 := allocate2DArray(m2, n2)
		M3 := allocate2DArray(m2, n2)
		M4 := allocate2DArray(m2, n2)
		M5 := allocate2DArray(m2, n2)
		M6 := allocate2DArray(m2, n2)
		M7 := allocate2DArray(m2, n2)

		go func(){
			// M1 = (A11 + A22)*(B11 + B22) 
			tAM1 := allocate2DArray(m2, p2)
			AddMats(tAM1, m2, p2, A11, A22)
			tBM1 := allocate2DArray(p2, n2)
			AddMats(tBM1, p2, n2, B11, B22)
			strassenMMult(0, m2, 0, n2, 0, p2, tAM1, tBM1, M1)
			done <- 0
		}()
		
		go func(){
			//M2 = (A21 + A22)*B11 
			tAM2 := allocate2DArray(m2, p2)
			AddMats(tAM2, m2, p2, A21, A22)
			strassenMMult(0, m2, 0, n2, 0, p2, tAM2, B11, M2)
			done <- 0
		}()

		go func(){
			//M3 = A11*(B12 - B22) 
			tBM3 := allocate2DArray(p2, n2)
			SubMats(tBM3, p2, n2, B12, B22)
			strassenMMult(0, m2, 0, n2, 0, p2, A11, tBM3, M3)
			done <- 0
		}()
		go func(){
			//M4 = A22*(B21 - B11) 
			tBM4 := allocate2DArray(p2, n2)
			SubMats(tBM4, p2, n2, B21, B11)
			strassenMMult(0, m2, 0, n2, 0, p2, A22, tBM4, M4)
			done <- 0
		}()

		go func(){
			//M5 = (A11 + A12)*B22 
			tAM5 := allocate2DArray(m2, p2)
			AddMats(tAM5, m2, p2, A11, A12)
			strassenMMult(0, m2, 0, n2, 0, p2, tAM5, B22, M5)
			done <- 0
		}()

		go func(){
			//M6 = (A21 - A11)*(B11 + B12) 
			tAM6 := allocate2DArray(m2, p2)
			SubMats(tAM6, m2, p2, A21, A11)
			tBM6 := allocate2DArray(p2, n2)
			AddMats(tBM6, p2, n2, B11, B12)
			strassenMMult(0, m2, 0, n2, 0, p2, tAM6, tBM6, M6)
			done <- 0
		}()

			//M7 = (A12 - A22)*(B21 + B22) 
			tAM7 := allocate2DArray(m2, p2)
			SubMats(tAM7, m2, p2, A12, A22)
			tBM7 := allocate2DArray(p2, n2)
			AddMats(tBM7, p2, n2, B21, B22)
			strassenMMult(0, m2, 0, n2, 0, p2, tAM7, tBM7, M7)
				
		for cnt := 6;cnt>0;cnt--{//synchronise all running goroutines
			<-done
		}

		for i := 0; i < m2; i++ {
			for j := 0; j < n2; j++ {
				C11[i][j] = M1[i][j] + M4[i][j] - M5[i][j] + M7[i][j]
				C12[i][j] = M3[i][j] + M5[i][j]
				C21[i][j] = M2[i][j] + M4[i][j]
				C22[i][j] = M1[i][j] - M2[i][j] + M3[i][j] + M6[i][j]
			}
		}
	}
}

func Mul(A, B, C *Matrix) {
	GRAIN = A.Rows*B.Columns*2
	strassenMMult(0, A.Rows, 0, A.Columns, 0, B.Columns, A.Data, B.Data, C.Data)
}

