package mod_rw_fair_rwlock_queue

import "sync"

type FairRWMutex struct {
	mu            sync.Mutex
	cond          *sync.Cond
	activeReaders int
	writerActive  bool
	nextTicket    uint64
	servingTicket uint64
	waitingWrites int
}

func NewFairRWMutex() *FairRWMutex {
	m := &FairRWMutex{}
	m.cond = sync.NewCond(&m.mu)
	return m
}

func (m *FairRWMutex) RLock() {
	m.mu.Lock()
	myTicket := m.nextTicket
	m.nextTicket++
	for m.writerActive || m.servingTicket != myTicket {
		m.cond.Wait()
	}
	m.servingTicket++
	m.activeReaders++
	m.cond.Broadcast()
	m.mu.Unlock()
}

func (m *FairRWMutex) RUnlock() {
	m.mu.Lock()
	m.activeReaders--
	m.cond.Broadcast()
	m.mu.Unlock()
}

func (m *FairRWMutex) Lock() {
	m.mu.Lock()
	myTicket := m.nextTicket
	m.nextTicket++
	m.waitingWrites++
	for m.writerActive || m.servingTicket != myTicket {
		m.cond.Wait()
	}
	for m.activeReaders > 0 {
		m.cond.Wait()
	}
	m.waitingWrites--
	m.writerActive = true
	m.servingTicket++
	m.mu.Unlock()
}

func (m *FairRWMutex) Unlock() {
	m.mu.Lock()
	m.writerActive = false
	m.cond.Broadcast()
	m.mu.Unlock()
}
