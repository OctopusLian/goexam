package main

import "fmt"

func main() {
	s := make([]int, 5)
	s = append(s, 1, 2, 3)
	fmt.Println(s)

	s2 := make([]int, 0)
	s = append(s2, 1, 2, 3, 4)
	fmt.Println(s)
}
