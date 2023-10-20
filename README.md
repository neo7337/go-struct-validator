# go-struct-validator

![Build Status](https://github.com/neo7337/go-struct-validator/actions/workflows/ci.yml/badge.svg)

Neo's Go Struct Validator is a simple validator built on the core of OAS Specification. The validator is heavily
inspired by the OAS Specification approach leading to the creation of structs in a generic manner.

The validator covers the specifications, and its respective validations according to OAS.

---

- [Installation](#installation)
- [Benchmarking](#benchmarking)
- [Quick Start Guide](#quick-start-guide)
- [Features](#features)
    - [Validations Supported](#validations-supported)
- [Contributing](#contributing)

---

### Installation

```bash
go get github.com/neo7337/go-struct-validator
```

### Benchmarking

```bash
go test -run=Bench -bench=. -benchtime 5000000x
```

| S.No. |             Name             | Ops     | BenchResult |
|:------|:----------------------------:|---------|-------------|
| 1     |     BenchmarkValidator-8     | 5000000 | 849 ns/op   |
| 1     | BenchmarkValidatorParallel-8 | 5000000 | 268 ns/op   |

### Quick Start Guide

It comes with a simple usage as explained below, just import the package, and you are good to go.

To add check for validations, add the `constraints` tag in the struct fields or configure a one of your desire using

```go
package main

import (
	"fmt"
	"github.com/neo7337/go-struct-validator"
)

var sv = validator.NewStructValidator()

type TestStruct struct {
	Name        string  `json:"name" constraints:"min-length=5"`
	Age         int     `json:"age" constraints:"min=21"`
	Description string  `json:"description" constraints:"max-length=50"`
	Cost        float64 `json:"cost" constraints:"exclusiveMin=200"`
	ItemCount   int     `json:"itemCount" constraints:"multipleOf=5"`
}

func main() {
	msg := TestStruct{
		Name:        "Test",
		Age:         25,
		Description: "this is bench testing",
		Cost:        299.9,
		ItemCount:   2000,
	}

	if err := sv.Validate(msg); err != nil {
		fmt.Errorf(err)
	}
}
```

### Features

#### Validations Supported

| S.No. |     Name     | Data Type Supported | Status |
|:------|:------------:|---------------------|--------|
| 4     |     min      | numeric             | ✅      |
| 5     |     max      | numeric             | ✅      |
| 6     | exclusiveMin | numeric             | ✅      |
| 7     | exclusiveMax | numeric             | ✅      |
| 8     |  multipleOf  | numeric             | ✅      |
| 9     |  max-length  | string              | ✅      |
| 10    |  min-length  | string              | ✅      |
| 11    |   pattern    | string              | ✅      |
| 11    |   notnull    | string              | ✅      |
| 12    |     enum     | all                 | ✅      |

### Contributing
We welcome contributions to the project. If you find a bug or would like to request a new feature, please open an issue 
on [GitHub](https://github.com/neo7337/go-struct-validator/issues).