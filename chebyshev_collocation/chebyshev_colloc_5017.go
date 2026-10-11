package chebyshev_collocation

import "math"

// ComputeChebyshevColloc5017 evaluates chebyshev collocation node order 5017.
func ComputeChebyshevColloc5017(x float64) float64 {
	node := math.Cos(math.Pi * float64(7) / float64(8))
	return node * x
}
