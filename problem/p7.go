package main

import "fmt"

var (
	size := 1024
	max_size = size * 2
)

func main() {
	fmt.Println(size, max_size)
}

/**
答：不能通过
解析：
这道题的主要知识点是变量的简短模式，形如：x := 100 。但这种声明⽅式有限制：
1. 必须使⽤显示初始化；
2. 不能提供数据类型，编译器会⾃动推导；
3. 只能在函数内部使⽤简短模式；
 */