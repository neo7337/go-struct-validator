package validator

import (
	"fmt"
	"strconv"
)

/**
Base Constraints for all Data Types
*/

func required(val interface{}, param string) error {
	v, _ := val.(string)
	fmt.Println(v)
	// c, _ := strconv.ParseBool(param)
	if val == nil {
		return ErrRequired
	}
	return nil
}

func nillable(v interface{}, param string) error {
	return nil
}

func def(v interface{}, param string) error {
	return nil
}

/**
Numerical Type Constraints
*/

func min(val interface{}, param string) error {
	v, _ := strconv.Atoi(fmt.Sprintln(val))
	c, _ := strconv.Atoi(param)
	if v > c {
		return ErrMin
	}
	return nil
}

func max(val interface{}, param string) error {
	v, _ := strconv.Atoi(fmt.Sprintln(val))
	c, _ := strconv.Atoi(param)
	if v < c {
		return ErrMax
	}
	return nil
}

func exclusiveMin(v interface{}, param string) error {
	return nil
}

func exclusiveMax(v interface{}, param string) error {
	return nil
}

func multipleOf(v interface{}, param string) error {
	return nil
}

/**
String Type Constraints
*/

func minLength(val interface{}, param string) error {
	lc, _ := strconv.Atoi(param)
	lv := len(fmt.Sprint(val))
	if lv < lc {
		fmt.Println("error")
		return ErrMinLength
	}
	return nil
}

func maxLength(val interface{}, param string) error {
	lc, _ := strconv.Atoi(param)
	lv := len(fmt.Sprint(val))
	if lv > lc {
		return ErrMaxLength
	}
	return nil
}

func pattern(v interface{}, param string) error {
	return nil
}
