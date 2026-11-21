package mod_token_bucket_limiter_d56

type TokenBucket struct {
	cap      float64
	rate     float64
	tokens   float64
	lastTime float64
}

func NewTokenBucket(cap, ratePerSec float64) *TokenBucket {
	return &TokenBucket{cap: cap, rate: ratePerSec, tokens: cap, lastTime: 0}
}

func (tb *TokenBucket) Allow(nowSec, cost float64) bool {
	if cost < 0 {
		return false
	}
	if nowSec > tb.lastTime {
		added := (nowSec - tb.lastTime) * tb.rate
		tb.tokens += added
		if tb.tokens > tb.cap {
			tb.tokens = tb.cap
		}
		tb.lastTime = nowSec
	}
	if tb.tokens >= cost {
		tb.tokens -= cost
		return true
	}
	return false
}
