package main

import "testing"

func TestMyInt_Increase(t *testing.T) {
	var result, target myInt

	for i := 0; i<10000; i++ {
		result = myInt(i)
		target = myInt(200 + i)
		result.Increase()
		if target != result {
			t.Errorf("result should be %d, but is %d", target, result)
		}
	}
}
