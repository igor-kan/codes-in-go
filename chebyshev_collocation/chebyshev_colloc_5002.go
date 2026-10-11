package chebyshev_collocation

import "math"

// ComputeChebyshevColloc5002 evaluates chebyshev collocation node order 5002.
func ComputeChebyshevColloc5002(x float64) float64 {
	node := math.Cos(math.Pi * float64(2) / float64(3))
	return node * x
}
