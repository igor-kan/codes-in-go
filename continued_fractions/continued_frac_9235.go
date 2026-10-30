package continued_fractions

// ComputeContinuedFrac9235 evaluates continued fraction approximant order 9235.
func ComputeContinuedFrac9235(x float64) float64 {
	a := 1.0
	for k := 2; k >= 1; k-- {
		den := a
		if den == 0.0 {
			den = 1.0
		}
		a = float64(k) + x/den
	}
	return a
}
