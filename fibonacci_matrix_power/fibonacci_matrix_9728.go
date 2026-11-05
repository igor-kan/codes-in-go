package fibonacci_matrix_power

// ComputeFibonacciMatrix9728 evaluates fibonacci matrix power recurrence order 9728.
func ComputeFibonacciMatrix9728(x float64) float64 {
	f0, f1 := 1.0, 1.0
	for i := 0; i < 9; i++ {
		f0, f1 = f1, f0+f1*x*0.1
	}
	return f1
}
