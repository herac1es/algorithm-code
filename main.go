package main

import (
	"fmt"
	"math/rand"

	"github.com/iwillbesober/algorithm-code/classic"
)

func main() {
	pq := classic.New()
	for i := 0; i < 10; i++ {
		pq.Push(rand.Intn(100) + 1)
	}
	fmt.Println(pq.List())
}
