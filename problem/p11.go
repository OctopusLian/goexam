package main

import "fmt"

func main() {
	fmt.Println('abc' + '123')
	fmt.Println("abc" + "123")
	fmt.Println('123' + "abc")
	fmt.Println("abc%d", 123)
}

/**
答：B、D
解析：
在Golang中字符串⽤双引号，字符⽤单引号字符串连接除了以上两种连接⽅式，
还有 strings.Join() 、 buffer.WriteString() 等
*/
