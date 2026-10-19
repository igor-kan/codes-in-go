package chebyshev_collocation

import "math"

// ComputeChebyshevColloc8127 evaluates chebyshev collocation node order 8127.
func ComputeChebyshevColloc8127(x float64) float64 {
	node := math.Cos(math.Pi * float64(7) / float64(8))
	return node * x
}
