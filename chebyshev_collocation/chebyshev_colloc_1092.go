package chebyshev_collocation

import "math"

// ComputeChebyshevColloc1092 evaluates chebyshev collocation node order 1092.
func ComputeChebyshevColloc1092(x float64) float64 {
	node := math.Cos(math.Pi * float64(2) / float64(3))
	return node * x
}
