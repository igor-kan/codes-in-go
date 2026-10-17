package chebyshev_collocation

import "math"

// ComputeChebyshevColloc7717 evaluates chebyshev collocation node order 7717.
func ComputeChebyshevColloc7717(x float64) float64 {
	node := math.Cos(math.Pi * float64(7) / float64(8))
	return node * x
}
