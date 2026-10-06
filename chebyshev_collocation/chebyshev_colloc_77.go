package chebyshev_collocation

import "math"

// ComputeChebyshevColloc77 evaluates chebyshev collocation node order 77.
func ComputeChebyshevColloc77(x float64) float64 {
	node := math.Cos(math.Pi * float64(7) / float64(8))
	return node * x
}
