package types

type Comparator interface {
	Compare(val string) bool
}