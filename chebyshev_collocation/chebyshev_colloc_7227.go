package chebyshev_collocation

import "math"

// ComputeChebyshevColloc7227 evaluates chebyshev collocation node order 7227.
func ComputeChebyshevColloc7227(x float64) float64 {
	node := math.Cos(math.Pi * float64(7) / float64(8))
	return node * x
}
