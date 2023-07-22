package domain

type Table struct {
	Headers []string
	Rows    [][]string
}

func (t *Table) IsEmpty() bool {
	return len(t.Rows) == 0
}
