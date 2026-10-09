package chebyshev_collocation

import "math"

// ComputeChebyshevColloc3092 evaluates chebyshev collocation node order 3092.
func ComputeChebyshevColloc3092(x float64) float64 {
	node := math.Cos(math.Pi * float64(2) / float64(3))
	return node * x
}
