package main
import (
"fmt"
"rand"
)
func main(){
	var row,col int
	row=20
	col=20
	/*.Print("Enter rows & cols: ")
	fmt.Scan(&row,&col)*/
	a:=make([][]int,row)
	b:=make([][]int,row)
	for i := range a {
        a[i] = make([]int, col)
    }	
    for i := range b {
        b[i] = make([]int, col)
    }	
	randomize(a,row,col)
	randomize(b,row,col)
	c:=make([][]int,row)
	rows,cols,extra:=len(a),len(b[0]),len(b)
	for i:=0;i<rows;i++{
c[i]=make([]int,col)
		for j:=0;j<cols;j++{
				for k:=0;k<extra;k++{
			c[i][j]+=a[i][k]*b[k][j]
}
}
}
print(a)
print(b)
fmt.Print("Result:")
print(c)

}
func print(x[][]int){
	fmt.Println()
	for _,r := range x{
fmt.Println(r)}
}
func randomize(x [][]int,row,col int){
	for i:=0;i<row;i++{
		//x[i]=make([]int,col)
		for j:=0;j<col;j++{
			x[i][j]=rand.Int()%10
}
}
}
