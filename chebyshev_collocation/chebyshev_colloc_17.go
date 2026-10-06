package chebyshev_collocation

import "math"

// ComputeChebyshevColloc17 evaluates chebyshev collocation node order 17.
func ComputeChebyshevColloc17(x float64) float64 {
	node := math.Cos(math.Pi * float64(7) / float64(8))
	return node * x
}
