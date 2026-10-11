package chebyshev_collocation

import "math"

// ComputeChebyshevColloc5092 evaluates chebyshev collocation node order 5092.
func ComputeChebyshevColloc5092(x float64) float64 {
	node := math.Cos(math.Pi * float64(2) / float64(3))
	return node * x
}
