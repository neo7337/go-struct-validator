package validator_test

import (
	"testing"
	"validator"
)

/**
mandatory constraint validations test
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
	want := "mandatory fields not present"
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
min, max validations test
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
exclusive min/max validation test
*/

type ExlStruct struct {
	MinC int `json:"minC" constraints:"required=true,nillable=true,exclusiveMin=10"`
	MaxC int `json:"maxC" constraints:"required=true,nillable=true,exclusiveMax=50"`
}

func TestExlConstraint(t *testing.T) {
	msg1 := ExlStruct{
		MinC: 10,
		MaxC: 50,
	}
	sv1 := validator.NewStructValidator()
	if err := sv1.Validate(msg1); err != nil {
		t.Errorf("Error in validation: %s", err)
	}
}

/**
min-length, max-length validations test
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

/**
multiple validations test
*/
type MulStruct struct {
	Num int `json:"num" constraints:"required=true,nillable=true,multipleOf=5"`
}

func TestValConstraint(t *testing.T) {
	msg1 := MulStruct{
		Num: 10,
	}
	sv1 := validator.NewStructValidator()
	if err := sv1.Validate(msg1); err != nil {
		t.Errorf("Error in validation: %s", err)
	}

	msg2 := MulStruct{
		Num: 11,
	}
	sv2 := validator.NewStructValidator()
	err2 := sv2.Validate(msg2)
	got2 := err2.Error()
	want2 := "multipleOf validation failed"
	if got2 != want2 {
		t.Errorf("Expected: %s, got: %s", got2, want2)
	}
}

/**
required validations test
*/

type ReqStruct struct {
	Name string `json:"name" constraints:"required=true,nillable=false"`
}

func TestReqConstraints(t *testing.T) {
	msg1 := ReqStruct{
		Name: "abcd",
	}
	sv1 := validator.NewStructValidator()
	if err := sv1.Validate(msg1); err != nil {
		t.Errorf("Error in validation: %s", err)
	}

	msg2 := ReqStruct{
		Name: "",
	}
	sv2 := validator.NewStructValidator()
	err2 := sv2.Validate(msg2)
	got2 := err2.Error()
	want2 := "required validation failed"
	if got2 != want2 {
		t.Errorf("Expected: %s, got: %s", got2, want2)
	}
}

/**
pattern validations
*/

type PattStruct struct {
	Str string `json:"str" constraints:"required=true,nillable=false,pattern=^[tes]{4}.*"`
}

type PattStruct2 struct {
	Str2 string `json:"str2" constraints:"required=true,nillable=false,pattern=["`
}

func TestPatternConstraints(t *testing.T) {
	msg1 := PattStruct{
		Str: "test1234",
	}
	sv1 := validator.NewStructValidator()
	if err := sv1.Validate(msg1); err != nil {
		t.Errorf("Error in validation: %s", err)
	}

	msg2 := PattStruct{
		Str: "abcd1234",
	}
	sv2 := validator.NewStructValidator()
	err2 := sv2.Validate(msg2)
	got2 := err2.Error()
	want2 := "pattern validation failed"
	if got2 != want2 {
		t.Errorf("Expected: %s, got: %s", got2, want2)
	}

	msg3 := PattStruct2{
		Str2: "tsst1234",
	}
	sv3 := validator.NewStructValidator()
	err3 := sv3.Validate(msg3)
	got3 := err3.Error()
	want3 := "invalid constraint value"
	if got3 != want3 {
		t.Errorf("Expected: %s, got: %s", got3, want3)
	}
}

/**
Nested Structure Testing
*/

type Example struct {
	Reference
	Summary     string      `json:"summary,omitempty" constraints:"required=true,nillable=false"`
	Description string      `json:"description,omitempty" constraints:"required=true,nillable=false"`
	Value       interface{} `json:"example,omitempty" constraints:"required=true,nillable=false"`
}

type Reference struct {
	Ref         *string `json:"$ref,omitempty" constraints:"required=true,nillable=false"`
	Summary     string  `json:"summary,omitempty" constraints:"required=true,nillable=false"`
	Description string  `json:"description,omitempty" constraints:"required=true,nillable=false"`
}

func TestNested(t *testing.T) {
	msg := Example{
		Reference: Reference{
			Ref:         nil,
			Summary:     "ref summary",
			Description: "ref description",
		},
		Summary:     "summary",
		Description: "description",
		Value:       nil,
	}

	sv := validator.NewStructValidator()
	if err := sv.Validate(msg); err != nil {
		t.Errorf("Error in validation: %s", err)
	}
}

/**
Empty Struct Validation
// TODO
*/

type EmptyExample struct {
	Field string `json:"field" constraints:"required=false,nillable=false"`
}

func TestEmptyStruct(t *testing.T) {
	msg := EmptyExample{}
	sv := validator.NewStructValidator()
	if err := sv.Validate(msg); err != nil {
		t.Errorf("Error in validation: %s", err)
	}

}

type ConstExample struct {
	Summary     string `json:"summary"`
	Description string `json:"description" constraints:"required=false,nillable=false"`
}

func TestConstStruct(t *testing.T) {
	msg := ConstExample{
		Summary:     "testing",
		Description: "this is testing",
	}
	sv := validator.NewStructValidator()
	if err := sv.Validate(msg); err != nil {
		t.Errorf("Error in validation: %s", err)
	}
}
