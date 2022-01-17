package validator

import (
	"fmt"
	"reflect"
	"strings"
)

type StructValidatorFunc func(v interface{}, param string) error

type StructValidator struct {
	validationFuncs map[string]StructValidatorFunc
	tagName         string
}

func NewStructValidator() *StructValidator {
	return &StructValidator{
		validationFuncs: map[string]StructValidatorFunc{
			// Base Constraints
			// boolean value
			// mandatory field
			"required": required,
			// boolean value
			// mandatory field
			"nillable": nillable,
			"default":  def,
			// Numeric Constraints
			"min":          min,
			"max":          max,
			"exclusiveMin": exclusiveMin,
			"exclusiveMax": exclusiveMax,
			"multipleOf":   multipleOf,
			// String Constraints
			"min-length": minLength,
			"max-length": maxLength,
			"pattern":    pattern,
		},
		tagName: "constraints",
	}
}

func (sv *StructValidator) Validate(v interface{}) error {

	value := reflect.ValueOf(v)
	typ := value.Type()
	fmt.Println(value)
	for i := 0; i < typ.NumField(); i++ {
		f := typ.Field(i)
		fmt.Println(f)
		fmt.Println("in loop")
		tag := f.Tag.Get("constraints")
		fmt.Println(tag)
		constraints := parseTag(tag)
		fieldValue := value.Field(i)
		fmt.Println("fieldName", f.Name)
		fmt.Println("fieldValue", value.Field(i))
		fmt.Println("fieldType", f.Type.Kind())
		fmt.Println("tagNames", constraints)

		if err := sv.executeValidators(fieldValue, f.Type, constraints); err != nil {
			return err
		}
	}

	return nil
}

func parseTag(tag string) map[string]string {
	m := make(map[string]string)
	split := strings.Split(tag, ",")
	for _, str := range split {
		constraintName := strings.Split(str, "=")[0]
		constraintValue := strings.Split(str, "=")[1]
		m[constraintName] = constraintValue
	}
	return m
}

func (sv *StructValidator) executeValidators(value reflect.Value, typ reflect.Type, constraint map[string]string) error {
	for i, v := range constraint {
		if err := sv.validationFuncs[i](value, v); err != nil {
			return err
		} else {
			continue
		}
	}
	return nil
}
