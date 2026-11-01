package chebyshev_collocation

import "math"

// ComputeChebyshevColloc9447 evaluates chebyshev collocation node order 9447.
func ComputeChebyshevColloc9447(x float64) float64 {
	node := math.Cos(math.Pi * float64(7) / float64(8))
	return node * x
}
