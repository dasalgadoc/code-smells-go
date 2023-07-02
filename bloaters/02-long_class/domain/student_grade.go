package domain

type NumericGrade float32

type Grade struct {
	Value  NumericGrade
	Weight int
}

func (s *Grade) CalculateGrade() NumericGrade {
	return s.Value * NumericGrade(s.Weight) / 100
}
