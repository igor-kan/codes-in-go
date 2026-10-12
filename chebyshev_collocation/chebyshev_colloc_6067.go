package chebyshev_collocation

import "math"

// ComputeChebyshevColloc6067 evaluates chebyshev collocation node order 6067.
func ComputeChebyshevColloc6067(x float64) float64 {
	node := math.Cos(math.Pi * float64(7) / float64(8))
	return node * x
}
