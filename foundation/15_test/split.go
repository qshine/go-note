package split

import (
	"strings"
)

func Split(s, seq string) (result []string) {
	// 使用该行优化. 原因是把result初始化为了一个容量足够大的切片, 节省了每次append分配内存
	result = make([]string, 0, strings.Count(s, seq)+1)

	i := strings.Index(s, seq)
	for i > -1 {
		result = append(result, s[:i])
		s = s[i+len(seq):]
		i = strings.Index(s, seq)
	}
	result = append(result, s)
	return
}
