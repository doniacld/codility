package mix

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFactorials(t *testing.T) {
	res := Factorials(3)
	assert.Equal(t, "6, 2, 1", res)
	res = Factorials(0)
	assert.Equal(t, "", res)
	res = Factorials(15)
	assert.Equal(t, "1307674368000, 87178291200, 6227020800, 479001600, 39916800, 3628800, 362880, 40320, 5040, 720, 120, 24, 6, 2, 1", res)
}


var result string

var (
	in1  = 1
	in5  = 5
	in10 = 10
	in15 = 15
)

func benchmarkFactorials(input int, b *testing.B) {
	var r string

	for n := 0; n < b.N; n++ {
		r = Factorials(input)
	}
	result = r
}

func BenchmarkFactorial1(b *testing.B)  { benchmarkFactorials(in1, b) }
func BenchmarkFactorial5(b *testing.B)  { benchmarkFactorials(in5, b) }
func BenchmarkFactorial10(b *testing.B) { benchmarkFactorials(in10, b) }
func BenchmarkFactorial15(b *testing.B) { benchmarkFactorials(in15, b) }

/* Benchmark result

goos: linux
goarch: amd64
pkg: github.com/doniacld/codility/mix
BenchmarkFactorial1
BenchmarkFactorial1-8    	10625452	       101 ns/op
BenchmarkFactorial5
BenchmarkFactorial5-8    	 1839907	       640 ns/op
BenchmarkFactorial10
BenchmarkFactorial10-8   	  806451	      1260 ns/op
BenchmarkFactorial15
BenchmarkFactorial15-8   	  753147	      1490 ns/op
PASS
*/
