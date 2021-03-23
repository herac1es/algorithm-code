package main

import "fmt"

func main() {
	a := make(map[int]struct{}, 10)
	fmt.Println(len(a))
}
