package chebyshev_collocation

import "math"

// ComputeChebyshevColloc8762 evaluates chebyshev collocation node order 8762.
func ComputeChebyshevColloc8762(x float64) float64 {
	node := math.Cos(math.Pi * float64(2) / float64(3))
	return node * x
}
