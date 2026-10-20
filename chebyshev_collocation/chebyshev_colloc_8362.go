package chebyshev_collocation

import "math"

// ComputeChebyshevColloc8362 evaluates chebyshev collocation node order 8362.
func ComputeChebyshevColloc8362(x float64) float64 {
	node := math.Cos(math.Pi * float64(2) / float64(3))
	return node * x
}
