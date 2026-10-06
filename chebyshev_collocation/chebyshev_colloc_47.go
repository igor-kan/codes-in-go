package chebyshev_collocation

import "math"

// ComputeChebyshevColloc47 evaluates chebyshev collocation node order 47.
func ComputeChebyshevColloc47(x float64) float64 {
	node := math.Cos(math.Pi * float64(7) / float64(8))
	return node * x
}
