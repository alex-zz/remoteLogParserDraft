package types

const (
	TypeDate = "{{type.date}}"
	TypeString = "{{type.string}}"
	TypeRangeDate = "{{type.range(date)}}"
)

func GetAvailableTypes() []string {
	return []string{TypeDate, TypeString, TypeRangeDate}
}

func BuildType(typeName string, data Data) (Comparator, error) {
	var c Comparator
	var err error

	if err = data.Validate(typeName) ; err != nil {
		return c, err
	}

	/**switch typeName {
	case TypeDate:
		c = &Date{
			Value: data.Value,
			Format: data.Format,
		}
	case TypeString:
		c = &String{
			Value: data.Value,
		}
	case TypeRangeDate:
		from := Date{
			Value: data.From,
			Format: data.Format,
		}

		to := Date{
			Value: data.To,
			Format: data.Format,
		}

		c = &RangeDate{
			From: from,
			To: to,
		}
	default:
		err = errors.New("unsupported type")
	}*/

	return c, err
}