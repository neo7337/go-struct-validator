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

| S.No. |         Name         | Ops     | BenchResult |
|:------|:--------------------:|---------|-------------|
| 1     | BenchmarkValidator-8 | 105066  | 9718 ns/op  |
| 2     | BenchmarkValidator-8 | 150873  | 7100 ns/op  |
| 3     | BenchmarkValidator-8 | 170964  | 6709 ns/op  |
| 4     | BenchmarkValidator-8 | 5000000 | 6546 ns/op  |
| 5     | BenchmarkValidator-8 | 5000000 | 1315 ns/op  |

### WIP
1. Base Struct Validator
2. Enums Validation 
3. Nested Struct Validation
4. Benchmarking
5. Cache Implementation
