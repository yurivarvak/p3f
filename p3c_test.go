package main

import "testing"

func Test_100(t *testing.T) {
	p := primes(10000)
	s, _ := prime3factor(100, p)
	if s != 520 {
		t.Error("fail 100")
	}
	s, _ = prime3factor(212, p)
	if s != 4209 {
		t.Error("fail 212")
	}
	s, _ = prime3factor(3732, p)
	if s != 2560656 {
		t.Error("fail 3732")
	}
	s, _ = prime3factor(10000, p)
	if s != 19186879 {
		t.Error("fail 10000")
	}
}

func Benchmark_1000000(b *testing.B) {
	n := 1000000
	p := primes(n)
	//	b.ResetTimer()
	prime3factor(n, p)
}
