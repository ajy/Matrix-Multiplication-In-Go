package main

type Matrix struct {
	Rows,Columns int
	Data [][]int8
}

type MatrixRowColPair struct {
    Row,Col int
    RowData []int8
    ColData []int8
}

type MatEl struct {
	Row,Column int
	Element int
}
