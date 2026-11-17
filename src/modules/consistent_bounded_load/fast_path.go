package mod_consistent_bounded_load

import "math"

func fastMaxAllowed(totalReq, numServers int, factor float64) int {
	return int(math.Ceil(factor * float64(totalReq) / float64(numServers)))
}
