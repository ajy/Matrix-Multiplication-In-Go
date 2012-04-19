package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"runtime"
	"runtime/pprof"
	"time"
)

var mat1 string
var mat2 string
var NumWorkers int

func init() {
	flag.StringVar(&mat1, "mat1", "../data1.csv", "Path to the CSV data file.")
	flag.StringVar(&mat2, "mat2", "../data2.csv", "Path to the CSV data file.")
	flag.IntVar(&NumWorkers, "workers", 5, "number of goroutines doing the work")
}

func main() {

	runtime.GOMAXPROCS(4)
	flag.Parse()
	//Memory and CPU Profiling.Use gopprof matmul cpuprofile and gopprof matmul memprofile to see profiling information
	var cpuprofile = "cpuprofile"
	var memprofile = "memprofile"
	f1, err := os.Create(cpuprofile)
	if err != nil {

		log.Fatal(err)
	}
	f2, err := os.Create(memprofile)
	if err != nil {

		log.Fatal(err)
	}

	pprof.StartCPUProfile(f1)
	defer pprof.StopCPUProfile() //Happens when main() returns
	pprof.WriteHeapProfile(f2)   //Memory profiler

	//Reading the matrices from csv files
	start := time.Now()
	mat1 := OpenCsv(mat1)
	mat2 := OpenCsv(mat2)
	end := time.Now()
	rtime := end.Sub(start)
	fmt.Printf("\nTime Taken to read Matrices %v s\n", rtime.Seconds())
	matres := Matrix{mat1.Rows, mat2.Columns, make([][]int, mat1.Rows)}
	done := make(chan bool)
	initMatrix(&matres) //matres.initMatrix() make it this way
	rowCol := make(chan MatrixRowColPair)

	matValidate := Matrix{mat1.Rows, mat2.Columns, make([][]int, mat1.Rows)} //Matrix for validating the results
	initMatrix(&matValidate)
	fmt.Println("\nExecuting Parallel Matrix Multiplication")
	start = time.Now()
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
		go RowColMultiplier(&matres, rowCol, done)
	}
	for i := 0; i < NumWorkers; i++ {
		<-done
	}
	end = time.Now()

	mtime := end.Sub(start)

	fmt.Printf("\nParallel matrix multiplication done in %v s ", mtime.Seconds())
	fmt.Println()
	fmt.Printf("\nTotal time taken %v s ", (rtime + mtime).Seconds())

	//Validation
	fmt.Println("\nChecking for errors using standard matrix multiplication")
	seqMatMult(mat1.Data, mat2.Data, matValidate.Data)
	if CheckResults(matres.Data,matValidate.Data) {
		fmt.Println("\nNo errors occured")
	} else {
		fmt.Println("\nError detected\n")
	}
}

func seqMatMult(A [][]int, B [][]int, C [][]int) {
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
