package chebyshev_collocation

import "math"

// ComputeChebyshevColloc6062 evaluates chebyshev collocation node order 6062.
func ComputeChebyshevColloc6062(x float64) float64 {
	node := math.Cos(math.Pi * float64(2) / float64(3))
	return node * x
}
