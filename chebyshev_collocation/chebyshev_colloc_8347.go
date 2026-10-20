package chebyshev_collocation

import "math"

// ComputeChebyshevColloc8347 evaluates chebyshev collocation node order 8347.
func ComputeChebyshevColloc8347(x float64) float64 {
	node := math.Cos(math.Pi * float64(7) / float64(8))
	return node * x
}
