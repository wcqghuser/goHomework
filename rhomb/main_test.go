package main

import (
	"testing"
	"reflect"
)

func TestGetRhombArr(t *testing.T) {
	length, height := 5, 2
	targetArr := [][]string{
		{" "," ","*"," "," "},
		{" ","*","*","*"," "},
		{"*","*","*","*","*"},
		//{" ","*","*","*","*"},   //wrong cotent
		{" ","*","*","*"," "}, //right content
		{" "," ","*"," "," "},
	}
	resultArr := getRhombArr(length, height)
	if !StringSliceReflectEqual(targetArr, resultArr) {
		t.Errorf("result should be:\n %s\n but is:\n %s", targetArr, resultArr)
	}

	length, height = 7, 3
	targetArr = [][]string{
		{" "," "," ","*"," "," "," "},
		{" "," ","*","*","*"," "," "},
		{" ","*","*","*","*","*"," "},
		{"*","*","*","*","*","*","*"},
		{" ","*","*","*","*","*"," "},
		{" "," ","*","*","*"," "," "},
		{" "," "," ","*"," "," "," "},
	}
	resultArr = getRhombArr(length, height)
	if !StringSliceReflectEqual(targetArr, resultArr) {
		t.Errorf("result should be:\n %s\n but is:\n %s", targetArr, resultArr)
	}
}


func StringSliceReflectEqual(a, b [][]string) bool {
	return reflect.DeepEqual(a, b)
}