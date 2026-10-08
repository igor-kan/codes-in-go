package chebyshev_collocation

import "math"

// ComputeChebyshevColloc1042 evaluates chebyshev collocation node order 1042.
func ComputeChebyshevColloc1042(x float64) float64 {
	node := math.Cos(math.Pi * float64(2) / float64(3))
	return node * x
}
