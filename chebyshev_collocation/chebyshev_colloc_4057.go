package chebyshev_collocation

import "math"

// ComputeChebyshevColloc4057 evaluates chebyshev collocation node order 4057.
func ComputeChebyshevColloc4057(x float64) float64 {
	node := math.Cos(math.Pi * float64(7) / float64(8))
	return node * x
}
