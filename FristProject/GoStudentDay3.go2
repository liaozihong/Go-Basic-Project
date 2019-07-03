package main

import (
	"fmt"
	"time"
)

func main() {
	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)
	//在数组上使用range将传入index和值两个变量。上面那个例子我们不需要使用该元素的序号，所以我们使用空白符"_"省略了。有时侯我们确实需要知道它的索引。
	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}
	//range 也可以用在map的键值对上
	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}
	//range也可以用来枚举Unicode字符串。第一个参数是字符的索引，第二个是字符（Unicode的值）本身。
	for i, c := range "go" {
		fmt.Println(i, c)
	}

	/*使用 MAP 函数*/
	var countryCapitalMap map[string]string /*创建集合 */
	countryCapitalMap = make(map[string]string)

	/* map插入key - value对,各个国家对应的首都 */
	countryCapitalMap["France"] = "巴黎"
	countryCapitalMap["Italy"] = "罗马"
	countryCapitalMap["Japan"] = "东京"
	countryCapitalMap["India "] = "新德里"

	/*使用键输出地图值 */
	for country := range countryCapitalMap {
		fmt.Println(country, "首都是", countryCapitalMap[country])
	}

	/*查看元素在集合中是否存在 */
	capital, ok := countryCapitalMap["American"] /*如果确定是真实的,则存在,否则不存在 */
	/*fmt.Println(capital) */
	/*fmt.Println(ok) */
	if ok {
		fmt.Println("American 的首都是", capital)
	} else {
		fmt.Println("American 的首都不存在")
	}
	//delete 函数
	delete(countryCapitalMap, "Japan")
	for country := range countryCapitalMap {
		fmt.Println(country, "删除后遍历："+countryCapitalMap[country])
	}

	//简单递归函数，没啥，思想等同就好
	var i int = 10
	fmt.Printf("%d 的阶乘是 %d\n", i, Factorial(uint64(i)))

	for index := 0; index < 10; index++ {
		fmt.Printf("%d \t", fibonacci(index))
	}
	fmt.Println()
	//类型转换， float32(target)
	var sum2 int = 17
	var count int = 5
	var mean float32

	mean = float32(sum2) / float32(count)
	fmt.Printf("mean 的值为: %f\n", mean)

	//测试接口
	var phone Phone
	phone = new(NokiaPhone)
	phone.call()
	phone = new(IPhone)
	phone.call()

	/*异常处理*/
	// 正常情况
	if result, errorMsg := Divide(100, 10); errorMsg == "" {
		fmt.Println("100/10 = ", result)
	}
	// 当被除数为零的时候会返回错误信息
	if _, errorMsg := Divide(100, 0); errorMsg != "" {
		fmt.Println("errorMsg is: ", errorMsg)
	}

	//Go 并发
	go say("你妹哦")
	say("啊哈")

	//channel 测试
	sliceTest := []int{7, 2, 8, -9, 4, 0}
	channelTest := make(chan int)
	go calculateSum(sliceTest[:len(sliceTest)/2], channelTest)
	go calculateSum(sliceTest[len(sliceTest)/2:], channelTest)
	fmt.Println(<-channelTest, <-channelTest) //从通道channelTest中接收

	//channel 缓冲区测试
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	fmt.Println(<-ch, <-ch)

	channalBuffer := make(chan int, 10)
	go testChannelBuffer(cap(channalBuffer), channalBuffer)
	// range 函数遍历每个从通道接收到的数据，因为 c 在发送完 10 个
	// 数据之后就关闭了通道，所以这里我们 range 函数在接收到 10 个数据
	// 之后就结束了。如果上面的 c 通道不关闭，那么 range 函数就不
	// 会结束，从而在接收第 11 个数据的时候就阻塞了。
	for i := range channalBuffer {
		fmt.Println(i)
	}
}

/*计算阶乘递归函数*/
func Factorial(n uint64) (result uint64) {
	if n > 0 {
		result = n * Factorial(n-1)
		return result
	}
	return 1
}

//计算斐波那契数列
func fibonacci(n int) int {
	if n < 2 {
		return n
	}
	return fibonacci(n-2) + fibonacci(n-1)

}

//定义一个接口
type Phone interface {
	call()
}

//定义一个接口实现
type NokiaPhone struct {
}

func (nokiaPhone NokiaPhone) call() {
	fmt.Println("I am Nokia, I can call you!")
}

//再定义一个接口实现
type IPhone struct {
}

func (iPhone IPhone) call() {
	fmt.Println("I am iPhone, I can call you!")
}

//定义一个DivideError结构
type DivideError struct {
	dividee int
	divider int
}

// 实现`error`接口
func (de *DivideError) Error() string {
	strFormat := `
		Cannot proceed, this divider is zero,
		dividee: %d
		divider: 0
	`
	return fmt.Sprintf(strFormat, de.dividee)
}

func Divide(varDividee int, varDivider int) (result int, errorMsg string) {
	if varDivider == 0 {
		dData := DivideError{
			dividee: varDividee,
			divider: varDivider,
		}
		errorMsg = dData.Error()
		return
	} else {
		return varDividee / varDivider, ""
	}
}

//测试并发方法
func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func calculateSum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum
}

func testChannelBuffer(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}
