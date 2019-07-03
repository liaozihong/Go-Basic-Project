package main

import (
	"fmt"
	"math"
	"unsafe"
)

const (
	a = "abc"
	b = len(a)
	c = unsafe.Sizeof(a)
)

func main() {
	a := 10

	if a < 20 {
		fmt.Printf("a 小于 20\n")
	}
	fmt.Printf("a 的 值为： %d\n", a)

	var i, j int
	for i = 2; i < 100; i++ {
		for j = 2; j <= (i / j); j++ {
			if i%j == 0 {
				break
			}
		}
		if j > (i / j) {
			fmt.Printf("%d 是素数!\n", i)
		}
	}

	var a1 int = 4
	var b int32
	var c float32
	var ptr *int

	/* 运算符实例 */
	fmt.Printf("第 1 行 - a1 变量类型为 = %T\n", a1)
	fmt.Printf("第 2 行 - b 变量类型为 = %T\n", b)
	fmt.Printf("第 3 行 - c 变量类型为 = %T\n", c)

	/*  & 和 * 运算符实例 */
	ptr = &a1 /* 'ptr' 包含了 'a' 变量的地址 */
	fmt.Printf("a 的值为  %d\n", a1)
	fmt.Printf("*ptr 为 %d\n", *ptr)

	//基于goto进行遍历
	a3 := 10
Loop:
	for a3 < 20 {
		if a3 == 15 {
			a3 = a3 + 1
			goto Loop
		}
		fmt.Printf("a的值为：%d\n", a3)
		a3++
	}
	//面积计算
	const LENGTH int = 10
	const WIDTH int = 5
	var area int
	const constA, constB, constC = 1, false, "str"
	area = LENGTH * WIDTH
	fmt.Printf("面积为:%d", area)
	println()
	println(constA, constB, constC)

	//多返回值
	resultA, resultB := swap("GOOGLE", "Runoob")
	fmt.Println(resultA, resultB)

	getSquareRoot := func(x float64) float64 {
		return math.Sqrt(x)
	}
	fmt.Println(getSquareRoot(4))
}

func swap(x, y string) (string, string) {
	return y, x
}
