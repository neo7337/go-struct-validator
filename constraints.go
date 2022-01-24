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

func required(val reflect.Value, typ reflect.Type, param string) error {
	//v, _ := val.(string)
	//fmt.Println(v)
	//// c, _ := strconv.ParseBool(param)
	//if val == nil {
	//	return ErrRequired
	//}
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
	valid := true
	switch typ.Kind() {
	case reflect.Int:
		c, err := convertInt(param, 0)
		if err != nil {
			return err
		}
		cInt := int(c)
		in, _ := val.Interface().(int)
		valid = in > cInt
	case reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		/*c, err := convertInt(param)
		if err != nil {
			return err
		}
		in := val.Interface().(int8)
		valid = in > c*/
		valid = true
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		/*c, err := convertUint(param)
		if err != nil {
			return err
		}
		valid = input.Uint() < c*/
		valid = true
	case reflect.Float32:
		/*c, err := convertFloat(param)
		if err != nil {
			return err
		}
		valid = input.Float() < c*/
		valid = true
	case reflect.Float64:
		valid = true
	}
	if !valid {
		return ErrMin
	}
	return nil
}

func max(val reflect.Value, typ reflect.Type, param string) error {
	valid := true
	switch typ.Kind() {
	case reflect.Int:
		c, err := convertInt(param, 0)
		if err != nil {
			return err
		}
		cInt := int(c)
		in, _ := val.Interface().(int)
		valid = in < cInt
	case reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		/*c, err := convertInt(param)
		if err != nil {
			return err
		}
		in := val.Interface().(int8)
		valid = in > c*/
		valid = true
	case reflect.Uint:
		c, err := convertUint(param, 0)
		if err != nil {
			return err
		}
		cUint := uint(c)
		in, _ := val.Interface().(uint)
		valid = in < cUint
		valid = true
	case reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		/*c, err := convertUint(param)
		if err != nil {
			return err
		}
		valid = input.Uint() < c*/
		valid = true
	case reflect.Float32:
		c, err := convertFloat(param, 32)
		if err != nil {
			return err
		}
		cFloat := float32(c)
		in, _ := val.Interface().(float32)
		valid = in < cFloat
	case reflect.Float64:
		c, err := convertFloat(param, 64)
		if err != nil {
			return err
		}
		cFloat := c
		in, _ := val.Interface().(float64)
		valid = in < cFloat
	}
	if !valid {
		return ErrMax
	}
	return nil
}

func exclusiveMin(val reflect.Value, typ reflect.Type, param string) error {
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

func exclusiveMax(val reflect.Value, typ reflect.Type, param string) error {
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
	return nil
}
