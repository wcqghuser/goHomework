package main

import (
	"fmt"
)

type Student struct {
	name string
	sex int
	no int
	age int
	score int
}

func main()  {
	maps := map[int]Student {
		1: {"wcq", 1, 1, 27, 79},
		2: {"xiaoming", 1, 2, 28, 89},
		3: {"lucy", 0, 3, 26, 70},
		4: {"hanmeimei", 0, 4, 37, 80},
		5: {"xxx", 1, 5, 16, 75},
		6: {"yyy", 1, 6, 18, 73},
		7: {"zzz", 1, 7, 22, 60},
		8: {"qqq", 1, 8, 19, 66},
		9: {"eee", 1, 9, 20, 62},
		10: {"aaa", 1, 10, 44, 100},
	}

	studentSlice := make([]Student, 0, 10)
	i := 0
	for _, stu := range maps {
		studentSlice = append(studentSlice, stu)
		for j, s:= range studentSlice {
			if stu.score > s.score  {
				studentSlice[j], studentSlice[i] = studentSlice[i], studentSlice[j]
			}
		}
		i++
	}

	for _, s := range studentSlice {
		fmt.Println(s)
	}
}

