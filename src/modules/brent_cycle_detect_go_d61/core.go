package mod_brent_cycle_detect_go_d61

func BrentCycle(next func(int) int, start int) CycleInfo {
	power := 1
	lambda := 1
	tortoise := start
	hare := next(start)
	for tortoise != hare {
		if power == lambda {
			tortoise = hare
			power *= 2
			lambda = 0
		}
		hare = next(hare)
		lambda++
	}
	tortoise = start
	hare = start
	for i := 0; i < lambda; i++ {
		hare = next(hare)
	}
	mu := 0
	for tortoise != hare {
		tortoise = next(tortoise)
		hare = next(hare)
		mu++
	}
	return CycleInfo{Mu: mu, Lambda: lambda, HasCycle: true}
}

type BrentDetector struct{}

func (b BrentDetector) Detect(next func(int) int, start int) CycleInfo {
	return BrentCycle(next, start)
}

var _ CycleDetector = BrentDetector{}
