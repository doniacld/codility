package genomicrangequery

import (
	"testing"

	"github.com/magiconair/properties/assert"
)

// TestGenomicRangeQuery tests GenomicRangeQuery method
func TestGenomicRangeQuery(t *testing.T) {
	tt := []struct {
		description string
		s           string
		p           []int
		q           []int
		expectedRes []int
	}{
		{description: "nominal case", s: "CAGCCTA", p: []int{2, 5, 0}, q: []int{4, 5, 6}, expectedRes: []int{2, 4, 1}},
	}

	for _, tc := range tt {
		t.Run(tc.description, func(t *testing.T) {
			res := GenomicRangeQuery(tc.s, tc.p, tc.q)
			assert.Equal(t, tc.expectedRes, res)
		})
	}
}
