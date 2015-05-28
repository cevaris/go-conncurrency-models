package sum

import (
	"testing"
)

func numGenerator(n int) []int {
	var s []int = make([]int, n)
	for i := 0; i < n; i++ {
		s[i] = i
	}
	return s
}

func TestNumGenerator(t *testing.T) {
	expected := []int{0, 1, 2, 3, 4}
	actual := numGenerator(5)

	for i, v := range(actual) {
		if expected[i] != v {
			t.Error(actual, "not equal", expected)
		}
	}
}

func TestSimpleSum(t *testing.T) {
	var expected int64 = 45
	data := numGenerator(10)
	var actual int64 = SimpleSum(&data)
	if actual != expected {
		t.Error(actual, "not equal", expected)
	}
}

func benchmarkSimpleSum(size int, b *testing.B) {
	data := numGenerator(size)
	for n := 0; n < b.N; n++ {
		SimpleSum(&data)
	}
}

func BenchmarkSimpleSum1K(b *testing.B) {
	benchmarkSimpleSum(10*100, b)
}
func BenchmarkSimpleSum10K(b *testing.B) {
	benchmarkSimpleSum(100*100, b)
}
func BenchmarkSimpleSum100K(b *testing.B) {
	benchmarkSimpleSum(10*100*100, b)
}
func BenchmarkSimpleSum1M(b *testing.B) {
	benchmarkSimpleSum(100*100*100, b)
}
func BenchmarkSimpleSum10M(b *testing.B) {
	benchmarkSimpleSum(10*100*100*100, b)
}
func BenchmarkSimpleSum100M(b *testing.B) {
	benchmarkSimpleSum(100*100*100*100, b)
}



func benchmarkSyncSum(size int, b *testing.B) {
	data := numGenerator(size)
	for n := 0; n < b.N; n++ {
		SyncSum(&data, 1, 1, 1)
	}
}

func BenchmarkSyncSum1K(b *testing.B) {
	benchmarkSyncSum(10*100, b)
}
func BenchmarkSyncSum10K(b *testing.B) {
	benchmarkSyncSum(100*100, b)
}
func BenchmarkSyncSum100K(b *testing.B) {
	benchmarkSyncSum(10*100*100, b)
}
func BenchmarkSyncSum1M(b *testing.B) {
	benchmarkSyncSum(100*100*100, b)
}
func BenchmarkSyncSum10M(b *testing.B) {
	benchmarkSyncSum(10*100*100*100, b)
}
func BenchmarkSyncSum100M(b *testing.B) {
	benchmarkSyncSum(100*100*100*100, b)
}
