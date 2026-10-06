package chebyshev_collocation

import "math"

// ComputeChebyshevColloc52 evaluates chebyshev collocation node order 52.
func ComputeChebyshevColloc52(x float64) float64 {
	node := math.Cos(math.Pi * float64(2) / float64(3))
	return node * x
}
