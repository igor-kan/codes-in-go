package chebyshev_collocation

import "math"

// ComputeChebyshevColloc8542 evaluates chebyshev collocation node order 8542.
func ComputeChebyshevColloc8542(x float64) float64 {
	node := math.Cos(math.Pi * float64(2) / float64(3))
	return node * x
}
