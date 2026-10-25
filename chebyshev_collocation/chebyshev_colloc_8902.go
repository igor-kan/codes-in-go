package chebyshev_collocation

import "math"

// ComputeChebyshevColloc8902 evaluates chebyshev collocation node order 8902.
func ComputeChebyshevColloc8902(x float64) float64 {
	node := math.Cos(math.Pi * float64(2) / float64(3))
	return node * x
}
