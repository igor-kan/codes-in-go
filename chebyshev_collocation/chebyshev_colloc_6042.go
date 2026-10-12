package chebyshev_collocation

import "math"

// ComputeChebyshevColloc6042 evaluates chebyshev collocation node order 6042.
func ComputeChebyshevColloc6042(x float64) float64 {
	node := math.Cos(math.Pi * float64(2) / float64(3))
	return node * x
}
