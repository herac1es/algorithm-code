package jzoffer

import "strconv"

// 1 位： 0-9  10个
// 2 位： 10-99 90个
// 3 位： 100-999 900个.
// 4 位:  1000-9999 9000个
func findNthDigit(n int) int {
	if n < 10 {
		return n
	}
	start := 1
	b := 1
	c := 9
	for n > c {
		n -= c
		b++
		start *= 10
		c = 9 * start * b
	}
	num := start + (n-1)/b
	return int(strconv.Itoa(num)[(n-1)%b] - '0')
}
