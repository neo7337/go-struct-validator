package validator

import (
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
	//logger.Info("starting struct validation")
	// add a logic to check for the empty struct input in order to skip the validation of the struct
	if err := sv.deepFields(v); err != nil {
		return err
	}
	return nil
}

func (sv *StructValidator) deepFields(itr interface{}) error {
	ifv := reflect.ValueOf(itr)
	ift := ifv.Type()
	for i := 0; i < ift.NumField(); i++ {
		vi := ifv.Field(i)
		v := ift.Field(i)
		switch v.Type.Kind() {
		case reflect.Struct:
			if err := sv.deepFields(vi.Interface()); err != nil {
				return err
			}
		default:
			tag := v.Tag.Get("constraints")
			if tag == "" {
				logger.InfoF("constraint not present for field : %s, skip to next field", v.Name)
				continue
			}
			fieldValue := ifv.Field(i)
			if err := sv.parseTag(fieldValue, tag, v.Type); err != nil {
				return err
			}
		}
	}
	return nil
}

func (sv *StructValidator) parseTag(fieldValue reflect.Value, tag string, typ reflect.Type) error {
	split := strings.Split(tag, ",")
	// fix this logic to check the mandatory tags
	if err := check(split); err != true {
		return ErrMandatoryFields
	}
	for _, str := range split {
		constraintName := strings.Split(str, "=")[0]
		constraintValue := strings.Split(str, "=")[1]
		//m[constraintName] = constraintValue
		if err := sv.validationFuncs[constraintName](fieldValue, typ, constraintValue); err != nil {
			logger.ErrorF("constraint validation failed")
			return err
		} else {
			continue
		}
	}
	return nil
}

func check(list []string) bool {
	count := 0
	for _, v := range list {
		for _, m := range mandatory {
			if m == strings.Split(v, "=")[0] {
				count++
			}
		}
	}
	if count == len(mandatory) {
		return true
	}
	return false
}
