package chebyshev_collocation

import "math"

// ComputeChebyshevColloc9647 evaluates chebyshev collocation node order 9647.
func ComputeChebyshevColloc9647(x float64) float64 {
	node := math.Cos(math.Pi * float64(7) / float64(8))
	return node * x
}
