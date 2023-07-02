package domain

type TeachersExtraPoint struct {
	value []TeacherExtraPoint
}

func (t *TeachersExtraPoint) GetExtraPoint(teacher string, grade int) bool {
	for _, v := range t.value {
		if v.Name == teacher {
			return v.GivenAnExtraPoint[grade]
		}
	}
	return false
}

func NewTeachersExtraPoint(value []TeacherExtraPoint) TeachersExtraPoint {
	return TeachersExtraPoint{
		value: value,
	}
}
