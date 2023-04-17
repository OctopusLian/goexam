package main

import "fmt"

func main() {
	a := []int{7, 8, 9}
	fmt.Printf("%+v\n", a)
	ap(a)
	fmt.Printf("%+v\n", a)
	app(a)
	fmt.Printf("%+v\n", a)
}

func ap(a []int) {
	a = append(a, 10)
}

func app(a []int) {
	a[0] = 1
}

/**
[7 8 9]
[7 8 9]
[1 8 9]
解析：
因为append导致底层数组重新分配内存了，append中的a这个alice的底层数组和外⾯不是⼀个，
并没有改变外⾯的。
*/
