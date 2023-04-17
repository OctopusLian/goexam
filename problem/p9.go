package main

import "fmt"

type MyInt1 int
type MyInt2 = int

func main() {
	var i int = 0
	var i1 MyInt1 = i
	var i2 MyInt2 = i
	fmt.Println(i1, i2)
}

/**
答：不能通过
解析：
这道题考的是 类型别名 与 类型定义 的区别
第5⾏代码是基于类型 int 创建了新类型 MyInt1，第6⾏代码是创建了int的类型别名 MyInt2，注意类型别名的定义是 = 。
所以，第10⾏代码相当于是将int类型的变量赋值给MyInt1类型的变量，Go是强类型语⾔，编译当然不通过；
⽽MyInt2只是int的别名，本质上还是int，可以赋值。
第10⾏代码的赋值可以使⽤强制类型转换 var i1 MyInt1 = MyInt1(i)
*/
