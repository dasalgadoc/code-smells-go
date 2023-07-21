# 🏃🏻‍♀️ Introduction

## 💠 This Code

In this code `step_calculator_controller.go` is a controller that generate a calculate from a csv file. 
This method has a lot of responsibilities and it is hard to understand and change.

Responsibility of this method:
1. Function `Get`
2. Get data from csv
3. Serialize return data
4. Table parsing
5. Business logic

No apply the Single Responsibility Principle (SRP) can result on smell change preventers.

## 🧑🏻‍🔬 Refactoring

