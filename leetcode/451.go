package leetcode

// 给定一个字符串，请将字符串里的字符按照出现的频率降序排列。
//
// 示例 1:
//
//
// 输入:
// "tree"
//
// 输出:
// "eert"
//
// 解释:
// 'e'出现两次，'r'和't'都只出现一次。
// 因此'e'必须出现在'r'和't'之前。此外，"eetr"也是一个有效的答案。
//
//
// 示例 2:
//
//
// 输入:
// "cccaaa"
//
// 输出:
// "cccaaa"
//
// 解释:
// 'c'和'a'都出现三次。此外，"aaaccc"也是有效的答案。
// 注意"cacaca"是不正确的，因为相同的字母必须放在一起。
//
//
// 示例 3:
//
//
// 输入:
// "Aabb"
//
// 输出:
// "bbAa"
//
// 解释:
// 此外，"bbaA"也是一个有效的答案，但"Aabb"是不正确的。
// 注意'A'和'a'被认为是两种不同的字符。
//
// Related Topics 堆 哈希表
// 👍 230 👎 0
var gm map[rune]int

// leetcode submit region begin(Prohibit modification and deletion)
func frequencySort(s string) string {
	gm = make(map[rune]int)
	for _, v := range []rune(s) {
		gm[v]++
	}
	rs := make([]rune, 0, len(gm))
	for i := range gm {
		rs = append(rs, i)
	}
	size := len(rs)
	for i := size / 2; i >= 1; i-- {
		sink(rs, i, size)
	}
	ret := make([]rune, 0)
	for size > 0 {
		for i := 0; i < gm[rs[0]]; i++ {
			ret = append(ret, rs[0])
		}
		rs[0], rs[size-1] = rs[size-1], rs[0]
		size--
		sink(rs, 1, size)
	}
	return string(ret)
}

//  submit region end(Prohibit modification and deletion)

func sink(chars []rune, idx, size int) {
	for (idx << 1) <= size {
		t := idx << 1
		if t+1 <= size && (gm[chars[t+1-1]] > gm[chars[t-1]]) {
			t = t + 1
		}
		if gm[chars[t-1]] <= gm[chars[idx-1]] {
			break
		}
		chars[idx-1], chars[t-1] = chars[t-1], chars[idx-1]
		idx = t
	}
}
