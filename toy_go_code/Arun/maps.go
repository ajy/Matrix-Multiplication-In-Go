package main

import(
//"fmt"
	"strings"
	"go-tour.googlecode.com/hg/wc"
)

func WordCount(s string) map[string]int {
	//s :="I ate a donut. Then I ate another donut."
	w :=strings.Fields(s)
	//fmt.Println(w)
	m:=make(map[string]int)
	for i:=0 ;i<len(w) ;i++ {
		m[w[i]]=strings.Count(s,w[i])
	}
	return m
}

func main() {
	/*s :=strings.TrimSpace("I ate a donut. Then I ate another donut.")
	w :=strings.Fields(s)
	
	fmt.Println(w)
	m:=make(map[string]int)
	for i:=0 ;i<len(w) ;i++ {
		m[w[i]]=strings.Count(s,w[i])
	}
	fmt.Println(m)	*/
wc.Test(WordCount)
}
