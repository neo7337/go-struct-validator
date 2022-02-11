package validator

import (
	"go.nandlabs.io/l3"
	"reflect"
	"strings"
	"sync"
)

var logger = l3.Get()

var mandatory = [...]string{"required", "nillable"}

type StructValidatorFunc func(v reflect.Value, typ reflect.Type, param string) error

type field struct {
	name        string
	constraints sync.Map //map[string]string
}

type structFields struct {
	list []field
}

type StructValidator struct {
	fields         structFields
	validationFunc map[string]StructValidatorFunc
	tagName        string
}

var validatorCache sync.Map

func NewStructValidator() *StructValidator {
	return &StructValidator{
		validationFunc: map[string]StructValidatorFunc{
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

	//check for cache
	val := StructValidator{fields: cachedTypeFields(reflect.ValueOf(v).Type())}
	if err := val.validateFields(); err != nil {
		return err
	}

	if err := sv.deepFields(v); err != nil {
		return err
	}
	return nil
}

func (sv *StructValidator) validateFields() error {
	for i, v = range sv.fields.list {

	}
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
		if err := sv.validationFunc[constraintName](fieldValue, typ, constraintValue); err != nil {
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

func parseFields(t reflect.Type) structFields {

}

var fieldCache sync.Map //map[reflect.Type]structFields

func cachedTypeFields(t reflect.Type) structFields {
	if f, ok := fieldCache.Load(t); ok {
		return f.(structFields)
	}
	f, _ := fieldCache.LoadOrStore(t, parseFields(t))
	return f.(structFields)
}
