package validator

import (
	"fmt"
	"reflect"
	"regexp"
	"strings"
	"sync"
)

func (sv *StructValidator) validateFields() error {
	for _, field := range sv.fields.list {
		// check if tag-name is present or not, skip any kind of validation for which the constraints are not passed
		if (reflect.DeepEqual(field.constraints[0], tStruct{})) {
			logger.InfoF("skipping validation check for field: %s", field.name)
			continue
		}
		for _, val := range field.constraints {
			if err := val.fnc(field, val.value); err != nil {
				return err
			}
		}
	}
	return nil
}

// parseTag returns the map of constraints
func (sv *StructValidator) parseTag(tag string) ([]tStruct, error) {
	tl := splitUnescapedComma(tag)
	t := make([]tStruct, len(tl))

	for i, s := range tl {
		s = strings.Replace(s, `\,`, ",", -1)
		v := strings.SplitN(s, "=", 2)
		name := strings.Trim(v[0], " ")

		// Check for blank tag name
		if len(v) > 1 {
			value := strings.Trim(v[1], " ")

			// Lookup validation function and populate tStruct
			if fnc, found := sv.validationFunc[name]; found {
				t[i] = tStruct{name, value, fnc}
			} else {
				return nil, fmt.Errorf("validation function not found for tag: %s", name)
			}
		} else {
			// No value provided, simply set the name
			t[i] = tStruct{name: name}
		}
	}

	return t, nil
}

var sepPattern = regexp.MustCompile(`((?:^|[^\\])(?:\\\\)*);`)

func splitUnescapedComma(str string) []string {
	var ret []string
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

	var current []field
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
				tag := sf.Tag.Get(sv.tagName)
				// if the tag-name tag is -, skip the field validation
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
	if sv.enableCache {
		t := reflect.ValueOf(v).Type()
		if f, ok := fieldCache.Load(t); ok {
			return f.(structFields)
		}
		f, _ := fieldCache.LoadOrStore(t, sv.parseFields(v))
		return f.(structFields)
	}
	f := sv.parseFields(v)
	return f
}
