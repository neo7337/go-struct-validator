package validator_test

import (
	"testing"
	"validator"
)

type Message struct {
	Name string `json:"name" constraints:"min-length=5"`
	Age  int    `json:"age" constraints:"required=true,nillable=true,min=0"`
}

func TestNewStructValidator(t *testing.T) {
	msg := Message{
		Name: "Test",
		Age:  21,
	}
	sv := validator.NewStructValidator()
	if err := sv.Validate(msg); err != nil {
		t.Errorf("Error in validation: %d", err)
	}
}
