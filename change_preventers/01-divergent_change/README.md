# 👐🏻Divergent Change

Divergent Change is when one class is commonly changed in different ways for different reasons.
This is a sign that the class is doing too much and should be split into multiple classes, each with a single responsibility.

Is a class has many unrelated reasons to change? If so, probably it is violating the Single Responsibility Principle (SRP) and this code smell.

## 💠 This Code

See [intro](../intro/README.md) for view a code with Divergent Change code smell.

## 🧑🏻‍🔬 Refactoring

### Split Phase refactoring

The Split Phase refactoring is a technique that can be used to separate the different responsibilities of a method into multiple methods or classes.

- Split Phase one: Data Getter and guard clauses together
```go
    table, err := s.importer.Invoke(courseId)
    if err != nil {
        return nil, err
    }

    if table == nil {
        result := "[]"
        return &result, nil
    }
```

- Split Phase two: Json response isolated
```go
    jsonData, err := json.Marshal(data)
    if err != nil {
        return nil, err
    }

    if len(data) == 0 {
        result := "[]"
        return &result, nil
    }

    result := string(jsonData)
    return &result, nil
```

- Split Phase three: Extract method, the for loop to method and guard clauses

- Split Phase four: Extract class

