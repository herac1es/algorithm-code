package main

import "fmt"

func main() {
	a := make([]int, 0)
	a = append(a, get()...)
	fmt.Println(a)
}

func get() []int {
	return nil
}
