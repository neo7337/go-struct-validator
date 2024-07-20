package validator

import (
	"reflect"

	"oss.nandlabs.io/golly/l3"
)

var logger = l3.Get()

type StructValidatorFunc func(field field, param string) error

type tStruct struct {
	name  string
	value string
	fnc   StructValidatorFunc
}

type field struct {
	name        string
	value       reflect.Value
	typ         reflect.Type
	index       []int
	constraints []tStruct
	inter       interface{}
}

type structFields struct {
	list []field
}

type StructValidator struct {
	fields         structFields
	validationFunc map[string]StructValidatorFunc
	tagName        string
	enableCache    bool
}

// NewStructValidator : Generates the new StructValidator object
func NewStructValidator() *StructValidator {
	return &StructValidator{
		validationFunc: map[string]StructValidatorFunc{
			// Base Constraints
			// Numeric Constraints
			// <, > only
			"min": min,
			"max": max,
			// <=, >= this is inclusive of the input value
			"exclusiveMin": exclusiveMin,
			"exclusiveMax": exclusiveMax,
			"multipleOf":   multipleOf,
			// String Constraints
			// boolean value
			"notnull":    notnull,
			"min-length": minLength,
			"max-length": maxLength,
			// regex pattern support
			"pattern": pattern,
			// enums support
			"enum": enum,
		},
		tagName:     "constraints",
		enableCache: false,
	}
}

// SetTagName : provide the custom tag-name for your constraints
// By default : 'constraints' tag-name is used
func (sv *StructValidator) SetTagName(tag string) *StructValidator {
	sv.tagName = tag
	return sv
}

// SetCache : provide true/false to enable or disable cache in the validator
func (sv *StructValidator) SetCache(action bool) *StructValidator {
	sv.enableCache = action
	return sv
}

// Validate : Runs the core of logic of the validations schema
func (sv *StructValidator) Validate(v interface{}) error {
	//check for cache
	sv.fields = sv.cachedTypeFields(v)
	if err := sv.validateFields(); err != nil {
		return err
	}
	return nil
}
