package main

import (
	"fmt"
	"strconv"
	"os"
	"io"
	"csv"
)

func main() {
	OpenCsv("/home/akarsh/Documents/go_code/data.csv")
}

func OpenCsv(s string) {
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
	data0 := make([]int,len(data[0])*len(data))
	for i:=0;i < len(data);i++ {
		for j:=0;j<len(data[i]);j++ {
			data0[(i*len(data[0]))+j],_ = strconv.Atoi(data[i][j])
		}
	}
	fmt.Println(data0)
}
