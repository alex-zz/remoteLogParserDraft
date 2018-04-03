package pool

type Creator interface {
	Create() (adapter Adapter, err error)
}
