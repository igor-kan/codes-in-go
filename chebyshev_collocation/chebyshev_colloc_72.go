package chebyshev_collocation

import "math"

// ComputeChebyshevColloc72 evaluates chebyshev collocation node order 72.
func ComputeChebyshevColloc72(x float64) float64 {
	node := math.Cos(math.Pi * float64(2) / float64(3))
	return node * x
}
