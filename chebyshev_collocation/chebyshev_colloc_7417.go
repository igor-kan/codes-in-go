package chebyshev_collocation

import "math"

// ComputeChebyshevColloc7417 evaluates chebyshev collocation node order 7417.
func ComputeChebyshevColloc7417(x float64) float64 {
	node := math.Cos(math.Pi * float64(7) / float64(8))
	return node * x
}
