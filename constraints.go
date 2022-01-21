package validator

import (
	"fmt"
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
	// based on type it can be
	//int
	//int8
	//int16
	//int32
	//int64
	//uint
	//uint8
	//uint16
	//uint32
	//uint64
	//uintptr
	//byte - alias for uint8
	//rune
	v, _ := strconv.Atoi(fmt.Sprintln(val))
	c, _ := strconv.Atoi(param)
	if v > c {
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

func exclusiveMin(v interface{}, typ reflect.Type, param string) error {
	switch typ.Kind() {
	case reflect.Int:
	case reflect.Int8:
	case reflect.Int16:
	case reflect.Int32:
	case reflect.Int64:
	case reflect.Uint:
	case reflect.Uint8:
	case reflect.Uint16:
	case reflect.Uint32:
	case reflect.Uint64:
	case reflect.Uintptr:
	default:

	}
	return nil
}

func exclusiveMax(v interface{}, typ reflect.Type, param string) error {
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
