package main

import (
	"runtime"
	"fmt"
)

func main()  {
	runtime.GOMAXPROCS(runtime.NumCPU())
	ch := make(chan int, 1)
	for i := 0; i< 10; i++ {
		go func(i int) {
			fmt.Println(i)
			ch <- 1
		}(i)
		<-ch
	}
}
