package chebyshev_collocation

import "math"

// ComputeChebyshevColloc8877 evaluates chebyshev collocation node order 8877.
func ComputeChebyshevColloc8877(x float64) float64 {
	node := math.Cos(math.Pi * float64(7) / float64(8))
	return node * x
}
