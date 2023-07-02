package domain

type TeachersExtraPoint struct {
	value []TeacherExtraPoint
}

func (t *TeachersExtraPoint) TeacherGivesAnExtraPoint(teacher string, year int) bool {
	for _, v := range t.value {
		if v.Name == teacher {
			return v.GivenAnExtraPoint[year]
		}
	}
	return false
}

func (t *TeachersExtraPoint) GetExtraPoint(teacher string, year int) NumericGrade {
	if t.TeacherGivesAnExtraPoint(teacher, year) {
		return 1
	}
	return 0
}

func NewTeachersExtraPoint(value []TeacherExtraPoint) TeachersExtraPoint {
	return TeachersExtraPoint{
		value: value,
	}
}
