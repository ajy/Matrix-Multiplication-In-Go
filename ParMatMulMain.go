package main

import (
	"./ParallelMatMul"
	"fmt"
	"flag"
	"log"
	"os"
	"runtime"
	"runtime/pprof"
	"time"
	. "./comm"
)
var mat1loc string
var mat2loc string
var NumWorkers int

func init() {
	flag.StringVar(&mat1loc, "mat1", "./data1.csv", "Path to the CSV data file.")
	flag.StringVar(&mat2loc, "mat2", "./data2.csv", "Path to the CSV data file.")
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
	mat1 := OpenCsv(mat1loc)
	mat2 := OpenCsv(mat2loc)
	end := time.Now()
	rtime := end.Sub(start)
	fmt.Printf("\nTime Taken to read Matrices %v s\n", rtime.Seconds())
	matres := Matrix{mat1.Rows, mat2.Columns, make([][]int, mat1.Rows)}
	InitMatrix(&matres) //matres.initMatrix() make it this way
	
	//ParallelMatMul.CheckResults(mat1.Data,mat2.Data)see what's in the matrices
	
	matValidate := Matrix{mat1.Rows, mat2.Columns, make([][]int, mat1.Rows)} //Matrix for validating the results
	InitMatrix(&matValidate)
	ParallelMatMul.NumWorkers = NumWorkers// need not be set, has default
	fmt.Println("\nExecuting Parallel Matrix Multiplication")
	start = time.Now()
	ParallelMatMul.ParMatMul(mat1, mat2, &matres)
	end = time.Now()

	mtime := end.Sub(start)

	fmt.Printf("\nParallel matrix multiplication done in %v s ", mtime.Seconds())
	fmt.Println()
	fmt.Printf("\nTotal time taken %v s ", (rtime + mtime).Seconds())

	//Validation
	fmt.Println("\nChecking for errors using standard matrix multiplication")
	ParallelMatMul.SeqMatMult(mat1.Data, mat2.Data, matValidate.Data)
	if ParallelMatMul.CheckResults(matres.Data,matValidate.Data) {
		fmt.Println("\nNo errors occured")
	} else {
		fmt.Println("\nError detected\n")
	}
}
