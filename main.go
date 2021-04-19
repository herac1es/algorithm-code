package main

import "fmt"

func main() {
	a := int(89)
	fmt.Println(uint8(a))
}

func Apd(nums []int) {
	nums = nums[:len(nums)-1]
}
