package domain

type StandardField struct {
	name  string
	index int
}

func newStandardField(name string, index int) *StandardField {
	return &StandardField{name: name, index: index}
}

var standardFields = []StandardField{
	*newStandardField("Start", 1),
	*newStandardField("Frei Parken", 9),
}
