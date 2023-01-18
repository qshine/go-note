package split

import "testing"

// 性能测试
// 执行必须带 -bench 参数, 如: go unittest -bench=Split

func BenchmarkSplit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Split("沙河有沙又有河", "沙")
	}
}
