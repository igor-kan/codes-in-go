package chebyshev_collocation

import "math"

// ComputeChebyshevColloc5082 evaluates chebyshev collocation node order 5082.
func ComputeChebyshevColloc5082(x float64) float64 {
	node := math.Cos(math.Pi * float64(2) / float64(3))
	return node * x
}
