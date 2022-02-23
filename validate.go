package validator

import (
	"go.nandlabs.io/l3"
	"reflect"
	"regexp"
	"strings"
	"sync"
)

var logger = l3.Get()

var mapMandatory = map[string]string{
	"required": "required",
	"nillable": "nillable",
}

type StructValidatorFunc func(v reflect.Value, typ reflect.Type, param string) error

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
			// enums support
			"enum": enum,
		},
		tagName: "constraints",
	}
}

func (sv *StructValidator) Validate(v interface{}) error {
	//logger.Info("starting struct validation")
	// add a logic to check for the empty struct input in order to skip the validation of the struct

	//check for cache
	sv.fields = sv.cachedTypeFields(v)
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
		for _, val := range v.constraints {
			if err := val.fnc(v.value, v.typ, val.value); err != nil {
				return err
			}
		}
	}
	return nil
}

func checkForMandatory(constraint []tStruct) error {
	count := 0
	for _, t := range constraint {
		if _, ok := mapMandatory[t.name]; ok {
			count++
		}
	}
	if count != len(mapMandatory) {
		return ErrMandatoryFields
	}
	return nil
}

//parseTag returns the map of constraints
func (sv *StructValidator) parseTag(tag string) ([]tStruct, error) {
	tl := splitUnescapedComma(tag)
	t := make([]tStruct, 0, len(tl))

	for _, s := range tl {
		s = strings.Replace(s, `\,`, ",", -1)
		tg := tStruct{}
		v := strings.SplitN(s, "=", 2)
		tg.name = strings.Trim(v[0], " ")
		//check for blank tag name
		if len(v) > 1 {
			tg.value = strings.Trim(v[1], " ")
		}
		tg.fnc, _ = sv.validationFunc[tg.name]
		// check for not found
		t = append(t, tg)
	}

	return t, nil
}

var sepPattern = regexp.MustCompile(`((?:^|[^\\])(?:\\\\)*);`)

func splitUnescapedComma(str string) []string {
	ret := []string{}
	indexes := sepPattern.FindAllStringIndex(str, -1)
	last := 0
	for _, is := range indexes {
		ret = append(ret, str[last:is[1]-1])
		last = is[1]
	}
	ret = append(ret, str[last:])
	return ret
}

// reference from go encoder
func (sv *StructValidator) parseFields(v interface{}) structFields {

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
				consts, _ := sv.parseTag(tag)
				// add check for error

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

func (sv *StructValidator) cachedTypeFields(v interface{}) structFields {
	t := reflect.ValueOf(v).Type()
	if f, ok := fieldCache.Load(t); ok {
		return f.(structFields)
	}
	f, _ := fieldCache.LoadOrStore(t, sv.parseFields(v))
	return f.(structFields)
}
