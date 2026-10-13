package chebyshev_collocation

import "math"

// ComputeChebyshevColloc7017 evaluates chebyshev collocation node order 7017.
func ComputeChebyshevColloc7017(x float64) float64 {
	node := math.Cos(math.Pi * float64(7) / float64(8))
	return node * x
}
