package models

type IdentifiedRecord struct {
	ID   string
	Data any
}

type Records struct {
	TotalCount int64
	Items      []IdentifiedRecord
}
