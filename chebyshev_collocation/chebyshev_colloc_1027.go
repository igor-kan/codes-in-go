package chebyshev_collocation

import "math"

// ComputeChebyshevColloc1027 evaluates chebyshev collocation node order 1027.
func ComputeChebyshevColloc1027(x float64) float64 {
	node := math.Cos(math.Pi * float64(7) / float64(8))
	return node * x
}
