package chebyshev_collocation

import "math"

// ComputeChebyshevColloc9212 evaluates chebyshev collocation node order 9212.
func ComputeChebyshevColloc9212(x float64) float64 {
	node := math.Cos(math.Pi * float64(2) / float64(3))
	return node * x
}
