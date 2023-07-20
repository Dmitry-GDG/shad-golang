package bench

import (
	"fmt"
	"testing"
)

func BenchmarkSample(b *testing.B) {
	b.ReportAllocs() // бенчмарк будет считать аллокации памяти
	b.SetBytes(2) // установить кол-во байт на одну операцию
	for i := 0; i < b.N; i++ { // код запускактся b.N раз, b.N сам подкручивает кол-во, пока не будет уверен
		if x := fmt.Sprintf("%d", 42); x != "42" {
			b.Fatalf("Unexpected string: %s", x)
		}
	}
}

// func BenchmarkSample2(b *testing.B) {
// 	b/Setbytes(2)
// 	for i:= 0; i < b.N; i++ {
// 		if x := fmt.Sprintf("%d", 42); x != "42" {
// 			b.Fatalf("Unexpected string: %s", x)
// 		}
// 	}
// }

// go test -v -bench=BenchmarkSample$ ./format_test.go

