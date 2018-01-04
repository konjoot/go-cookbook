package search

import "testing"

var (
	dataMap10     dMap
	dataSlice10   dSlice
	dataMap100    dMap
	dataSlice100  dSlice
	dataMap1000   dMap
	dataSlice1000 dSlice
)

func init() {
	dataMap10 = newDMap(10)
	dataSlice10 = newDSlice(10)
	dataMap100 = newDMap(100)
	dataSlice100 = newDSlice(100)
	dataMap1000 = newDMap(1000)
	dataSlice1000 = newDSlice(1000)
}

func benchmarkMap10(key string, b *testing.B) {
	for n := 0; n < b.N; n++ {
		dataMap10.includes(key)
	}
}

func benchmarkSlice10(key string, b *testing.B) {
	for n := 0; n < b.N; n++ {
		dataSlice10.includes(key)
	}
}

func benchmarkMap100(key string, b *testing.B) {
	for n := 0; n < b.N; n++ {
		dataMap100.includes(key)
	}
}

func benchmarkSlice100(key string, b *testing.B) {
	for n := 0; n < b.N; n++ {
		dataSlice100.includes(key)
	}
}

func benchmarkMap1000(key string, b *testing.B) {
	for n := 0; n < b.N; n++ {
		dataMap1000.includes(key)
	}
}

func benchmarkSlice1000(key string, b *testing.B) {
	for n := 0; n < b.N; n++ {
		dataSlice1000.includes(key)
	}
}

// 10 elements
// map
func BenchmarkMap10Start(b *testing.B) {
	benchmarkMap10("0", b)
}

func BenchmarkMap10Middle(b *testing.B) {
	benchmarkMap10("4", b)
}

func BenchmarkMap10End(b *testing.B) {
	benchmarkMap10("9", b)
}

func BenchmarkMap10Out(b *testing.B) {
	benchmarkMap10("10", b)
}

// slice
func BenchmarkSlice10Start(b *testing.B) {
	benchmarkSlice10("0", b)
}

func BenchmarkSlice10Middle(b *testing.B) {
	benchmarkSlice10("4", b)
}

func BenchmarkSlice10End(b *testing.B) {
	benchmarkSlice10("9", b)
}

func BenchmarkSlice10Out(b *testing.B) {
	benchmarkSlice10("10", b)
}

// 100 elements
// map
func BenchmarkMap100Start(b *testing.B) {
	benchmarkMap100("0", b)
}

func BenchmarkMap100Middle(b *testing.B) {
	benchmarkMap100("49", b)
}

func BenchmarkMap100End(b *testing.B) {
	benchmarkMap100("99", b)
}

func BenchmarkMap100Out(b *testing.B) {
	benchmarkMap100("100", b)
}

// slice
func BenchmarkSlice100Start(b *testing.B) {
	benchmarkSlice100("0", b)
}

func BenchmarkSlice100Middle(b *testing.B) {
	benchmarkSlice100("49", b)
}

func BenchmarkSlice100End(b *testing.B) {
	benchmarkSlice100("99", b)
}

func BenchmarkSlice100Out(b *testing.B) {
	benchmarkSlice100("100", b)
}

// 1000 elements
// map
func BenchmarkMap1000Start(b *testing.B) {
	benchmarkMap1000("0", b)
}

func BenchmarkMap1000Middle(b *testing.B) {
	benchmarkMap1000("499", b)
}

func BenchmarkMap1000End(b *testing.B) {
	benchmarkMap1000("999", b)
}

func BenchmarkMap1000Out(b *testing.B) {
	benchmarkMap1000("1000", b)
}

// slice
func BenchmarkSlice1000Start(b *testing.B) {
	benchmarkSlice1000("0", b)
}

func BenchmarkSlice1000Middle(b *testing.B) {
	benchmarkSlice1000("499", b)
}

func BenchmarkSlice1000End(b *testing.B) {
	benchmarkSlice1000("999", b)
}

func BenchmarkSlice1000Out(b *testing.B) {
	benchmarkSlice1000("1000", b)
}
