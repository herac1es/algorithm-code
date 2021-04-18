package main

import "fmt"

func main() {
	a := []int{1, 2, 3}
	b := append([][]int{}, a)
	fmt.Println(a, b)
	a[1] = 8
	fmt.Println(a, b)
}

func Apd(nums []int) {
	nums = nums[:len(nums)-1]
}
