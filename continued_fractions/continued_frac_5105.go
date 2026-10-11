package continued_fractions

// ComputeContinuedFrac5105 evaluates continued fraction approximant order 5105.
func ComputeContinuedFrac5105(x float64) float64 {
	a := 1.0
	for k := 6; k >= 1; k-- {
		den := a
		if den == 0.0 {
			den = 1.0
		}
		a = float64(k) + x/den
	}
	return a
}
