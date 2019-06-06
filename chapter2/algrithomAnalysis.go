package main

import (
	"fmt"
	"math"
	"math/rand"
)

func main() {
	//N := 101
	//sushu := isSushu(N)
	//fmt.Println(N,"是素数吗？ ",sushu)
	//Eratosthenes(N)

	//fmt.Println("power1：",power1(2,50))
	//fmt.Println("power2：",power2(2,50))

	majorElement([]int{3,3,4,2,4,4,2,4,4})
	majorElement([]int{3,3,4,2,4,4,2,4})
	majorElement([]int{3,2,3,2,3,2,3})

}

/*
 2.19 找数组的主要元素：大小为N的数组，主要元素出现次数超过N/2
 首先找出主要元素的一个候选元素。这个候选元素是唯一有可能是主要元素的元素。
第二步确定是否该候选元素实际上就是主要元素。
为找出数组A的一个候选元，构造第二个数组B。比较A1和A2。如果它们相等，则取其中之一加到数组B中，否则什么也不做。然后比较A3，A4...直到读完该数组
然后递归得寻找数组B的候选元；它也是A的候选元

该算法之所以有效，是因为主要元素出现次数超过N/2
如    x 2 x 2 x 2 x 2 x
*/
func majorElement(A []int) int {
	B := candidateElements(A)
	if len(B) == 0 {
		fmt.Println("no major element")
		return -1
	}
	major := B[0]
	count := 0
	for _,v := range A  {
		if v == major  {
			count++
		}
	}
	if  count > len(A) / 2{
		fmt.Println("the major element is ",major)
		return major
	}
	fmt.Println("no major element")
	return -1
}

func candidateElements(A []int) []int {
	if len(A) <= 1 {
		return A
	}else{
		B := []int{}
		for i:= 0; i< len(A);i += 2  {
			if i+ 1 == len(A) {
				if A[i] == A[i - 2]  || A[i] == A[i - 1]{
					B = append(B, A[i])
				}
			}else {
				if A[i] == A[i + 1] {
					B = append(B, A[i])
				}
			}

		}
		return candidateElements(B)
	}
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

func sqrt(N int) int {
	return int(math.Sqrt(float64(N)))
}

/* 2.14 Eratosthenes筛：计算小于N的所有素数
	制作2到N的数组，找出最小的未被删除的整数i，打印i，然后删除1i,2i,3i..... 当i>sqrt(N)时退出
*/
func Eratosthenes(N int)  {
	arr := make([]int,N-2)
	for i := 2; i < N ; i++ {
		arr[i - 2] = i
	}
	for i := arr[0]; i <= sqrt(N) ;i = arr[0]{
		fmt.Println(i)
		for j := 0; j < len(arr); j ++{
			if arr[j] % i == 0 {
				if j >= len(arr) - 1 {
					arr = arr[0:j]
				}else {
					arr = append(arr[0:j],arr[j+1:len(arr)]...)
				}
			}
		}
	}
	for _,v := range arr {
		fmt.Println(v)
	}
}

/* 求幂 时间复杂度0(log N)*/
func power1(a,b int) int {
	if b == 0 {
		return 1
	}else if b % 2 == 0 {
		return power1(a * a, b / 2)
	}else {
		return power1(a * a, ( b - 1) / 2) * a
	}
}

/* 不用递归，写出快速求幂的程序 */
func power2(a,b int) int {
	if b == 0 {
		return 1
	}
	ops := []string{}
	//pow := a
	for b > 1 {
		//pow *= pow
		if b % 2 == 1 {
			b -= 1
			//pow *= a
			ops = append(ops,"*")
		}
		ops = append(ops,"pow")
		b /= 2;
	}
	p := a
	for i := len(ops) - 1; i >= 0; i--{
		switch ops[i] {
		case "pow":{
				p *= p
			}
		case "*":{
				p *= a
			}
		}
	}
	return p
}

func Abs(a int) int  {
	if a < 0{
		return -a
	}
	return a
}

// 求最大子序列乘积
// lastNegMul:记录上一个为负数的乘积。
// firstNegMul:记录第一个为负数的乘积
// 如果max mul为负数，则除以last or first使其为正。
func maxMulOfSubsequence1(arr []int) {
	maxMul, thisMul := 1, 1
	firstNegMul, lastNegMul := 1,0
	for _, v := range arr {
		thisMul *= v
		lastNegMul *= v	// 记录出现的负数，到下一个负数（不包含）之间的数的乘积。  循环结束后，该值为最后一个出现的负数及其以后数字的乘积。
		if v < 0 {
			lastNegMul = v;
		}
		if thisMul <0 && firstNegMul == 1 { // 记录第一个负数积
			firstNegMul = thisMul
		}
		if Abs(thisMul) > maxMul{
			maxMul = thisMul
		}
		if thisMul == 0 {
			thisMul = 1
			firstNegMul = 1
		}
	}
	if maxMul < 0 && firstNegMul < 0 && lastNegMul < 0{
		if firstNegMul < lastNegMul {
			maxMul /= lastNegMul
		}else {
			maxMul /= firstNegMul
		}
	}
	fmt.Println("max mul: ",maxMul)
}

func maxMulOfSubsequence2(arr []int) {
	maxMul, thisMul := 1, 1
	for i, _ := range arr {
		thisMul = 1
		for j:=i; j<len(arr); j++{
			thisMul *= arr[j]
			if thisMul > maxMul {
				maxMul = thisMul
			} else if arr[j] == 0 {
				thisMul = 1
			}
		}
	}
	fmt.Println("min mul: ",maxMul)
}

/* 求最小子序列和*/
func minSumOfSubsequence(arr []int)  {
	minSum, sum := 0,0

	for _,v := range arr{
		sum += v
		if sum < minSum {
			minSum = sum
		}else if (sum >= 0){
			sum = 0
		}
	}
	fmt.Println("min sum: ",minSum)
}


/* 生成前N个自然数的随机置换 */
func randNumber1(n int)  {
	arr := make([]int,n)
	for i:=0; i < n; {
		r := randInt(1,n)
		if findInt(arr,r) == -1 {
			arr[i] = r
			i++
		}
	}
	fmt.Println(arr)
}

func randNumber2(n int)  {
	arr := make([]int,n)
	used := make(map[int] int)
	for i:=0; i < n; {
		r := randInt(1,n)
		if _, ok := used[r]; ok == false {
			arr[i] = r
			used[r] = 1
			i++
		}
	}
	fmt.Println(arr)
}

func randNumber3(n int)  {
	arr := make([]int,n)
	for i:=0; i < n; i++ {
		arr[i] = i + 1
	}
	for i:= 0; i < n; i++{
		r := randInt(1,n) - 1
		arr[i], arr[r] = arr[r],arr[i]
	}
	fmt.Println(arr)
}

func findInt(arr []int,d int) int {
	for i,v := range arr{
		if v == d {
			return i
		}
	}
	return  -1
}


/* 随机数生成器：以相同的概率生成i到j之间的一个整数 */
func randInt(i,j int) int{
	min,d := i, j - i
	if j < i {
		min = j
		d = -d
	}
	d++
	if d == 0 {
		return min
	}

	return min + rand.Int() % d
}
