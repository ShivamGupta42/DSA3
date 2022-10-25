package benchmark

import "testing"

func BenchmarkCallMethod1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CallMethod1()
	}
}

func BenchmarkCallMethod2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CallMethod2()
	}
}

func BenchmarkCallMethod3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CallMethod3()
	}
}
