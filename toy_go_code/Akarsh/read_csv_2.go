package main

import (
	"fmt"
	"strconv"
	"os"
	"io"
	"csv"
)

type Matrix struct {
	Rows,Columns int
	Data [][]int
}
func main() {
	mat := Matrix{100,100,make([][]int,100)}
	for i := 0;i < mat.Rows;i++ {
		mat.Data[i] = make([]int,mat.Columns)
	}
	OpenCsv("/home/akarsh/Documents/go_code/data1.csv",&mat)
	for i:=0;i<mat.Rows;i++ {
		col1 := mat.GetCol(i)
		row1 := mat.GetRow(i)
		fmt.Println(RowColMultiplier(row1,col1))
	}
}

func OpenCsv(s string,mat *Matrix) {
	f,err := os.Open(s)
	if err != nil {
		fmt.Println("Could not Open the CSV File")
		return
	}
	read := csv.NewReader(io.Reader(f))
	data,err := read.ReadAll()
	if err != nil {
		fmt.Println("Failed to read from the CSV File(Maybe the file does not comply to the CSV standard defined in RFC 4180)")
	}
	for i:=0;i < len(data);i++ {
		for j:=0;j<len(data[i]);j++ {
			ret,_ := strconv.Atoi(data[i][j])
			mat.Data[i][j] = int(ret / 100)
		}
	}
}

func (mat * Matrix) GetCol(colNum int)  []int {
	col := make([]int,mat.Rows)
	for i := 0; i < mat.Rows ; i++ {
			col[i] = mat.Data[i][colNum]
	}
	return col
}

func (mat * Matrix) GetRow(rowNum int)  []int {
        //row := make([]int,mat.Columns)
	//fmt.Println(rowNum)
        /*for i := 0; i < mat.Columns ; i++ {
                        row = mat.Data[i]
        }*/
        return mat.Data[rowNum]
}

