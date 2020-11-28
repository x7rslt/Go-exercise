//练习2.4：写一个用于统计位的 PopCount，它在其实际参数的64位上执行移位操作，每次判断最右边的位，进而实现统计功能。
//把它与快查表版本的性能进行对比。
package popcount

func PopCount(x uint64) int {
	var count int
	for x != 0 {
		if x&1 == 1 {
			count++
		}
		x >>= 1
	}
	return count
}
