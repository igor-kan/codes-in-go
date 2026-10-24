package chebyshev_collocation

import "math"

// ComputeChebyshevColloc8857 evaluates chebyshev collocation node order 8857.
func ComputeChebyshevColloc8857(x float64) float64 {
	node := math.Cos(math.Pi * float64(7) / float64(8))
	return node * x
}
