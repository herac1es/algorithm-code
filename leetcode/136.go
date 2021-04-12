package leetcode

// 136.只出现一次的数据
// 给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。
//
// 说明：
//
// 你的算法应该具有线性时间复杂度。 你可以不使用额外空间来实现吗？
//
// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/single-number
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

// 利用 ^ 异或的交换律
// 时间：O(n)
// 空间：O(1)
func singleNumber(nums []int) int {
	ret := 0
	for _, v := range nums {
		ret ^= v
	}
	return ret
}

// hash解法
// 时间：O（n）
// 空间：O（n)
func singleNumberHash(nums []int) int {
	hmap := make(map[int]struct{}, len(nums))
	for _, v := range nums {
		if _, ok := hmap[v]; !ok {
			hmap[v] = struct{}{}
			continue
		}
		delete(hmap, v)
	}
	ret := 0
	for k := range hmap {
		ret = k
	}
	return ret
}
