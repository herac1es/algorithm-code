package leetcode

// 17.电话号码的字母组合
func letterCombinations(digits string) []string {
	if digits == "" {
		return []string{}
	}

	return letterCombRecur(digits, 0)

}

func letterCombRecur(digits string, index int) []string {

	if index == len(digits)-1 {
		return dict[digits[index]]
	}
	ret := make([]string, 0)
	last := letterCombRecur(digits, index+1)
	for _, v1 := range dict[digits[index]] {

		tmp := make([]string, 0, len(last))
		for _, v2 := range last {
			tmp = append(tmp, v1+v2)
		}
		ret = append(ret, tmp...)
	}
	return ret
}

var dict = map[byte][]string{
	'2': {"a", "b", "c"},
	'3': {"d", "e", "f"},
	'4': {"g", "h", "i"},
	'5': {"j", "k", "l"},
	'6': {"m", "n", "o"},
	'7': {"p", "q", "r", "s"},
	'8': {"t", "u", "v"},
	'9': {"w", "x", "y", "z"},
}
