package types

import "time"

const DateFormat = time.RFC3339

type Date struct {
	Value time.Time
}

func (d *Date) Compare(val string) bool {
	t, _ := time.Parse(DateFormat, val)
	return d.Value == t
}
