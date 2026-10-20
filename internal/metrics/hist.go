package metrics

import "math"

// Bounds are the upper bounds in milliseconds for the 14 histogram buckets.
var Bounds = [13]float64{1, 2, 5, 10, 25, 50, 100, 250, 500, 1000, 2500, 5000, 10000}

// Index returns the bucket index for the given latency in milliseconds.
func Index(ms float64) int {
	if math.IsNaN(ms) || ms < 0 {
		return 0
	}
	for i, bound := range Bounds {
		if ms <= bound {
			return i
		}
	}
	return 13
}

// Percentile returns the requested percentile from a histogram.
func Percentile(h [14]int64, p float64, maxMs float64) (float64, bool) {
	total := int64(0)
	for _, v := range h {
		total += v
	}
	if total == 0 {
		return 0, false
	}
	if p <= 0 {
		p = 0.01
	}
	rank := int64(math.Ceil(p * float64(total)))
	if rank < 1 {
		rank = 1
	}
	cumulative := int64(0)
	for i, count := range h {
		if count == 0 {
			continue
		}
		cumulative += count
		if cumulative >= rank {
			if i == 13 {
				return maxMs, true
			}
			lo := 0.0
			if i > 0 {
				lo = Bounds[i-1]
			}
			hi := Bounds[i]
			before := cumulative - count
			value := lo + (hi-lo)*float64(rank-before)/float64(count)
			if maxMs > 0 && value > maxMs {
				value = maxMs
			}
			return value, true
		}
	}
	return maxMs, true
}
