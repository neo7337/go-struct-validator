package validator

import (
	"fmt"
	"reflect"
	"regexp"
	"strconv"
)

/**
Base Constraints for all Data Types
*/

func required(val reflect.Value, typ reflect.Type, param string) error {
	switch typ.Kind() {
	case reflect.String:
		c, err := convertBool(param)
		if err != nil {
			return err
		}
		if c == true {
			in, _ := val.Interface().(string)
			if in == "" {
				return ErrRequired
			}
		}
	case reflect.Bool:
	case reflect.Int:
	case reflect.Float32:
	case reflect.Uint:
	}
	return nil
}

func nillable(val reflect.Value, typ reflect.Type, param string) error {
	return nil
}

func def(val reflect.Value, typ reflect.Type, param string) error {
	return nil
}

/**
Numerical Type Constraints
*/

func min(val reflect.Value, typ reflect.Type, param string) error {
	return checkMin(val, typ, param, false)
}

func max(val reflect.Value, typ reflect.Type, param string) error {
	return checkMax(val, typ, param, false)
}

/**
move the below functions to a generic function to consider the both min and exclusive-min
*/
func exclusiveMin(val reflect.Value, typ reflect.Type, param string) error {
	return checkMin(val, typ, param, true)
}

func exclusiveMax(val reflect.Value, typ reflect.Type, param string) error {
	return checkMax(val, typ, param, true)
}

func multipleOf(val reflect.Value, typ reflect.Type, param string) error {
	valid := true
	in, _ := val.Interface().(int)
	c, err := convertInt(param, 0)
	cInt := int(c)
	if err != nil {
		return err
	}
	valid = in%cInt == 0
	if !valid {
		return ErrMultipleOf
	}
	return nil
}

/**
String Type Constraints
*/

func minLength(val reflect.Value, typ reflect.Type, param string) error {
	valid := true
	lc, _ := strconv.Atoi(param)
	lv := len(fmt.Sprint(val))
	valid = lv > lc
	if !valid {
		return ErrMinLength
	}
	return nil
}

func maxLength(val reflect.Value, typ reflect.Type, param string) error {
	valid := true
	lc, _ := strconv.Atoi(param)
	lv := len(fmt.Sprint(val))
	valid = lv < lc
	if !valid {
		return ErrMaxLength
	}
	return nil
}

func pattern(val reflect.Value, typ reflect.Type, param string) error {
	in, _ := val.Interface().(string)
	if typ.Kind() != reflect.String {
		return ErrNotSupported
	}
	re, err := regexp.Compile(param)
	if err != nil {
		return ErrBadConstraint
	}
	if !re.MatchString(in) {
		return ErrPattern
	}
	return nil
}
