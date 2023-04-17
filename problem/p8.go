package main

import "fmt"

func main() {
	sn1 := struct {
		age
		int
		name string
	}{age: 11, name: "qq"}
	sn2 := struct {
		age
		int
		name string
	}{age: 11, name: "11"}

	if sn1 == sn2 {
		fmt.Println("sn1 == sn2")
	}
	sm1 := struct {
		age int
		m   map[string]string
	}{age: 11, m: map[string]string{"a": "1"}}
	sm2 := struct {
		age int
		m   map[string]string
	}{age: 11, m: map[string]string{"a": "1"}}
	if sm1 == sm2 {
		fmt.Println("sm1 == sm2")
	}
}

/**
答：不能通过,invalid operation: sm1 == sm2
解析：
考点是结构体的⽐较，有⼏个需要注意的地⽅：
1. 结构体只能⽐较是否相等，但是不能⽐较⼤⼩；
2. 想同类型的结构体才能进⾏⽐较，结构体是否相同不但与属性类型有关，还与属性顺序相关；
3. 如果struct的所有成员都可以⽐较，则该struct就可以通过==或!=进⾏⽐较是否相同，⽐较时逐个项
进⾏⽐较，如果每⼀项都相等，则两个结构体才相等，否则不相等；

那有什么是可以⽐较的呢？
常⻅的有bool、数值型、字符、指针、数组等

不能⽐较的有
slice、map、函数
*/
