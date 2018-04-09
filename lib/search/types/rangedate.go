package types

import "time"

type RangeDate struct {
	From Date
	To Date
}

func (r *RangeDate) Compare(val string) bool {
	t, _ := time.Parse(DateFormat, val)
	return r.From.Value.After(t) && r.To.Value.Before(t)
}
