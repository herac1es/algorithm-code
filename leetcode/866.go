package leetcode

import (
	"math"
	"strconv"
)

//求出大于或等于 N 的最小回文素数。
//
// 回顾一下，如果一个数大于 1，且其因数只有 1 和它自身，那么这个数是素数。
//
// 例如，2，3，5，7，11 以及 13 是素数。
//
// 回顾一下，如果一个数从左往右读与从右往左读是一样的，那么这个数是回文数。
//
// 例如，12321 是回文数。
//
//
//
// 示例 1：
//
// 输入：6
//输出：7
//
//
// 示例 2：
//
// 输入：8
//输出：11
//
//
// 示例 3：
//
// 输入：13
//输出：101
//
//
//
// 提示：
//
//
// 1 <= N <= 10^8
// 答案肯定存在，且小于 2 * 10^8。
//
//
//
//
//
// Related Topics 数学
// 👍 61 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func primePalindrome(N int) int {
	if 8 <= N && N <= 11 {
		return 11
	}

	for i := N; i <= 200000000; {
		str := strconv.FormatInt(int64(i), 10)
		// 偶数位直接升位
		if len(str)%2 == 0 {
			i = int(math.Pow10(len(str)))
			continue
		}
		if i >= N && str == reverse(str) && isPrime(i) {
			return i
		}
		i++
	}
	return -1
}

func reverse(s string) string {
	if len(s) <= 1 {
		return s
	}
	ret := []byte(s)
	pre := 0
	last := len(s) - 1
	for pre < last {
		ret[pre], ret[last] = ret[last], ret[pre]
		pre++
		last--
	}
	return string(ret)
}

// 判断质数
func isPrime(n int) bool {
	if n < 2 || n%2 == 0 {
		return n == 2
	}
	for i := 3; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

//leetcode submit region end(Prohibit modification and deletion)
