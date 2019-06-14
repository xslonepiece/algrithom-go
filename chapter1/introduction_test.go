package main

import "testing"

func TestFindNum(t *testing.T)  {
	tests := []struct{
		arr []int
		k int
		res int
	}{
		{[]int{8,1,3,5,2,9},1,9},
		{[]int{8,1,3,5,2,9},2,8},
		{[]int{8,1,3,5,2,9},3,5},
		{[]int{8,1,3,5,2,9},4,3},
		{[]int{8,1,3,5,2,9},5,2},

		{[]int{8,1,3,5,2,9},100,-1},
		{[]int{},1,-1},
	}
	for _,tt := range tests{
		actual := findNum(tt.arr,tt.k)
		if actual != tt.res {
			 t.Errorf("arr：%v，the %d max num is %d, but expect is %d\n",tt.arr,tt.k,actual,tt.res)
		}
	}
}