package main

import (
	"./ParStrassen"
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
	if(A.Columns!=B.Rows){
		fmt.Println("These matrices cannot be multiplied, %s has %d columns and %s has %d rows",mat1,A.Columns,mat2,B.Rows)
		os.Exit(1)
	}
	C := Matrix{A.Rows, B.Columns, make([][]int, A.Rows)}
	InitMatrix(&C)
	C4 := Matrix{A.Rows, B.Columns, make([][]int, A.Rows)}
	InitMatrix(&C4)

	runtime.GOMAXPROCS(CoreNo)//set number of cores Go can use

	fmt.Printf("Executing Standard matrix multiplication\n")
	before := time.Now()
	SeqMatMult(A.Data, B.Data, C.Data)
	after := time.Now()
	fmt.Printf("Standard matrix multiplication done in %v s\n\n\n", after.Sub(before).Seconds())

	fmt.Print("Executing Strassen matrix multiplication\n")
	before = time.Now()
	ParStrassen.Mul(A, B, &C4)
	after = time.Now()
	fmt.Printf("Strassen matrix multiplication done in %v s\n\n\n", after.Sub(before).Seconds())

	fmt.Println("Checking for errors")
	if CheckResults(C.Data, C4.Data) {
		fmt.Println("\nNo errors occured")
	} else {
		fmt.Println("\nError detected\n")
	}
}
