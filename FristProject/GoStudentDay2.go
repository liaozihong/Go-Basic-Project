package main

import "fmt"

/* 声明全局变量 */
var a int = 20

type Books struct {
	title   string
	author  string
	subject string
	book_id int
}

func main() {

	/**初始化数组**/
	var balance = [5]float32{1000.0, 2.0, 3.4, 7.0, 50.0}
	fmt.Println(balance)
	balance = [...]float32{1000.0, 2.0, 3.4, 7.0, 50.0}
	fmt.Println(balance)
	//访问数组元素并赋值
	var salary float32 = balance[3]
	fmt.Println(salary)

	/**for循环使用方式，1**/
	for i, x := range balance {
		fmt.Printf("for循环no.1 ,第 %d 位 值 为：%f\n", i, x)
	}

	/**for循环使用方式，2**/
	for i := 0; i < len(balance); i++ {
		fmt.Printf("for循环no.2 ,第 %d 位 值 为：%f\n", i, balance[i])
	}

	/**for循环使用方式，3**/
	var i int
	for i < len(balance) {
		fmt.Printf("for循环no.3 ,第 %d 位 值 为：%f\n", i, balance[i])
		i++
	}

	/* nextNumber 为一个函数，函数 i 为 0 */
	nextNumber := getSequence()

	/* 调用 nextNumber 函数，i 变量自增 1 并返回 */
	fmt.Println(nextNumber())
	fmt.Println(nextNumber())
	fmt.Println(nextNumber())

	/* 创建新的函数 nextNumber1，并查看结果 */
	nextNumber1 := getSequence()
	fmt.Println(nextNumber1())
	fmt.Println(nextNumber1())

	/* main 函数中声明局部变量 */
	var a int = 10
	var b int = 20
	var c int = 0

	fmt.Printf("main()函数中 a = %d\n", a)
	c = sum(a, b)
	fmt.Printf("main()函数中 a = %d\n", a)
	fmt.Printf("main()函数中 c = %d\n", c)

	//声明指针变量
	var ip *int
	ip = &a /**指正变量的存储地址*/
	//查看变量地址值
	fmt.Printf("变量的地址：%x\n", &a)
	/* 指针变量的存储地址 */
	fmt.Printf("ip 变量储存的指针地址: %x\n", ip)

	/* 使用指针访问值 */
	fmt.Printf("*ip 变量的值: %d\n", *ip)

	/*Go空指针nil 指针也称为空指针。
	nil在概念上和其它语言的null、None、nil、NULL一样，都指代零值或空值。*/
	var ptr *int
	fmt.Printf("ptr 的值为%x\n", ptr)

	//指针数组
	pointArr := []int{10, 100, 200}
	var index int
	var ptrs [3]*int

	for index = 0; index < 3; index++ {
		ptrs[index] = &pointArr[index] /* 整数地址赋值给指针数组 */
	}
	for index = 0; index < 3; index++ {
		fmt.Printf("a[%d] = %d\n", index, *ptrs[index])
	}
	//结构体使用，相当于对象
	// 创建一个新的结构体
	fmt.Println(Books{"Go 语言", "www.runoob.com", "Go 语言教程", 6495407})

	// 也可以使用 key => value 格式
	fmt.Println(Books{title: "Go 语言", author: "www.runoob.com", subject: "Go 语言教程", book_id: 6495407})

	// 忽略的字段为 0 或 空
	fmt.Println(Books{title: "Go 语言", author: "www.runoob.com"})
	var Book1 Books /* 声明 Book1 为 Books 类型 */
	var Book2 Books /* 声明 Book2 为 Books 类型 */

	/* book 1 描述 */
	Book1.title = "Go 语言"
	Book1.author = "www.runoob.com"
	Book1.subject = "Go 语言教程"
	Book1.book_id = 6495407

	/* book 2 描述 */
	Book2.title = "Python 教程"
	Book2.author = "www.runoob.com"
	Book2.subject = "Python 语言教程"
	Book2.book_id = 6495700

	/* 打印 Book1 信息 */
	printBook(Book1)

	/* 打印 Book2 信息 */
	printBook(Book2)

	/*结构体指针使用*/
	/* 打印 Book1 信息 */
	printBook2(&Book1)

	/* 打印 Book2 信息 */
	printBook2(&Book2)

	/*语言分片 Slice，定义语法如下：
		var slice1 []type = make([]type, len)
		也可以简写为
		slice1 := make([]type, len)
		也可以指定容量，其中capacity为可选参数。
	make([]T, length, capacity)
	*/
	var number = make([]int, 3, 5)
	printSlice(number)

	//切片的截取
	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	printSlice(numbers)
	//打印子切片从索引1(包含)到索引4(不包含)
	fmt.Println("numbers[1:4] == ", numbers[1:4])
	/* 默认下限为 0*/
	fmt.Println("numbers[:3] ==", numbers[:3])

	/* 默认上限为 len(s)*/
	fmt.Println("numbers[4:] ==", numbers[4:])

	number1 := make([]int, 0, 5)
	printSlice(number1)

	/* 打印子切片从索引  0(包含) 到索引 2(不包含) */
	number2 := numbers[:2]
	printSlice(number2)

	/* 打印子切片从索引 2(包含) 到索引 5(不包含) */
	number3 := numbers[2:5]
	printSlice(number3)

}

func printSlice(x []int) {
	fmt.Printf("len=%d,cap=%d slice=%v\n", len(x), cap(x), x)
}

/* 函数定义-两数相加 */
func sum(a, b int) int {
	a = a + 1
	fmt.Printf("sum() 函数中 a = %d\n", a)
	fmt.Printf("sum() 函数中 b = %d\n", b)
	return a + b
}

func getSequence() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}
func printBook(book Books) {
	fmt.Printf("Book title : %s\n", book.title)
	fmt.Printf("Book author : %s\n", book.author)
	fmt.Printf("Book subject : %s\n", book.subject)
	fmt.Printf("Book book_id : %d\n", book.book_id)
}

func printBook2(book *Books) {
	fmt.Printf("Book title pointer : %s\n", book.title)
	fmt.Printf("Book author pointer: %s\n", book.author)
	fmt.Printf("Book subject pointer: %s\n", book.subject)
	fmt.Printf("Book book_id pointer: %d\n", book.book_id)
}
