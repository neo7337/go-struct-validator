package validator

import (
	"reflect"
	"strconv"
	"strings"
)

/**
constraints conversion from string
*/

func convertInt(param string, bit int) (int64, error) {
	i, err := strconv.ParseInt(param, 0, bit)
	if err != nil {
		return 0, ErrBadConstraint
	}
	return i, nil
}

func convertUint(param string, bit int) (uint64, error) {
	i, err := strconv.ParseUint(param, 0, bit)
	if err != nil {
		return 0, ErrBadConstraint
	}
	return i, nil
}

func convertFloat(param string, bit int) (float64, error) {
	i, err := strconv.ParseFloat(param, bit)
	if err != nil {
		return 0, ErrBadConstraint
	}
	return i, nil
}

func convertBool(param string) (bool, error) {
	i, err := strconv.ParseBool(param)
	if err != nil {
		return false, ErrBadConstraint
	}
	return i, nil
}

func checkMin(val reflect.Value, typ reflect.Type, param string, isExclusive bool) error {
	valid := true
	switch typ.Kind() {
	case reflect.Int:
		c, err := convertInt(param, 0)
		if err != nil {
			return err
		}
		cInt := int(c)
		in, _ := val.Interface().(int)
		if isExclusive {
			valid = in >= cInt
		} else {
			valid = in > cInt
		}
	case reflect.Int8:
		c, err := convertInt(param, 8)
		if err != nil {
			return err
		}
		cInt := int8(c)
		in, _ := val.Interface().(int8)
		if isExclusive {
			valid = in >= cInt
		} else {
			valid = in > cInt
		}
	case reflect.Int16:
		c, err := convertInt(param, 16)
		if err != nil {
			return err
		}
		cInt := int16(c)
		in, _ := val.Interface().(int16)
		if isExclusive {
			valid = in >= cInt
		} else {
			valid = in > cInt
		}
	case reflect.Int32:
		c, err := convertInt(param, 32)
		if err != nil {
			return err
		}
		cInt := int32(c)
		in, _ := val.Interface().(int32)
		if isExclusive {
			valid = in >= cInt
		} else {
			valid = in > cInt
		}
	case reflect.Int64:
		c, err := convertInt(param, 64)
		if err != nil {
			return err
		}
		in, _ := val.Interface().(int64)
		if isExclusive {
			valid = in >= c
		} else {
			valid = in > c
		}
	case reflect.Uint:
		c, err := convertUint(param, 0)
		if err != nil {
			return err
		}
		cUint := uint(c)
		in, _ := val.Interface().(uint)
		if isExclusive {
			valid = in >= cUint
		} else {
			valid = in > cUint
		}
	case reflect.Uint8:
		c, err := convertUint(param, 8)
		if err != nil {
			return err
		}
		cUint := uint8(c)
		in, _ := val.Interface().(uint8)
		if isExclusive {
			valid = in >= cUint
		} else {
			valid = in > cUint
		}
	case reflect.Uint16:
		c, err := convertUint(param, 16)
		if err != nil {
			return err
		}
		cUint := uint16(c)
		in, _ := val.Interface().(uint16)
		if isExclusive {
			valid = in >= cUint
		} else {
			valid = in > cUint
		}
	case reflect.Uint32:
		c, err := convertUint(param, 32)
		if err != nil {
			return err
		}
		cUint := uint32(c)
		in, _ := val.Interface().(uint32)
		if isExclusive {
			valid = in >= cUint
		} else {
			valid = in > cUint
		}
	case reflect.Uint64:
		c, err := convertUint(param, 64)
		if err != nil {
			return err
		}
		in, _ := val.Interface().(uint64)
		if isExclusive {
			valid = in >= c
		} else {
			valid = in > c
		}
	case reflect.Uintptr:
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
		if isExclusive {
			valid = in >= cFloat
		} else {
			valid = in > cFloat
		}
	case reflect.Float64:
		c, err := convertFloat(param, 64)
		if err != nil {
			return err
		}
		cFloat := c
		in, _ := val.Interface().(float64)
		if isExclusive {
			valid = in >= cFloat
		} else {
			valid = in > cFloat
		}
	}
	if !valid {
		if isExclusive {
			return ErrExclusiveMin
		} else {
			return ErrMin
		}
	}
	return nil
}

func checkMax(val reflect.Value, typ reflect.Type, param string, isExclusive bool) error {
	valid := true
	switch typ.Kind() {
	case reflect.Int:
		c, err := convertInt(param, 0)
		if err != nil {
			return err
		}
		cInt := int(c)
		in, _ := val.Interface().(int)
		if isExclusive {
			valid = in <= cInt
		} else {
			valid = in < cInt
		}
	case reflect.Int8:
		c, err := convertInt(param, 8)
		if err != nil {
			return err
		}
		cInt := int8(c)
		in, _ := val.Interface().(int8)
		if isExclusive {
			valid = in <= cInt
		} else {
			valid = in < cInt
		}
	case reflect.Int16:
		c, err := convertInt(param, 16)
		if err != nil {
			return err
		}
		cInt := int16(c)
		in, _ := val.Interface().(int16)
		if isExclusive {
			valid = in <= cInt
		} else {
			valid = in < cInt
		}
	case reflect.Int32:
		c, err := convertInt(param, 32)
		if err != nil {
			return err
		}
		cInt := int32(c)
		in, _ := val.Interface().(int32)
		if isExclusive {
			valid = in <= cInt
		} else {
			valid = in < cInt
		}
	case reflect.Int64:
		c, err := convertInt(param, 64)
		if err != nil {
			return err
		}
		in, _ := val.Interface().(int64)
		if isExclusive {
			valid = in <= c
		} else {
			valid = in < c
		}
	case reflect.Uint:
		c, err := convertUint(param, 0)
		if err != nil {
			return err
		}
		cUint := uint(c)
		in, _ := val.Interface().(uint)
		if isExclusive {
			valid = in <= cUint
		} else {
			valid = in < cUint
		}
	case reflect.Uint8:
		c, err := convertUint(param, 8)
		if err != nil {
			return err
		}
		cUint := uint8(c)
		in, _ := val.Interface().(uint8)
		if isExclusive {
			valid = in <= cUint
		} else {
			valid = in < cUint
		}
	case reflect.Uint16:
		c, err := convertUint(param, 16)
		if err != nil {
			return err
		}
		cUint := uint16(c)
		in, _ := val.Interface().(uint16)
		if isExclusive {
			valid = in <= cUint
		} else {
			valid = in < cUint
		}
	case reflect.Uint32:
		c, err := convertUint(param, 32)
		if err != nil {
			return err
		}
		cUint := uint32(c)
		in, _ := val.Interface().(uint32)
		if isExclusive {
			valid = in <= cUint
		} else {
			valid = in < cUint
		}
	case reflect.Uint64:
		c, err := convertUint(param, 64)
		if err != nil {
			return err
		}
		in, _ := val.Interface().(uint64)
		if isExclusive {
			valid = in <= c
		} else {
			valid = in < c
		}
	case reflect.Uintptr:
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
		if isExclusive {
			valid = in <= cFloat
		} else {
			valid = in < cFloat
		}
	case reflect.Float64:
		c, err := convertFloat(param, 64)
		if err != nil {
			return err
		}
		cFloat := c
		in, _ := val.Interface().(float64)
		if isExclusive {
			valid = in <= cFloat
		} else {
			valid = in < cFloat
		}
	}
	if !valid {
		if isExclusive {
			return ErrExclusiveMax
		} else {
			return ErrMax
		}
	}
	return nil
}

func checkIfEnumExists(val string, param string, separator string) bool {
	flag := false
	params := strings.Split(param, separator)
	for _, en := range params {
		if val == en {
			flag = true
		}
	}
	if flag == false {
		return false
	}
	return true
}
