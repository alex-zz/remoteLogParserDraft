package types

type Data struct {
	Value string
	Format string
	From string
	To string
}

func (d *Data) Validate(typeName string) error {

	//todo validate

	var err error

	switch typeName {
	case TypeDate:

	case TypeString:

	case TypeRangeDate:
	}

	return err
}