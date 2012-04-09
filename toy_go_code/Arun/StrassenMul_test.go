package main

import (
    "testing" //import go package for testing related functionality
	"flag"
	"fmt"
	"time"

    )
    

func Test_Multiply1024(t *testing.T) { //test function starts with "Test" and takes a pointer to type testing.T
    flag.Parse()
	M := MatSize
	N := MatSize
	P := MatSize
	GRAIN=1024*1024
	A := Allocate2DArray(M, P)
	B := Allocate2DArray(P, N)
	C := Allocate2DArray(M, N)
	C4 := Allocate2DArray(M, N)

	randomize(M,P,A)
	identity(P,N,B)
	
	seqMatMult(M, N, P, A, B, C)
	/*fmt.Printf("Execute Standard matmult\n\n")
	before := time.Now()
	seqMatMult(M, N, P, A, B, C)
	after := time.Now()
	fmt.Printf("Standard matrix function done in %v ns\n\n\n", (after.Sub(before)))*/
	fmt.Printf("GRAIN= %d",GRAIN)
	before := time.Now()
	matmultS(M, N, P, A, B, C4)
	after := time.Now()
	fmt.Printf("Strassen matrix function done in %v \n\n\n", (after.Sub(before)))

	if CheckResults(M, N, C, C4) {
		fmt.Printf("Error in matmultS\n\n")
	} else {
		fmt.Printf("OKAY\n\n")
	}    
    }

func Test_Multiply512(t *testing.T) { //test function starts with "Test" and takes a pointer to type testing.T
    flag.Parse()
	M := MatSize
	N := MatSize
	P := MatSize
	GRAIN=512*512
	A := Allocate2DArray(M, P)
	B := Allocate2DArray(P, N)
	C := Allocate2DArray(M, N)
	C4 := Allocate2DArray(M, N)

	randomize(M,P,A)
	identity(P,N,B)
	
/*	
	fmt.Printf("Execute Standard matmult\n\n")
	before := time.Now()
	seqMatMult(M, N, P, A, B, C)
	after := time.Now()
	fmt.Printf("Standard matrix function done in %v ns\n\n\n", (after.Sub(before)))
*/
	seqMatMult(M, N, P, A, B, C)
	fmt.Printf("GRAIN= %d",GRAIN)
	before := time.Now()
	matmultS(M, N, P, A, B, C4)
	after := time.Now()
	fmt.Printf("Strassen matrix function done in %v ns\n\n\n", (after.Sub(before)))

	if CheckResults(M, N, C, C4) {
		fmt.Printf("Error in matmultS\n\n")
	} else {
		fmt.Printf("OKAY\n\n")
	}   
    }

