# ♖ Strategy

Strategy pattern is a behavioral design pattern that lets you define a family of algorithms, put each of them into a separate class, and make their objects interchangeable.

## 💠 This Code

See the `date_generator.go` to see a Code smell.

There is a problem and many ways to solved it. The `date_generator.go` is a solution that is not flexible and not scalable, it fails with SRP and OCP.
```go
func (l *LimitDates) GenerateDates() {
    nov := CurrentDate()
    if l.Type == DayLimitType {
        l.Gte = time.Date(nov.Year(), nov.Month(), nov.Day(), 0, 0, 0, 0, time.UTC)
        l.Lte = time.Date(nov.Year(), nov.Month(), nov.Day(), 23, 59, 59, 999, time.UTC)
    }
    if l.Type == MonthLimitType {
        l.Gte = time.Date(nov.Year(), nov.Month(), 1, 0, 0, 0, 0, time.UTC)
        l.Lte = time.Date(nov.Year(), nov.Month()+1, 0, 23, 59, 59, 999, time.UTC)
    }
    if l.Type == YearLimitType {
        l.Gte = time.Date(nov.Year(), 1, 1, 0, 0, 0, 0, time.UTC)
        l.Lte = time.Date(nov.Year(), 12, 31, 23, 59, 59, 999, time.UTC)
    }
    // ... more strategies
}
```


## 🎯 Refactoring

In `date_generator_refactor.go` we see a better approach, using the Strategy Pattern.
With this patter we can add more strategies without changing the `LimitDates` struct, so we apply SRP and OCP.

Just we need to inject a strategy generator.

```go
func (l *LimitDatesRefactor) GenerateDates() {
    now := CurrentDate()
    l.Gte, l.Lte = l.Generator.Generate(now)
}
```

Generate is a interface that has a method `Generate` that returns two dates and concrete structs implement concrete strategies.

```go
type Generator interface {
    Generate(date time.Time) (time.Time, time.Time)
}
```