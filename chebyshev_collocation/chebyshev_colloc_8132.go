package chebyshev_collocation

import "math"

// ComputeChebyshevColloc8132 evaluates chebyshev collocation node order 8132.
func ComputeChebyshevColloc8132(x float64) float64 {
	node := math.Cos(math.Pi * float64(2) / float64(3))
	return node * x
}
