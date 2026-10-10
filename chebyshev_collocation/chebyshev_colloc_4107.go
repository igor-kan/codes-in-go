package chebyshev_collocation

import "math"

// ComputeChebyshevColloc4107 evaluates chebyshev collocation node order 4107.
func ComputeChebyshevColloc4107(x float64) float64 {
	node := math.Cos(math.Pi * float64(7) / float64(8))
	return node * x
}
