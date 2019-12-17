package main

import "testing"

func TestLoop(t *testing.T) {
	//单元测试
	t.Log("Loop",Loop(uint64(32)))
}

func TestFactorial(t *testing.T) {
	t.Log("factorial",Factorial(uint64(32)))
}

func BenchmarkLoop(b *testing.B) {
	//基准测试 (性能测试)　ns是纳秒平均16.7ns
	for i:=0;i<b.N;i++{
		Loop(uint64(40))
	}
}

func BenchmarkFactorial(b *testing.B) {
	for i:=0;i<b.N;i++{
		Factorial(uint64(40))
	}
}
