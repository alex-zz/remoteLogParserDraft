package types

type String struct {
	Name string
	Value string
}

func (s *String) Compare(val interface{}) bool {
	return s.Value == val
}