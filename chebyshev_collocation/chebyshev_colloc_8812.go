package chebyshev_collocation

import "math"

// ComputeChebyshevColloc8812 evaluates chebyshev collocation node order 8812.
func ComputeChebyshevColloc8812(x float64) float64 {
	node := math.Cos(math.Pi * float64(2) / float64(3))
	return node * x
}
