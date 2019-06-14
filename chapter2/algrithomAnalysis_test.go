package main

import (
	"testing"
)


var randNumbersTests = []int{
	-1,
	0,
	10,
	100,
	9999,
}

func checkRandNumbers(arr []int, n int) bool  {
	//if n <= 0 && len(arr) == 0 {
	//	return true
	//}
	if  len(arr) != n{
		return false
	}
	temp := make([]int,n)
	for _, v := range arr {
		if temp[v - 1] != 0 {
			return false
		}else {
			temp[v - 1] = 1
		}
	}
	return true
}

func TestRandNumber(t *testing.T)  {
	testRandNumber(t,randNumber1)
	testRandNumber(t,randNumber2)
	testRandNumber(t,randNumber3)
}

func testRandNumber(t *testing.T, f func(int)[]int)  {
	for _, tt := range randNumbersTests{
		arr := f(tt)
		if !checkRandNumbers(arr,tt) {
			t.Errorf("test %v: input %d, output: %v not expected\n",f,tt,arr)
		}
	}

}


func BenchmarkRandNumber(b *testing.B) {
	//benchmarkRandNumber(b,randNumber1)
	benchmarkRandNumber(b,randNumber3)
}

func benchmarkRandNumber(b *testing.B, randNumber func(int)[]int) {
	n := 9999
	for i := 0; i < b.N ; i++ {
		arr := randNumber(n)
		if !checkRandNumbers(arr,n) {
			b.Errorf("input %d, output: %v not expected\n",n,arr)
		}
	}
}



