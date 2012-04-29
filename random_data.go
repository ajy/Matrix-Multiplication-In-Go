package main

import (
    "fmt"
    "math/rand"
    "flag"
)

var s int
func init() {
    flag.IntVar(&s,"size",5,"Size of the random matrix to generate.")
}

func oneRow(s int) {
    data := make([]int,s,s)
    for i := range(data) {
        data[i] = rand.Int()
    }
    for i := range(data[:len(data)-1]) {
        fmt.Print(data[i],",")
    }
    fmt.Print(rand.Int(),"\n")
}

func main() {
    flag.Parse()
    for i:=0;i < s;i++ {
        oneRow(s)
    }
}
