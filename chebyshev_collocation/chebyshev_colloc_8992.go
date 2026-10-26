package chebyshev_collocation

import "math"

// ComputeChebyshevColloc8992 evaluates chebyshev collocation node order 8992.
func ComputeChebyshevColloc8992(x float64) float64 {
	node := math.Cos(math.Pi * float64(2) / float64(3))
	return node * x
}
