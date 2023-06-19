# 🐍 Long method code smell

## 🥷🏻 Detection

- Your code has two or more responsibilities. (No SRP)
- Your code is too long to understand.
- Your code is too long to test.
- Your code has too many levels of indentation
- Your code has too many lines (more than 10-20), parameters, variables, etc.

## 💠 This Code

See the students_grade_calculator.go file for view a code with Long method code smell.

All starts with a simple feature, an average:
```go
func (s *studentGradeCalculator) calculateGrades(
    examsGrades []domain.Grades) domain.Grades {
	
    if !(len(examsGrades) == 0) {
        var gradesSum domain.Grades
        var gradesCount domain.Grades

        for _, grade := range examsGrades {
            gradesSum += grade
            gradesCount++
        }

        return gradesSum / gradesCount
    } else {
        return 0
    }
	
}
```
Consider the following aspects:
- If there is no grades, the average is 0.

Later, a new feature is added, the student have to attend to some minimum classes to pass the subject.
If we add this feature to the previous code, we have:

```go
func (s *studentGradeCalculator) calculateGradesMinimumClasses(
    examsGrades []domain.Grades,
    hasReachMinimumGrades bool) domain.Grades {
	
    if !(len(examsGrades) == 0) {
        var gradesSum domain.Grades
        var gradesCount domain.Grades

        for _, grade := range examsGrades {
            gradesSum += grade
            gradesCount++
        }

        if hasReachMinimumGrades {
            return gradesSum / gradesCount
        } else {
            return 0
        }

    } else {
        return 0
    }
}
```

Consider the following aspects:
- If there is no grades, the average is 0.
- If the student has not reached the minimum classes, the average is 0.

Here we have a problem, the code is too long to understand, and we have a lot of levels of indentation, but it can be worse.

Later, we have a new feature, weighted average for some exams:
```go
func (s *studentGradeCalculator) calculateGradesMinimumClassesAndWeightedAverage(
    examsGrades []domain.StudentGrade,
    hasReachMinimumGrades bool) domain.Grades {
	
    if !(len(examsGrades) == 0) {
        var gradesSum domain.Grades
        var gradesCount domain.Grades
        weightSum := 0

        for _, grade := range examsGrades {
            gradesSum += grade.Value * domain.Grades(grade.Weight) / 100
            weightSum += grade.Weight
            gradesCount++
        }

        // errors handling: oue return type is grades, so we can't return an error
        // magic numbers'
        // concept overload
        if weightSum == 100 {
            if hasReachMinimumGrades {
                return gradesSum / gradesCount
            } else {
                return 0
            }
        } else if weightSum > 100 {
            return -1 // error type over-weighed
        } else {
            return -2 // error type under-weighed
        }

    } else {
        return 0
    }
}
```

Consider the following aspects:
- If there is no grades, the average is 0.
- If the student has not reached the minimum classes, the average is 0.
- The sum of the weights must be 100, otherwise you have to raise an error.

But, we have a simply return type, how we manage errors? through magic numbers, and we have a concept overload, the function is doing too much.
And, it can be worse, we can add more features.

Some teachers can give a bonus to the students:

```go
func (s *studentGradeCalculator) calculateGradesMinimumClassesAndWeightedAverageWithExtraPoint(
    examsGrades []domain.StudentGrade,
    hasReachMinimumGrades bool,
    teacher string) domain.Grades {
	
    if !(len(examsGrades) == 0) {
        var gradesSum domain.Grades
        var gradesCount domain.Grades
        weightSum := 0
        gotExtraPoint := false

        if s.teacherExtraPoint != nil {
            currentYear := time.Now().Year()
            if extra, ok := s.teacherExtraPoint[teacher][currentYear]; ok {
                gotExtraPoint = extra
            }
        }

        for _, grade := range examsGrades {
            gradesSum += grade.Value * domain.Grades(grade.Weight) / 100
            weightSum += grade.Weight
            gradesCount++
        }

        // errors handling: oue return type is grades, so we can't return an error
        // magic numbers'
        // concept overload
        if weightSum == 100 {
            if hasReachMinimumGrades {
                if gotExtraPoint {
                    return gradesSum/gradesCount + 1
                }
                return gradesSum / gradesCount
            } else {
                return 0
            }
        } else if weightSum > 100 {
            return -1 // error type over-weighed
        } else {
            return -2 // error type under-weighed
        }

    } else {
        return 0
    }
}
```

Consider the following aspects:
- If there is no grades, the average is 0.
- If the student has not reached the minimum classes, the average is 0.
- The sum of the weights must be 100, otherwise you have to raise an error.
- If the teacher gives an extra point, the average is the average plus one.

So, our simply average function growth enough and someday we will death for technical debt.

## 🤷🏻‍♀️ Issues

## 🧑🏻‍🔬 Refactoring