package chebyshev_collocation

import "math"

// ComputeChebyshevColloc9347 evaluates chebyshev collocation node order 9347.
func ComputeChebyshevColloc9347(x float64) float64 {
	node := math.Cos(math.Pi * float64(7) / float64(8))
	return node * x
}
