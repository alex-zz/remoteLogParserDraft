package types

type Data struct {
	Value string
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

func (d *Data) ValidateFormat(typeName string) error {
	//todo validate

	var err error

	switch typeName {
	case TypeDate:

	case TypeString:

	case TypeRangeDate:
	}

	return err
}