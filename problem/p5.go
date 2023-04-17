package main

import "fmt"

func main() {
	list := new([]int)
	list = append(list, 1)
	fmt.Println(list)
}

/**
答：不能通过
解析：
不能通过编译， new([]int) 之后的 list 是⼀个 *int[] 类型的指针，不能对指针执⾏ append 操作。
可以使⽤ make() 初始化之后再⽤。同样的， map 和 channel 建议使⽤ make() 或字⾯量的⽅式初始化，不要⽤ new 。
*/
