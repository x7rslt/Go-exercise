//练习2.5：使用 x&(x-1) 可以清除x最右边的非零位，利用该特点写一个 PopCount，然后评价它的性能。
package popcount

// 使用 x&(x-1) 可以清除x最右边的非零位
// 意思就是把最右边的一个1变成0，比如10是1010计算后是1000就是8，再计算1次就是0了
// 循环这个算法最后就会得到0，一次清一位，清理了几次就说明有多少位1
// 稀疏的情况下（1比较少），效率高
func PopCount(x uint64) int {
	var count int
	for x != 0 {
		x = x & (x - 1)
		count++
	}
	return count
}
