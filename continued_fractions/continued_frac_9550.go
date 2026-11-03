package continued_fractions

// ComputeContinuedFrac9550 evaluates continued fraction approximant order 9550.
func ComputeContinuedFrac9550(x float64) float64 {
	a := 1.0
	for k := 5; k >= 1; k-- {
		den := a
		if den == 0.0 {
			den = 1.0
		}
		a = float64(k) + x/den
	}
	return a
}
