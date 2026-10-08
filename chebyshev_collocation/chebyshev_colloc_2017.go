package chebyshev_collocation

import "math"

// ComputeChebyshevColloc2017 evaluates chebyshev collocation node order 2017.
func ComputeChebyshevColloc2017(x float64) float64 {
	node := math.Cos(math.Pi * float64(7) / float64(8))
	return node * x
}
