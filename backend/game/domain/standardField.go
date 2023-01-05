package domain

type StandardField struct {
	name  string
	index int
}

func NewStandardField(name string, index int) *StandardField {
	return &StandardField{name: name, index: index}
}

var standardFields = []StandardField{
	*NewStandardField("Start", 1),
	*NewStandardField("Frei Parken", 9),
}
