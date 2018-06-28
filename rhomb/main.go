package main

import (
	"os"
	"strconv"
	"fmt"
)

/**
 * @param length 画菱形所占区域的长度
 * @param heigth 菱形高度的一半
 */
func getRhombArr(length int, height int) ([][]string) {
	if length < 2 * height + 1 {
		panic("length need to be longer")
	}

	rhombArr := make([][]string, 0, height)

	for i := 0; i < height; i++ {
		arr := make([]string, 0, length)
		for j := 0; j < length; j++ {
			if (j > length/2 && j-length/2 <= i) || (j <= length/2 && length/2-j <= i) {
				arr = append(arr, "*")
			} else {
				arr = append(arr, " ")
			}
		}
		rhombArr = append(rhombArr, arr)
	}

	for i := height; i >= 0; i-- {
		arr := make([]string, 0, length)
		for j := 0; j < length; j++ {
			if (j > length/2 && j-length/2 <= i) || (j <= length/2 && length/2-j <= i) {
				arr = append(arr, "*")
			} else {
				arr = append(arr, " ")
			}
		}
		rhombArr = append(rhombArr, arr)
	}
	return rhombArr
}

func main()  {
	length, err := strconv.Atoi(os.Args[1])
	if err != nil {
		panic(err)
	}
	height, err := strconv.Atoi(os.Args[2])
	if err != nil {
		panic(err)
	}
	rhombArr := getRhombArr(length, height)
	for _, arr := range rhombArr {
		for _, r := range arr {
			fmt.Print(r)
		}
		fmt.Println()
	}
}
