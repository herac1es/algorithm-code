package jzoffer

// 时间: O(n!)
// 空间: O(n2)
func permutation(s string) []string {
	visited := make(map[int]struct{})
	ret := make([]string, 0)
	permutationRecur(s, visited, "", &ret)
	return ret
}

func permutationRecur(s string, visited map[int]struct{}, item string, ret *[]string) {

	if len(item) == len(s) {
		*ret = append(*ret, item)
		return
	}

	localVisited := make(map[byte]struct{}, len(s))
	for idx, v := range []byte(s) {
		if _, ok := localVisited[v]; ok {
			continue
		}
		if _, ok := visited[idx]; ok {
			continue
		}
		localVisited[v] = struct{}{}
		visited[idx] = struct{}{}
		permutationRecur(s, visited, item+string(v), ret)
		delete(visited, idx)
	}
}
