package main

func main() {

}

func funcMui(x, y int) (sum int, error) {
	return x + y, nil
}

/**
答：第⼆个返回值没有命名
解析：
在函数有多个返回值时，只要有⼀个返回值有命名，其他的也必须命名。如果有多个返回值必须加上括号();
如果只有⼀个返回值且命名也需要加上括号()。这⾥的第⼀个返回值有命名sum，第⼆个没有命名，所以错误。
*/
