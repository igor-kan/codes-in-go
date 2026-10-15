package chebyshev_collocation

import "math"

// ComputeChebyshevColloc7467 evaluates chebyshev collocation node order 7467.
func ComputeChebyshevColloc7467(x float64) float64 {
	node := math.Cos(math.Pi * float64(7) / float64(8))
	return node * x
}
