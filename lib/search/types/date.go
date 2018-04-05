package types

type Date struct {
	Value string
	Format string
}

func (d *Date) Compare(val string) bool {
	return d.Value == val
}
