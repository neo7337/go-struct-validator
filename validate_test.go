package validator_test

import (
	"testing"
	"validator"
)

/**
mandatory constraint validation
*/
type ReqMsg struct {
	Name string `json:"name" constraints:"min-length=5"`
	Age  int    `json:"age" constraints:"min=10"`
}

func TestRequiredConstraintFail(t *testing.T) {
	msg := ReqMsg{
		Name: "Test",
		Age:  11,
	}
	sv := validator.NewStructValidator()
	err := sv.Validate(msg)
	got := err.Error()
	want := "mandatory field required not present"
	if got != want {
		t.Errorf("Expected: %s, got: %s", got, want)
	}
}

type ReqMsg2 struct {
	Name string `json:"name" constraints:"required=true,nillable=true"`
	Age  int    `json:"age" constraints:"required=true,nillable=true"`
}

func TestSuccessValidation(t *testing.T) {
	msg := ReqMsg2{
		Name: "Test",
		Age:  11,
	}
	sv := validator.NewStructValidator()
	if err := sv.Validate(msg); err != nil {
		t.Errorf("Error in validation: %s", err)
	}
}

/**
min, max validation
*/

type NumStruct struct {
	MinC int `json:"minC" constraints:"required=true,nillable=true,min=10"`
	MaxC int `json:"maxC" constraints:"required=true,nillable=true,max=49"`
}

func TestRequiredConstraintSuccess(t *testing.T) {
	msg1 := NumStruct{
		MinC: 12,
		MaxC: 45,
	}
	sv1 := validator.NewStructValidator()
	if err := sv1.Validate(msg1); err != nil {
		t.Errorf("Error in validation: %s", err)
	}

	msg2 := NumStruct{
		MinC: 7,
		MaxC: 45,
	}
	sv2 := validator.NewStructValidator()
	err2 := sv2.Validate(msg2)
	got2 := err2.Error()
	want2 := "min value validation failed"
	if got2 != want2 {
		t.Errorf("Expected: %s, got: %s", got2, want2)
	}

	msg3 := NumStruct{
		MinC: 12,
		MaxC: 55,
	}
	sv3 := validator.NewStructValidator()
	err3 := sv3.Validate(msg3)
	got3 := err3.Error()
	want3 := "max value validation failed"
	if got3 != want3 {
		t.Errorf("Expected: %s, got: %s", got3, want3)
	}
}

/**
min-length, max-length validation
*/

type StrStruct struct {
	Str1 string `json:"str1" constraints:"required=true,nillable=true,min-length=10"`
	Str2 string `json:"str2" constraints:"required=true,nillable=true,max-length=15"`
}

func TestStrConstraint(t *testing.T) {
	msg1 := StrStruct{
		Str1: "hello_world",
		Str2: "hello_world_go",
	}
	sv1 := validator.NewStructValidator()
	if err := sv1.Validate(msg1); err != nil {
		t.Errorf("Error in validation: %s", err)
	}

	msg2 := StrStruct{
		Str1: "hell_worl",
		Str2: "hello_world_go",
	}
	sv2 := validator.NewStructValidator()
	err2 := sv2.Validate(msg2)
	got2 := err2.Error()
	want2 := "min-length validation failed"
	if got2 != want2 {
		t.Errorf("Expected: %s, got: %s", got2, want2)
	}

	msg3 := StrStruct{
		Str1: "hello_world",
		Str2: "hello_world_from_go",
	}
	sv3 := validator.NewStructValidator()
	err3 := sv3.Validate(msg3)
	got3 := err3.Error()
	want3 := "max-length validation failed"
	if got3 != want3 {
		t.Errorf("Expected: %s, got: %s", got3, want3)
	}
}
