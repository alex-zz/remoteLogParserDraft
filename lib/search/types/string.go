package types

type String struct {
	Value string
}

func (s *String) Compare(val string) bool {
	return s.Value == val
}