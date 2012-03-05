/*
Implement Pic. It should return a slice of length dy, each element of which is a slice of dx 8-bit unsigned integers. When you run the program, it will display your picture, interpreting the integers as grayscale (well, bluescale) values.

The choice of image is up to you. Interesting functions include x^y, (x+y)/2, and x*y.

(You need to use a loop to allocate each []uint8 inside the [][]uint8.)
*/
package main

import "go-tour.googlecode.com/hg/pic"

func Pic(dx, dy int) [][]uint8 {
	results := make([][]uint8,dy)
	for i:=0;i<dy;i++ {
		result := make([]uint8,dx)
		for j:=0;j<dx;j++ {
			result[j]=uint8(i+j)/2
		}
		results[i]=result
	}
	return results
}

func main() {
	pic.Show(Pic)
}
