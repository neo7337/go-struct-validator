package validator

import "testing"

type BenchTestStruct struct {
	Name        string  `json:"name" constraints:"required=true,nillable=true,min-length=5"`
	Age         int     `json:"age" constraints:"required=true,nillable=true,min=21"`
	Description string  `json:"description" constraints:"required=true,nillable=true,max-length=50"`
	Cost        float64 `json:"cost" constraints:"required=true,nillable=true,exclusiveMin=200"`
	ItemCount   int     `json:"itemCount" constraints:"required=true,nillable=true,multipleOf=5"`
}

func BenchmarkValidator(b *testing.B) {
	msg := BenchTestStruct{
		Name:        "BenchTest",
		Age:         25,
		Description: "this is bench testing",
		Cost:        299.9,
		ItemCount:   2000,
	}
	sv := NewStructValidator()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = sv.Validate(msg)
	}
}

func BenchmarkValidatorParallel(b *testing.B) {
	msg := BenchTestStruct{
		Name:        "BenchTest",
		Age:         25,
		Description: "this is bench testing",
		Cost:        299.9,
		ItemCount:   2000,
	}
	sv := NewStructValidator()

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_ = sv.Validate(msg)
		}
	})
}
