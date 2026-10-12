package chebyshev_collocation

import "math"

// ComputeChebyshevColloc6012 evaluates chebyshev collocation node order 6012.
func ComputeChebyshevColloc6012(x float64) float64 {
	node := math.Cos(math.Pi * float64(2) / float64(3))
	return node * x
}
