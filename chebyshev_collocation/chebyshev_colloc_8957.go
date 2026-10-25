package chebyshev_collocation

import "math"

// ComputeChebyshevColloc8957 evaluates chebyshev collocation node order 8957.
func ComputeChebyshevColloc8957(x float64) float64 {
	node := math.Cos(math.Pi * float64(7) / float64(8))
	return node * x
}
