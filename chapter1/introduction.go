package main

import (
	"fmt"
	"strings"
)

func main() {
	//arr := []int{8,1,3,5,2,9}
	//k := 3
	//fmt.Print(arr)
	//fmt.Printf("中第%d个最大值是：%d\n",k, findNum(arr,k))

	//printNumber(-8883)
	readCFileName("#include helloworld.h #include test.go #include #include s s")
}

/* 对于给定的N个数，确定其中第k个最大者 */
func findNum(arr []int, k int) int {
	n := len(arr)
	if k >= n {
		return -1
	}
	//plan 1
	bubblesort(arr)

	// plan 2
	//for i := 0; i < k; i++ {
	//	sorted := true
	//	for j := 0; j < n - i - 1; j++ {
	//		if arr[j] > arr[j + 1]{
	//			arr[j],arr[j + 1] = arr[j + 1] ,arr[j]
	//			sorted = false
	//		}
	//	}
	//	if sorted == true {
	//		break;
	//	}
	//}
	return arr[n - k]
}

/* 冒泡排序 */
func bubblesort(arr []int)  {
	for i := 0; i < len(arr); i++ {
		sorted := true
		for j := 0; j < len(arr) - i - 1; j++ {
			if arr[j] > arr[j + 1]{
				arr[j],arr[j + 1] = arr[j + 1] ,arr[j]
				sorted = false
			}
		}
		if sorted == true {
			break;
		}
	}
}

type Direction struct {
}

const (
	h = iota
	v
	d			//斜
)

type info struct {
	x,y,direction int
	s string
}


/*  字谜问题：从存放字母的二维数组中，横向，纵向，对角找出单词 */
func findEnglishWorld(arr []string)  {

}


/* 1.3 使用printDigit函数，编写一个过程以输出任意实数（可以是负的） */
func printNumber(num int){
	if abs(num) >= 10{
		printNumber(num / 10)
		printDigit(abs(num) % 10)
	}else {
		printDigit(num % 10)
	}
}

func abs(num int) int {
	if num <0 {
		return  -num
	}
	return  num
}

func printDigit(digit int)  {
	fmt.Print(digit)
}


/* 1.4 从#include filename中读取filename，支持嵌套
(因没有输入示例和输出示例，只能按自己理解的简单实现，不一定符合题意)*/
const prefix  = "#include "

func readCFileName(s string)  {
	idx := strings.Index(s,prefix)
	if idx >= 0{
		if idx > 0 {
			fmt.Println(s[0:idx])
		}
		readCFileName(s[idx + len(prefix):len(s)])
	}else{
		fmt.Println(s)
	}
}