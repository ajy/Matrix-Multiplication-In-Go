package main

import (
	"./ParallelMat"
	"./ParStrassen"
	"./Strassen"
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
	if(mat1.Columns!=mat2.Rows){
		fmt.Println("These matrices cannot be multiplied, %s has %d columns and %s has %d rows",mat1loc,mat1.Columns,mat2loc,mat2.Rows)
		os.Exit(1)
	}
	rtime := end.Sub(start)
	fmt.Printf("\nTime Taken to read Matrices %v s\n", rtime.Seconds())
	
	matValidate := Matrix{mat1.Rows, mat2.Columns, make([][]int, mat1.Rows)} //Matrix for validating the results
	InitMatrix(&matValidate)
	fmt.Println("\nExecuting Serial Matrix Multiplication")
	start = time.Now()
	SeqMatMult(mat1.Data, mat2.Data, matValidate.Data)
	end = time.Now()
	fmt.Printf("\nSerial matrix multiplication done in %v s\n", end.Sub(start).Seconds())
	
	
	matres := Matrix{mat1.Rows, mat2.Columns, make([][]int, mat1.Rows)}
	InitMatrix(&matres) //matres.initMatrix() make it this way
	
	
	ParallelMat.NumWorkers = NumWorkers// need not be set, has default
	fmt.Println("\nExecuting Parallel Matrix Multiplication")
	start = time.Now()
	ParallelMat.Mul(mat1, mat2, &matres)
	end = time.Now()
	fmt.Printf("Parallel matrix multiplication done in %v s\n", end.Sub(start).Seconds())
	
	//Validation
	fmt.Println("\nChecking for errors in Parallel Matrix Multiplication using standard matrix multiplication")
	if CheckResults(matres.Data,matValidate.Data) {
		fmt.Println("\nNo errors occured\n")
	} else {
		fmt.Println("\nError detected\n")
	}
	
	fmt.Print("\nExecuting Strassen matrix multiplication\n")
	start = time.Now()
	Strassen.Mul(mat1, mat2, &matres)
	end = time.Now()
	fmt.Printf("Strassen matrix multiplication done in %v s\n", end.Sub(start).Seconds())
	
	//Validation
	fmt.Println("\nChecking for errors in Strassen Matrix Multiplication using standard matrix multiplication")
	if CheckResults(matres.Data,matValidate.Data) {
		fmt.Println("\nNo errors occured\n")
	} else {
		fmt.Println("\nError detected\n")
	}
	
	fmt.Println("\nExecuting Parallel Strassen matrix multiplication")
	start = time.Now()
	ParStrassen.Mul(mat1, mat2, &matres)
	end = time.Now()
	fmt.Printf("Parallel Strassen matrix multiplication done in %v s\n", end.Sub(start).Seconds())
	
	//Validation
	fmt.Println("\nChecking for errors in Parallel Strassen Matrix Multiplication using standard matrix multiplication\n")
	if CheckResults(matres.Data,matValidate.Data) {
		fmt.Println("\nNo errors occured\n")
	} else {
		fmt.Println("\nError detected\n")
	}
	

}
