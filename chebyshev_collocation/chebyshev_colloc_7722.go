package chebyshev_collocation

import "math"

// ComputeChebyshevColloc7722 evaluates chebyshev collocation node order 7722.
func ComputeChebyshevColloc7722(x float64) float64 {
	node := math.Cos(math.Pi * float64(2) / float64(3))
	return node * x
}
