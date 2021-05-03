package jzoffer

// 丑数一定是从前某一个丑树x2、x3、x5得来
// 用三个数分别记录响应倍率已经算过的索引位置
// 每轮从这三个倍率结果中选最小值
// 时间：O(n)
// 空间: O(n)
func nthUglyNumber(n int) int {
	dp := make([]int, n)
	dp[0] = 1
	two, three, five := 0, 0, 0
	for i := 1; i < n; i++ {
		v := min(dp, &two, &three, &five)
		dp[i] = v
	}
	return dp[n-1]
}

// 从已知丑数的2、3、5倍乘积中找到最小值
func min(nums []int, two, three, five *int) int {
	ret := 0
	multiTwo := nums[*two] * 2
	multiThree := nums[*three] * 3
	multiFive := nums[*five] * 5

	if multiTwo <= multiThree && multiTwo <= multiFive {
		ret = multiTwo
	} else if multiThree <= multiTwo && multiThree <= multiFive {
		ret = multiThree
	} else {
		ret = multiFive
	}

	// 因为同一个丑数只能出现一次，所以要将相同的都跳过
	if multiTwo == ret {
		*two++
	}
	if multiThree == ret {
		*three++
	}
	if multiFive == ret {
		*five++
	}
	return ret
}
