package comm

import "fmt"

func InitMatrix(mat *Matrix) {
	for i := 0; i < mat.Rows; i++ {
		mat.Data[i] = make([]int, mat.Columns)
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
				if m < 20 && n<20 {
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
				}
				return false //return false if the matrix multiplication was not valid
			}
		}
	}
	return true //returning true on successfull validation
}
