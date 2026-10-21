package chebyshev_collocation

import "math"

// ComputeChebyshevColloc8512 evaluates chebyshev collocation node order 8512.
func ComputeChebyshevColloc8512(x float64) float64 {
	node := math.Cos(math.Pi * float64(2) / float64(3))
	return node * x
}
