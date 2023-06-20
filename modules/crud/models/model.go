package models

type IdentifiedModel struct {
	ID   string
	Data any
}

type Models struct {
	TotalCount int64
	Items      []IdentifiedModel
}
