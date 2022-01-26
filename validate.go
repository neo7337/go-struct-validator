package validator

import (
	"errors"
	"fmt"
	"go.nandlabs.io/l3"
	"reflect"
	"strings"
)

var logger = l3.Get()

var mandatory = [...]string{"required", "nillable"}

type StructValidatorFunc func(v reflect.Value, typ reflect.Type, param string) error

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
			// <, > only
			"min": min,
			"max": max,
			// <=, >= this is inclusive of the input value
			"exclusiveMin": exclusiveMin,
			"exclusiveMax": exclusiveMax,
			"multipleOf":   multipleOf,
			// String Constraints
			"min-length": minLength,
			"max-length": maxLength,
			// regex pattern support
			"pattern": pattern,
		},
		tagName: "constraints",
	}
}

func (sv *StructValidator) Validate(v interface{}) error {

	// add a logic to check for the empty struct input in order to skip the validation of the struct
	logger.Info("starting struct validation")
	value := reflect.ValueOf(v)
	typ := value.Type()
	for i := 0; i < typ.NumField(); i++ {
		f := typ.Field(i)
		tag := f.Tag.Get("constraints")
		constraints := parseTag(tag)
		fieldValue := value.Field(i)
		if err := sv.checkIfMandatoryTagPresent(constraints); err != nil {
			return err
		}
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
		if err := sv.validationFuncs[i](value, typ, v); err != nil {
			logger.ErrorF("constraint validation failed")
			return err
		} else {
			continue
		}
	}
	return nil
}

func (sv *StructValidator) checkIfMandatoryTagPresent(constraint map[string]string) error {
	for _, v := range mandatory {
		if _, ok := constraint[v]; !ok {
			logger.ErrorF("mandatory field `%s` not present", v)
			return errors.New(fmt.Sprintf("mandatory field %s not present", v))
		}
	}
	return nil
}
