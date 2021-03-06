package main

import (
	"fmt"
	"strconv"
	"os"
	"io"
	"csv"
	"flag"
)

type Matrix struct {
	Rows,Columns int
	Data [][]int
}

func init() {
        flag.IntVar(&BlockSize,"b",5,"specifies the size of blocks that will be run concurrently")
        flag.IntVar(&SliceSize,"s",5,"specifies the size of slices to be multiplied")
}

func main() {
	flag.Parse()
	fmt.Println(BlockSize)
	/*for i := 0;i < mat.Rows;i++ {
		mat.Data[i] = make([]int,mat.Columns)
	}*/
	mat := OpenCsv("/home/akarsh/Documents/go_code/data.csv")
	for i:=0;i<mat.Rows;i++ {
		col1 := mat.GetCol(i)
		row1 := mat.GetRow(i)
		fmt.Println(RowColMultiplier(row1,col1))
	}
}

func OpenCsv(s string) (mat *Matrix){
	f,err := os.Open(s)
	if err != nil {
		fmt.Println("Could not Open the CSV File")
		return
	}
	read := csv.NewReader(io.Reader(f))
	data,err := read.ReadAll()
	mat = &Matrix{len(data),len(data[0]),make([][]int,len(data))}
	initMatrix(mat)
	if err != nil {
		fmt.Println("Failed to read from the CSV File(Maybe the file does not comply to the CSV standard defined in RFC 4180)")
	}
	for i:=0;i < len(data);i++ {
		for j:=0;j<len(data[i]);j++ {
			ret,_ := strconv.Atoi(data[i][j])
			mat.Data[i][j] = ret
		}
	}
	return
}

func initMatrix(mat *Matrix) {
	for i := 0;i < mat.Rows;i++ {
                mat.Data[i] = make([]int,mat.Columns)
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

