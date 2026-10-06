package chebyshev_collocation

import "math"

// ComputeChebyshevColloc22 evaluates chebyshev collocation node order 22.
func ComputeChebyshevColloc22(x float64) float64 {
	node := math.Cos(math.Pi * float64(2) / float64(3))
	return node * x
}
