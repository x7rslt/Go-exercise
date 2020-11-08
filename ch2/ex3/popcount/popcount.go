//练习2.3：使用循环重写 PopCount 来代替单个表达式。对比两个版本的效率（11.4节会展示如何系统性地对比不同实现的性能。）

package popcount

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}
func PopCount(x uint64) int {
	var count int
	for i := 0; i < 8; i++ {
		count += int(pc[byte(x>>(i*8))])
	}
	return count
}
