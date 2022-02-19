# go-struct-validator

Neo's Go Struct Validator is a simple validator built on the core of OAS Specification. The validator is heavily inspired by the OAS Specification approach leading to the creation of structs in a generic manner.

The validator covers the specifications, and its respective validations according to OAS.

---

- [Installation](#installation)
- [Benchmarking](#benchmarking)
- [Quick Start Guide](#quick-start-guide)
- [Features](#features)
  - [Validations Supported](#validations-supported)
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

### Quick Start Guide

It comes with a simple usage as explained below, just import the package, and you are good to go.

To add check for validations, add the `constraints` tag in the struct fields. 

```go

    package main
    
    import (
        "fmt"
        "github.com/neo7337/go-struct-validator/v1"
    )
    
    var sv = validator.NewStructValidator()
    
    type TestStruct struct {
        Name        string  `json:"name" constraints:"required=true|nillable=true|min-length=5"`
        Age         int     `json:"age" constraints:"required=true|nillable=true|min=21"`
        Description string  `json:"description" constraints:"required=true|nillable=true|max-length=50"`
        Cost        float64 `json:"cost" constraints:"required=true|nillable=true|exclusiveMin=200"`
        ItemCount   int     `json:"itemCount" constraints:"required=true|nillable=true|multipleOf=5"`
    }
    
    func main() {
        msg := TestStruct{
            Name:        "Test",
            Age:         25,
            Description: "this is bench testing",
            Cost:        299.9,
            ItemCount:   2000,
        }
        sv := NewStructValidator()
        if err := sv.Validate(msg); err != nil {
            fmt.Errorf(err)
        }
    }

```

### Features

#### Validations Supported
1. required
2. nillable
3. default
4. min
5. max
6. exclusiveMin
7. exclusiveMax
8. multipleOf
9. min-length
10. max-length
11. pattern
12. enum

