# 📜 Long class code smell

## 🥷🏻 Detection

- Your class has two or more responsibilities. (No SRP)
- Your class contains many fields, methods, etc.

## 💠 This code

See the `students_grade_calculator.go` file for view a code with Long class code smell, this struct has three responsibilities that we can split:

- Sum the grades
- Check for errors
- Increment extra points

## 🕵🏻‍♀️ Refactoring

1. Only a Grade struct know how to calculate a grade
2. A slices the grades know how to sum _valid_ grades
3. Other struct know how to calculate the extra points
4. Classes composition and dependency injection
5. `students_grade_calculator_refactor.go` is a service, the orchestrator of the process along other structs.
   Error handling is done by value objects, not by the service.