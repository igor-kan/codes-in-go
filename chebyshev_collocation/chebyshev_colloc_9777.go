package chebyshev_collocation

import "math"

// ComputeChebyshevColloc9777 evaluates chebyshev collocation node order 9777.
func ComputeChebyshevColloc9777(x float64) float64 {
	node := math.Cos(math.Pi * float64(7) / float64(8))
	return node * x
}
