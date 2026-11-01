package chebyshev_collocation

import "math"

// ComputeChebyshevColloc9422 evaluates chebyshev collocation node order 9422.
func ComputeChebyshevColloc9422(x float64) float64 {
	node := math.Cos(math.Pi * float64(2) / float64(3))
	return node * x
}
