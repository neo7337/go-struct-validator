# go-struct-validator
Golang Struct Validator

The struct validator is a generic solution to manage the structs validation in golang.

---

- [Installation](#installation)
- [Benchmarking](#benchmarking)
- [WIP](#wip)

---

### Installation

```bash
go get ____
```

### Benchmarking

```bash
go test -run=Bench -bench=. -benchtime 5000000x
```

| S.No. |             Name             | Ops     | BenchResult |
|:------|:----------------------------:|---------|-------------|
| 1     |     BenchmarkValidator-8     | 5000000 | 1288 ns/op  |
| 1     | BenchmarkValidatorParallel-8 | 5000000 | 385 ns/op   |

### WIP
1. Base Struct Validator
2. Enums Validation 
3. Nested Struct Validation
4. Benchmarking
5. Cache Implementation
