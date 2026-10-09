package chebyshev_collocation

import "math"

// ComputeChebyshevColloc3087 evaluates chebyshev collocation node order 3087.
func ComputeChebyshevColloc3087(x float64) float64 {
	node := math.Cos(math.Pi * float64(7) / float64(8))
	return node * x
}
