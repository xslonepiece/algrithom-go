package main

import (
	"fmt"
	"math"
)

func main() {
	N := 6
	sushu := isSushu(N)
	fmt.Println(N,"是素数吗？ ",sushu)
}

/* 编写一个程序来确定正整数N是否是素数  要求时间复杂度O（开根号N）
素数只能被1和自身整除
判断其是否能被从2到N-1之间的数整除
*/
func isSushu(N int) bool {
	if N % 2 == 0 {
		return false
	}
	for i:= 3; i <= int(math.Sqrt(float64(N))); i += 2{
		if N % i == 0 {
			return false
		}
	}
	return true
}