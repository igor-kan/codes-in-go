package mod_circuit_halfopen_d58

type CircuitState int

const (
	StateClosed CircuitState = iota
	StateOpen
	StateHalfOpen
)

type Breaker struct {
	state    CircuitState
	fails    int
	limit    int
	cooldown int64
	openedAt int64
}

func NewBreaker(limit int, cooldown int64) *Breaker {
	return &Breaker{state: StateClosed, limit: limit, cooldown: cooldown}
}

func (cb *Breaker) Allow(nowMs int64) bool {
	if cb.state == StateOpen {
		if nowMs-cb.openedAt >= cb.cooldown {
			cb.state = StateHalfOpen
			return true
		}
		return false
	}
	return true
}

func (cb *Breaker) OnFail(nowMs int64) {
	cb.fails++
	if cb.state == StateHalfOpen || cb.fails >= cb.limit {
		cb.state = StateOpen
		cb.openedAt = nowMs
	}
}

func (cb *Breaker) OnOk() {
	cb.fails = 0
	cb.state = StateClosed
}
