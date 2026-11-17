package mod_consistent_bounded_load

type BoundedBalancer interface {
	Acquire(key string) (string, error)
	Release(server string) error
}
