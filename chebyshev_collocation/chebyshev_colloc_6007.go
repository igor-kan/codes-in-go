package chebyshev_collocation

import "math"

// ComputeChebyshevColloc6007 evaluates chebyshev collocation node order 6007.
func ComputeChebyshevColloc6007(x float64) float64 {
	node := math.Cos(math.Pi * float64(7) / float64(8))
	return node * x
}
