package pool

type Adapter interface {
	Destroy()
	Find()
}