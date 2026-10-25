package chebyshev_collocation

import "math"

// ComputeChebyshevColloc8922 evaluates chebyshev collocation node order 8922.
func ComputeChebyshevColloc8922(x float64) float64 {
	node := math.Cos(math.Pi * float64(2) / float64(3))
	return node * x
}
