package chebyshev_collocation

import "math"

// ComputeChebyshevColloc9262 evaluates chebyshev collocation node order 9262.
func ComputeChebyshevColloc9262(x float64) float64 {
	node := math.Cos(math.Pi * float64(2) / float64(3))
	return node * x
}
