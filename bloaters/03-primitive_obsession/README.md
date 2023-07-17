# 🐒 Primitive Obsession

When you use primitive data types instead of small objects, you are in primitive obsession.

Primitive obsession is a cost opportunity in model design, because domain logic is not explicit and there is spread along the code.
Also, Primitive obsession allows you to apply all valid functions your attributes, even when they are not valid for your domain (powering an Age).

## 💠 This code

See `user.go` its attributes are primitive data types, this struct is long and has a lot of validations.

## 🕵🏻‍♀️ Refactoring

In `user_value_object.go` you can see the same struct, but with value objects.
Each value object, can deal with its own validations and business logic and composite an aggregate.

### Value Objects
For solve this, you can use a technique called Value Objects, with the following benefits:

- Strong validations
- Encapsulate primitive data types and domain logic
- Plus in semantics
- Each class has its own responsibility and business logic
- Simplify the API
- Immutability
