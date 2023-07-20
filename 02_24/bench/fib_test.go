package bench

import "testing"

func Fib(n int) int {
	if n < 2 {
		return n
	}
	return Fib(n-1) + Fib(n-2)
}

func BenchmarkFib10(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Fib(10)
	}
}

// go test -v -bench=BenchmarkFib10$ ./fib_test.go
// go test -v -bench=BenchmarkFib10$ -cpuprofile=/tmp/cpu.out ./fib_test.go
