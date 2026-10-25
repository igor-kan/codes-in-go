package chebyshev_collocation

import "math"

// ComputeChebyshevColloc8882 evaluates chebyshev collocation node order 8882.
func ComputeChebyshevColloc8882(x float64) float64 {
	node := math.Cos(math.Pi * float64(2) / float64(3))
	return node * x
}
