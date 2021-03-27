package leetcode

import "strconv"

// 给定一个字符串数组，将字母异位词组合在一起。字母异位词指字母相同，但排列不同的字符串。
//
// 示例:
//
// 输入: ["eat", "tea", "tan", "ate", "nat", "bat"]
// 输出:
// [
//  ["ate","eat","tea"],
//  ["nat","tan"],
//  ["bat"]
// ]
//
// 说明：
//
//
// 所有输入均为小写字母。
// 不考虑答案输出的顺序。
//
// Related Topics 哈希表 字符串
// 👍 697 👎 0

// leetcode submit region begin(Prohibit modification and deletion)
func groupAnagrams(strs []string) [][]string {
	ret := make([][]string, 0)
	m := make(map[string]int)
	for i := range strs {
		if v, ok := m[encode(strs[i])]; ok {
			ret[v] = append(ret[v], strs[i])
		} else {
			ret = append(ret, []string{strs[i]})
			m[encode(strs[i])] = len(ret) - 1
		}
	}
	return ret
}

func encode(s string) string {
	arr := [26]int{0}
	for i := range s {
		arr[s[i]-97]++
	}
	ret := ""
	for k, v := range arr {
		if v > 0 {
			ret += string(k) + strconv.FormatInt(int64(v), 10)
		}
	}
	return ret
}

// leetcode submit region end(Prohibit modification and deletion)
