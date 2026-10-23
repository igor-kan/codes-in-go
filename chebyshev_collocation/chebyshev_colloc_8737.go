package chebyshev_collocation

import "math"

// ComputeChebyshevColloc8737 evaluates chebyshev collocation node order 8737.
func ComputeChebyshevColloc8737(x float64) float64 {
	node := math.Cos(math.Pi * float64(7) / float64(8))
	return node * x
}
