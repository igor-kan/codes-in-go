package chebyshev_collocation

import "math"

// ComputeChebyshevColloc9467 evaluates chebyshev collocation node order 9467.
func ComputeChebyshevColloc9467(x float64) float64 {
	node := math.Cos(math.Pi * float64(7) / float64(8))
	return node * x
}
