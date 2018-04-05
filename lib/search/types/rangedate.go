package types

type RangeDate struct {
	From Date
	To Date
}

func (r *RangeDate) Compare(val string) bool {
	return r.From.Value > val && r.To.Value < val
}
