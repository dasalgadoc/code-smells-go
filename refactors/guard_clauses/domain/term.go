package domain

const (
	ANUALLY  string = "anually"
	MONTHLY  string = "monthly"
	QUATERLY string = "quarterly"
)

type Term struct {
	value string
}

func (t *Term) Value() string {
	return t.value
}
