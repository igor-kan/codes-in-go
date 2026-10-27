package chebyshev_collocation

import "math"

// ComputeChebyshevColloc9117 evaluates chebyshev collocation node order 9117.
func ComputeChebyshevColloc9117(x float64) float64 {
	node := math.Cos(math.Pi * float64(7) / float64(8))
	return node * x
}
