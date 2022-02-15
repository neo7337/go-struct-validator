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
	value       reflect.Value
	typ         reflect.Type
	index       []int
	constraints map[string]string
	inter       interface{}
}

type structFields struct {
	list []field
}

type StructValidator struct {
	fields         structFields
	validationFunc map[string]StructValidatorFunc
	tagName        string
}

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
	sv.fields = cachedTypeFields(v)
	if err := sv.validateFields(); err != nil {
		return err
	}
	return nil
}

func (sv *StructValidator) validateFields() error {
	for _, v := range sv.fields.list {
		if err := checkForMandatory(v.constraints); err != nil {
			return err
		}
		for k, val := range v.constraints {
			if err := sv.validationFunc[k](v.value, v.typ, val); err != nil {
				return err
			}
		}
	}
	return nil
}

func checkForMandatory(constraint map[string]string) error {
	for _, v := range mandatory {
		if _, ok := constraint[v]; !ok {
			return ErrMandatoryFields
		}
	}
	return nil
}

//parseTag returns the map of constraints
func parseTag(tag string) map[string]string {
	m := make(map[string]string)
	if tag == "" {
		return m
	}
	split := strings.Split(tag, ",")
	for _, str := range split {
		constraintName := strings.Split(str, "=")[0]
		constraintValue := strings.Split(str, "=")[1]
		m[constraintName] = constraintValue
	}
	return m
}

// reference from go encoder
func parseFields(v interface{}) structFields {

	t := reflect.ValueOf(v).Type()
	fv := reflect.ValueOf(v)

	current := []field{}
	next := []field{{typ: t}}

	var count, nextCount map[reflect.Type]int

	visited := map[reflect.Type]bool{}

	var fields []field

	for len(next) > 0 {
		current, next = next, current[:0]
		count, nextCount = nextCount, map[reflect.Type]int{}

		for _, f := range current {
			if visited[f.typ] {
				continue
			}
			visited[f.typ] = true

			for i := 0; i < f.typ.NumField(); i++ {
				sf := f.typ.Field(i)
				if sf.Anonymous {
					t := sf.Type
					if t.Kind() == reflect.Ptr {
						t = t.Elem()
					}
				}
				tag := sf.Tag.Get("constraints")
				// if the constraints tag is -, skip the field validation
				if tag == "-" {
					continue
				}
				consts := parseTag(tag)

				index := make([]int, len(f.index)+1)
				copy(index, f.index)
				index[len(f.index)] = i

				ft := sf.Type
				if ft.Name() == "" && ft.Kind() == reflect.Ptr {
					ft = ft.Elem()
				}

				var val reflect.Value
				if !sf.Anonymous || ft.Kind() != reflect.Struct {
					if f.inter != nil {
						val = reflect.ValueOf(f.inter).Field(i)
					} else {
						val = fv.Field(i)
					}
					field := field{
						name:        sf.Name,
						typ:         ft,
						constraints: consts,
						value:       val,
					}
					fields = append(fields, field)
					if count[f.typ] > 1 {
						fields = append(fields, fields[len(fields)-1])
					}
					continue
				}

				nextCount[ft]++
				if nextCount[ft] == 1 {
					next = append(next, field{name: ft.Name(), index: index, typ: ft, inter: fv.Field(i).Interface()})
				}
			}
		}
	}
	return structFields{fields}
}

var fieldCache sync.Map //map[reflect.Type]structFields

func cachedTypeFields(v interface{}) structFields {
	t := reflect.ValueOf(v).Type()
	if f, ok := fieldCache.Load(t); ok {
		return f.(structFields)
	}
	f, _ := fieldCache.LoadOrStore(t, parseFields(v))
	return f.(structFields)
}
