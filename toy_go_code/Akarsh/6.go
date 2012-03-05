package main

import	(
	"fmt"
	"os"
)

func main()	{
	var a  = make([]byte,100)
	f,err := os.Open("/home/akarsh/remope.txt")
	if err != nil	{
		fmt.Println("Error in opening file")
	}	else	{
		contents,_ := f.Read(a) 
		fmt.Printf("%d,%s",contents,a)
	}
}
