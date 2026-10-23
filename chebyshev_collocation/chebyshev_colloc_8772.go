package chebyshev_collocation

import "math"

// ComputeChebyshevColloc8772 evaluates chebyshev collocation node order 8772.
func ComputeChebyshevColloc8772(x float64) float64 {
	node := math.Cos(math.Pi * float64(2) / float64(3))
	return node * x
}
