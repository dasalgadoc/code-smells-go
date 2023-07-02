# 📜 Long class code smell

## 🥷🏻 Detection

- Your class has two or more responsibilities. (No SRP)
- Your class contains many fields, methods, etc.

## 💠 This code

See the `students_grade_calculator.go` file for view a code with Long class code smell, this struct has three responsibilities that we can split:

- Sum the grades
- Check for errors
- Increment extra points

We can split responsibilities to models

```go
```

## 🕵🏻‍♀️ Refactoring