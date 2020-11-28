//运行Benchmark测试需要确保包内只有一个main函数
//go test -bench=. 执行测试
package ex3a2

import "testing"

func BenchmarkEcho1(b *testing.B){
	Echo1()
}
func BenchmarkEcho2(b *testing.B){
	Echo2()
}