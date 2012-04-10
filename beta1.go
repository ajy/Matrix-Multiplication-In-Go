package main
import (
	"fmt"
	"flag"
	"runtime"
	"runtime/pprof"
	"os"
	"time"
	"log"
)
var mat1 string
var mat2 string
var NumWorkers int

func init() {
	flag.StringVar(&mat1,"mat1","./data1.csv","Path to the CSV data file.")
	flag.StringVar(&mat2,"mat2","./data2.csv","Path to the CSV data file.")
	flag.IntVar(&NumWorkers,"workers",5,"number of goroutines doing the work")
}


func main() {

runtime.GOMAXPROCS(4)
    flag.Parse()
 //Memory and CPU Profiling.Use gopprof matmul cpuprofile and gopprof matmul memprofile to see profiling information
    var cpuprofile="cpuprofile"
    var memprofile="memprofile"
 f1, err := os.Create(cpuprofile)
 if err != nil {

            log.Fatal(err)
        }
  f2, err := os.Create(memprofile)
 if err != nil {

            log.Fatal(err)
        }

        pprof.StartCPUProfile(f1)
defer  pprof.StopCPUProfile() //Happens when main() returns
pprof.WriteHeapProfile(f2) //Memory profiler

//Reading the matrices from csv files
    start := time.Now() 
    mat1 := OpenCsv(mat1)
	mat2 := OpenCsv(mat2)
    end := time.Now() 
	rtime := end.Sub(start)
	fmt.Printf("===Time Taken to read Matrices %v s\n",rtime.Seconds())
	//fmt.Println()	
	matres := Matrix{mat1.Rows,mat2.Columns,make([][]int,mat1.Rows)}
	done := make(chan bool)
	initMatrix(&matres)//matres.initMatrix() make it this way
	rowCol := make(chan MatrixRowColPair)
	
	
	start = time.Now()
	go func() {
        for i:=0;i<mat2.Columns;i++ {
                col1 := mat2.GetCol(i)//was row1 := mat1.GetRow(i)
                for j:=0;j<mat1.Rows;j++ {
                        row1 := mat1.GetRow(j)//was col1 := mat2.GetCol(j)
                        matobj := MatrixRowColPair{j , i,row1 , col1}
                        rowCol <- matobj
                        }
        }
		close(rowCol)
	}();

	for i := 0;i < NumWorkers;i++ {
                go RowColMultiplier(&matres, rowCol,done)
	}
	for i := 0;i < NumWorkers;i++ {
                <-done
	}
	end = time.Now()	
	
	mtime := end.Sub(start)
	
	fmt.Printf("===Time taken for multiplication %v s ",mtime.Seconds())	
	fmt.Println()
	fmt.Printf("===Total time taken %v s ",(rtime+mtime).Seconds())

}
