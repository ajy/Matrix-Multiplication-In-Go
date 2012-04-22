package main

import (
	"./ParStrassenMul"
	"flag"
	"fmt"
	"time"
	"os"
	"runtime"
	. "./comm"
)

var mat1 string
var mat2 string
var CoreNo int

func init() {
	flag.StringVar(&mat1,"mat1","./data1.csv","Path to the CSV data file.")
	flag.StringVar(&mat2,"mat2","./data2.csv","Path to the CSV data file.")
	flag.IntVar(&CoreNo, "cores", 4, "specifies the number of cores Go can use to execute this code")
}

func main() {
	flag.Parse()
	A := OpenCsv(mat1)
	B := OpenCsv(mat2)
	M := A.Rows
	P := B.Columns
	if(A.Columns!=B.Rows){
		fmt.Println("These matrices cannot be multiplied, %s has %d columns and %s has %d rows",mat1,A.Columns,mat2,B.Rows)
		os.Exit(1)
	}
	N := A.Columns
	C := ParStrassenMul.Allocate2DArray(M, N)
	C4 := ParStrassenMul.Allocate2DArray(M, N)

	runtime.GOMAXPROCS(CoreNo)//set number of cores Go can use

	fmt.Printf("Executing Standard matrix multiplication\n")
	before := time.Now()
	ParStrassenMul.SeqMatMult(M, N, P, A.Data, B.Data, C)
	after := time.Now()
	fmt.Printf("Standard matrix multiplication done in %v s\n\n\n", after.Sub(before).Seconds())

	fmt.Print("Executing Strassen matrix multiplication\n")
	before = time.Now()
	ParStrassenMul.ParMatmultS(M, N, P, A.Data, B.Data, C4)
	after = time.Now()
	fmt.Printf("Strassen matrix multiplication done in %v s\n\n\n", after.Sub(before).Seconds())

	fmt.Println("Checking for errors")
	if ParStrassenMul.CheckResults(C, C4) {
		fmt.Println("\nNo errors occured")
	} else {
		fmt.Println("\nError detected\n")
	}
}
