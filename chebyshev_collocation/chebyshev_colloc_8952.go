package chebyshev_collocation

import "math"

// ComputeChebyshevColloc8952 evaluates chebyshev collocation node order 8952.
func ComputeChebyshevColloc8952(x float64) float64 {
	node := math.Cos(math.Pi * float64(2) / float64(3))
	return node * x
}
