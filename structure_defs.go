package main

type Matrix struct {
	Rows,Columns int
	Data [][]int
}

type MatrixRowColPair struct {
    Row,Col int
    RowData []int
    ColData []int
}

type MatEl struct {
	Row,Column int
	Element int
}
