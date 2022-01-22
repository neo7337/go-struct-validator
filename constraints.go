package validator

import (
	"fmt"
	"math"
	"reflect"
	"strconv"
)

/**
Base Constraints for all Data Types
*/

func required(val interface{}, typ reflect.Type, param string) error {
	v, _ := val.(string)
	fmt.Println(v)
	// c, _ := strconv.ParseBool(param)
	if val == nil {
		return ErrRequired
	}
	return nil
}

func nillable(v interface{}, typ reflect.Type, param string) error {
	return nil
}

func def(v interface{}, typ reflect.Type, param string) error {
	return nil
}

/**
Numerical Type Constraints
*/

func min(val interface{}, typ reflect.Type, param string) error {
	input := reflect.ValueOf(val)
	valid := true
	fmt.Println(input.Convert(typ))
	switch typ.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		c, err := convertInt(param)
		if err != nil {
			return err
		}
		in := val.(int)
		fmt.Println(in)
		fmt.Println(c)
		//valid = in < c
		fmt.Println(valid)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		c, err := convertUint(param)
		if err != nil {
			return err
		}
		valid = input.Uint() < c
	case reflect.Float32, reflect.Float64:
		c, err := convertFloat(param)
		if err != nil {
			return err
		}
		valid = input.Float() < c
	}
	if !valid {
		return ErrMin
	}
	return nil
}

func max(val interface{}, typ reflect.Type, param string) error {
	v, _ := strconv.Atoi(fmt.Sprintln(val))
	c, _ := strconv.Atoi(param)
	if v < c {
		return ErrMax
	}
	return nil
}

func exclusiveMin(val interface{}, typ reflect.Type, param string) error {
	valid := true
	input := reflect.ValueOf(val)
	switch input.Kind() {
	case reflect.Int:
		v := input.Int()
		if v < MinInt {
			valid = false
		}
	case reflect.Int8:
		v := input.Int()
		if v < math.MinInt8 {
			valid = false
		}
	case reflect.Int16:
		v := input.Int()
		if v < math.MinInt16 {
			valid = false
		}
	case reflect.Int32:
		v := input.Int()
		if v < math.MinInt32 {
			valid = false
		}
	case reflect.Int64:
		v := input.Int()
		if v < math.MinInt64 {
			valid = false
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		v := input.Int()
		if v < 0 {
			valid = false
		}
	}

	if !valid {
		return ErrExclusiveMin
	}

	return nil
}

func exclusiveMax(val interface{}, typ reflect.Type, param string) error {
	valid := true
	input := reflect.ValueOf(val)
	switch input.Kind() {
	case reflect.Int:
		v := input.Int()
		if v > MaxInt {
			valid = false
		}
	case reflect.Int8:
		v := input.Int()
		if v > math.MaxInt8 {
			valid = false
		}
	case reflect.Int16:
		v := input.Int()
		if v > math.MaxInt16 {
			valid = false
		}
	case reflect.Int32:
		v := input.Int()
		if v > math.MaxInt32 {
			valid = false
		}
	case reflect.Int64:
		v := input.Int()
		if v > math.MaxInt64 {
			valid = false
		}
	case reflect.Uint:
		v := input.Uint()
		if v > MaxUint {
			valid = false
		}
	case reflect.Uint8:
		v := input.Uint()
		if v > math.MaxUint8 {
			valid = false
		}
	case reflect.Uint16:
		v := input.Uint()
		if v > math.MaxUint16 {
			valid = false
		}
	case reflect.Uint32:
		v := input.Uint()
		if v > math.MaxUint32 {
			valid = false
		}
	case reflect.Uint64:
		v := input.Uint()
		if v > math.MaxUint64 {
			valid = false
		}
	}

	if !valid {
		return ErrExclusiveMax
	}

	return nil
}

func multipleOf(v interface{}, typ reflect.Type, param string) error {
	return nil
}

/**
String Type Constraints
*/

func minLength(val interface{}, typ reflect.Type, param string) error {
	lc, _ := strconv.Atoi(param)
	lv := len(fmt.Sprint(val))
	if lv < lc {
		fmt.Println("error")
		return ErrMinLength
	}
	return nil
}

func maxLength(val interface{}, typ reflect.Type, param string) error {
	lc, _ := strconv.Atoi(param)
	lv := len(fmt.Sprint(val))
	if lv > lc {
		return ErrMaxLength
	}
	return nil
}

func pattern(v interface{}, typ reflect.Type, param string) error {
	return nil
}

/**
utils
*/
/**
constraints conversion from string
*/
func convertInt(param string) (int64, error) {
	i, err := strconv.ParseInt(param, 0, 64)
	if err != nil {
		return 0, ErrBadConstraint
	}
	return i, nil
}

func convertUint(param string) (uint64, error) {
	i, err := strconv.ParseUint(param, 0, 64)
	if err != nil {
		return 0, ErrBadConstraint
	}
	return i, nil
}

func convertFloat(param string) (float64, error) {
	i, err := strconv.ParseFloat(param, 64)
	if err != nil {
		return 0, ErrBadConstraint
	}
	return i, nil
}

func convertToDesiredType() {

}