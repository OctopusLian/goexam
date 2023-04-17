package main

import "fmt"

func main() {
	s1 := []int{1, 2, 3}
	s2 := []int{4, 5}
	s1 = append(s1, s2)
	fmt.Println(s1)
}

/**
答：不能通过
解析：
append() 的第⼆个参数不能直接使⽤ slice ，需使⽤ ... 操作符，将⼀个切⽚追加到另⼀个切⽚上：
append(s1, s2...) 。或者直接跟上元素，形如： append(s1, 1, 2, 3) 。
*/
