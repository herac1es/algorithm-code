package leetcode

import "testing"

var a = " !\"#$%&'()*+,-./0123456789:;<=>?@ABCDEFGHIJKLMNOPQRSTUVWXYZ[\\]^_`abcdefghijklmnopqrstuvwxyz{|}~"

func Benchmark_lengthOfLongestSubstring(b *testing.B) {
	for i := 0; i < b.N; i++ {
		lengthOfLongestSubstring(a)
	}
}

func Benchmark_lengthOfLongestSubstringDp(b *testing.B) {
	for i := 0; i < b.N; i++ {
		lengthOfLongestSubstringDp(a)
	}
}
func Benchmark_lengthOfLongestSubstringMap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		lengthOfLongestSubstringMap(a)
	}
}

func Test_lengthOfLongestSubstringDp(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLongestSubstringDp(tt.args.s); got != tt.want {
				t.Errorf("lengthOfLongestSubstringDp() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_lengthOfLongestSubstringMap(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLongestSubstringMap(tt.args.s); got != tt.want {
				t.Errorf("lengthOfLongestSubstringMap() = %v, want %v", got, tt.want)
			}
		})
	}
}
