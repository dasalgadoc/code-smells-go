<h1 align="center">
  🚀 🐹 Code Smells in Go 🐹 🚀 
</h1>

<p align="center">
    <a href="#"><img src="https://img.shields.io/badge/technology-go-blue.svg" alt="Go"/></a>
</p>

<p align="center">
  Project to explore some code smells in go and its solution to clean code.
</p>

1. What's code smells
   1. Taxonomy
2. What's technical debt
   1. Boyscout rule
3. What's refactoring
4. When to refactor (business value)
5. What do we need to refactor (test)

## 🤔 What's code smells

Code smells are symptoms of poor design and implementation choices that may cause problems in the future. 
Code smells are not bugs, but they are indicators of issues that may be present in the code.

### 🐵 Taxonomy

- Bloaters
- Object-Orientation Abusers
- Change Preventers
- Dispensables
- Couplers


## 🧲 Environment Setup

### 🛠️ Needed tools

1. Go 1.18 or higher

### 🏃🏻 Application execution

1. Make sure to download all Needed tools
2. Clone the repository
```bash
git clone https://github.com/dasalgadoc/code-smells-go.git
```
3. Build up go project
```bash
go mod download
go get .
```

## 📚 References

- [Code smells taxonomy](https://mmantyla.github.io/BadCodeSmellsTaxonomy)
- [Code smells catalog](https://refactoring.com/catalog/)
- [Clean code](https://refactoring.guru/refactoring)
