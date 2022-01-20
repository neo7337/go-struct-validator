package validator_test

import (
	"testing"
	"validator"
)

type Message struct {
	Name string `json:"name" constraints:"min-length=5"`
	Age  int    `json:"age" constraints:"min=10"`
}

// negative test case
func TestRequiredConstraintFail(t *testing.T) {
	msg := Message{
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

// Positive Test Case
func TestSuccessValidation(t *testing.T) {
	msg := Message{
		Name: "Testy",
		Age:  11,
	}
	sv := validator.NewStructValidator()
	if err := sv.Validate(msg); err != nil {
		t.Errorf("Error in validation: %d", err)
	}
}

type BoolMessage struct {
	Present bool `json:"present" constraints:"required=true,nillable=true"`
}

func TestRequiredConstraintSuccess(t *testing.T) {
	msg := BoolMessage{
		Present: true,
	}
	sv := validator.NewStructValidator()
	if err := sv.Validate(msg); err != nil {
		t.Errorf("Error in validation: %d", err)
	}
}

type DefMessage struct {
	Count int `json:"count" constraints:"required=true,nillable=true,default=0"`
}

func TestDefConstraint(t *testing.T) {
	msg := DefMessage{
		Count: 0,
	}
	sv := validator.NewStructValidator()
	if err := sv.Validate(msg); err != nil {
		t.Errorf("Error in validation: %d", err)
	}
}
