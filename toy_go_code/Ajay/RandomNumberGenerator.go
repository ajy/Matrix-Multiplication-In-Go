package main

import (
	"fmt"
	"rand"
	"flag"
	"os"
)

var fileName = flag.String("o", "RandNumFile.txt", "Specifies output file")
var num = flag.Int("n", 0, "specifies number of random numbers to be generated")
var largest = flag.Int("max", 10000, "specifies largest random number generated")

func randGenerator(ch, done chan int){
	for {
		select {
			case ch <- rand.Intn(*largest):
			case <-done:
				return
		}
	}
}

func main(){
	flag.Parse()
	if *num <= 0 {
		fmt.Println("number of random numbers ", num,"to be created < 1")
		os.Exit(1)
	}
	fmt.Println("Starting write")
	fd, err := os.Create(*fileName)
	defer fd.Close()
	if err != nil {
		fmt.Println("could not create file ",*fileName)
		os.Exit(1)
	}
	ch := make(chan int)
	DoneCh := make(chan int)
	go randGenerator(ch, DoneCh)
	for i := *num; i>0; i-- {
		_, err:=fmt.Fprintln(fd, <-ch)
		if err != nil {
			fmt.Println("error while writing to file")
			os.Exit(1)
		}
	}
	DoneCh <- 0//asking randGenerator to stop
	fmt.Println("Done writing")
}



