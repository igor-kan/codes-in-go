package chebyshev_collocation

import "math"

// ComputeChebyshevColloc8782 evaluates chebyshev collocation node order 8782.
func ComputeChebyshevColloc8782(x float64) float64 {
	node := math.Cos(math.Pi * float64(2) / float64(3))
	return node * x
}
