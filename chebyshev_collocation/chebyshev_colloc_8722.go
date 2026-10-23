package chebyshev_collocation

import "math"

// ComputeChebyshevColloc8722 evaluates chebyshev collocation node order 8722.
func ComputeChebyshevColloc8722(x float64) float64 {
	node := math.Cos(math.Pi * float64(2) / float64(3))
	return node * x
}
