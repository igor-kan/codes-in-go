package chebyshev_collocation

import "math"

// ComputeChebyshevColloc8997 evaluates chebyshev collocation node order 8997.
func ComputeChebyshevColloc8997(x float64) float64 {
	node := math.Cos(math.Pi * float64(7) / float64(8))
	return node * x
}
