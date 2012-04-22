package main
import (
	"./StrassenMul"
	"flag"
	"fmt"
	"time"
	"os"
	. "./comm"
)

var mat1 string
var mat2 string

func init() {
	flag.StringVar(&mat1,"mat1","./data1.csv","Path to the CSV data file.")
	flag.StringVar(&mat2,"mat2","./data2.csv","Path to the CSV data file.")
}


func main() {
	N := 0
	flag.Parse()
	A := OpenCsv(mat1)
	B := OpenCsv(mat2)
	M := A.Rows
	P := B.Columns
	if(A.Columns==B.Rows){
		N = A.Columns
	} else {
		fmt.Println("These matrices cannot be multiplied, %s has %d columns and %s has %d rows",mat1,A.Columns,mat2,B.Rows)
		os.Exit(1)
	}
	C := StrassenMul.Allocate2DArray(M, N)
	C4 := StrassenMul.Allocate2DArray(M, N)

	fmt.Print("Executing Standard matrix multiplication\n")
	before := time.Now()
	SeqMatMult(A.Data, B.Data, C)
	after := time.Now()
	fmt.Printf("Standard matrix multiplication done in %v s\n\n", after.Sub(before).Seconds())

	fmt.Print("Executing Strassen matrix multiplication\n")
	before = time.Now()
	StrassenMul.MatmultS(M, N, P, A.Data, B.Data, C4)
	after = time.Now()
	fmt.Printf("Strassen matrix multiplication done in %v s\n\n", after.Sub(before).Seconds())

	fmt.Println("Checking for errors")
	if CheckResults(C, C4) {
		fmt.Println("\nNo errors occured")
	} else {
		fmt.Println("\nError detected\n")
	}
}
