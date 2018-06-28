package main

import (
	"fmt"
	"os"
)

func statistal(str string) (int, int, int, int) {
	a, b, c, d := 0, 0, 0, 0

	runes := []rune(str)
	for i := 0; i < len(runes); i ++ {
		r := runes[i]
		if (r >= 65 && r <= 90) || (r >= 97 && r <= 122) {
			a++
		} else if r == 32 {
			b++
		} else if r >= 48 && r <= 57 {
			c++
		} else {
			d++
		}
	}
	return a, b, c, d
}


//func GetBorderCharASCII() {
//	str := "azAZ09 你好*"
//	runes := []rune(str)
//	for i := 0; i < len(runes); i++ {
//		fmt.Println(runes[i])
//	}
//}

func main() {
	str := os.Args[1]
	a, b, c, d := statistal(str)
	fmt.Printf("%s:\n字母:%d,空格:%d,数字:%d,其它字符:%d\n\n", str, a, b, c, d)
}