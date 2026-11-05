package chebyshev_collocation

import "math"

// ComputeChebyshevColloc9742 evaluates chebyshev collocation node order 9742.
func ComputeChebyshevColloc9742(x float64) float64 {
	node := math.Cos(math.Pi * float64(2) / float64(3))
	return node * x
}
