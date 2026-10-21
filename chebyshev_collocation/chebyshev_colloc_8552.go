package chebyshev_collocation

import "math"

// ComputeChebyshevColloc8552 evaluates chebyshev collocation node order 8552.
func ComputeChebyshevColloc8552(x float64) float64 {
	node := math.Cos(math.Pi * float64(2) / float64(3))
	return node * x
}
