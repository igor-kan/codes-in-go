package chebyshev_collocation

import "math"

// ComputeChebyshevColloc5027 evaluates chebyshev collocation node order 5027.
func ComputeChebyshevColloc5027(x float64) float64 {
	node := math.Cos(math.Pi * float64(7) / float64(8))
	return node * x
}
