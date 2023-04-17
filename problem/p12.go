package main

import "fmt"

const (
	x = iota
	_
	y
	z = "zz"
	k
	p = iota
)

func main() {
	fmt.Println(x, y, z, k, p)
}

/**
0 2 zz zz 5
解析：
iota初始值为0，所以x为0，_表示不赋值，但是iota是从上往下加1的，所以y是2，z是“zz”,k和上⾯⼀个
同值也是“zz”,p是iota,从上0开始数他是5
*/
