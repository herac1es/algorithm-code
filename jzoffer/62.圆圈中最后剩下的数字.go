package jzoffer

// 数学推导法
// f(n,m) = (f(n-1,m)+m)%n f(1,m) = 0

func lastRemaining(n int, m int) int {

	last := 0
	for i := 2; i <= n; i++ {
		last = (last + m) % i
	}
	return last
}

// 模拟：超时
// func lastRemaining(n int, m int) int {
//     queue := make([]int, n)
//     for i := range queue {
//         queue[i] = i
//     }

//     idx := 0
//     for len(queue) > 1 {
//         idx = (m-1+idx)%len(queue)
//         queue = append(queue[:idx], queue[idx+1:]...)
//     }
//     return queue[0]
// }
