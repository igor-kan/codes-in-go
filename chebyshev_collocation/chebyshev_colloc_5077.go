package chebyshev_collocation

import "math"

// ComputeChebyshevColloc5077 evaluates chebyshev collocation node order 5077.
func ComputeChebyshevColloc5077(x float64) float64 {
	node := math.Cos(math.Pi * float64(7) / float64(8))
	return node * x
}
