package main

import (
	"fmt"
	"strconv"
	"os"
)

type myInt int

func (i *myInt) Increase() {
	*i += 200
}

func main()  {
	arg, err := strconv.Atoi(os.Args[1])
	if err != nil {
		panic(err)
	}

	a := myInt(arg)
	a.Increase()
	fmt.Println(a)
}
