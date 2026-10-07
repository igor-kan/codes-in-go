package chebyshev_collocation

import "math"

// ComputeChebyshevColloc602 evaluates chebyshev collocation node order 602.
func ComputeChebyshevColloc602(x float64) float64 {
	node := math.Cos(math.Pi * float64(2) / float64(3))
	return node * x
}
