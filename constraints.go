package validator

import (
	"errors"
	"fmt"
	"strconv"
)

/**
Base Constraints for all Data Types
*/

func required(v interface{}, param string) error {
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

func min(v interface{}, param string) error {
	return nil
}

func max(v interface{}, param string) error {
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
	fmt.Println("minlength func")
	lc, _ := strconv.Atoi(param)
	lv := len(fmt.Sprint(val))
	fmt.Println(lc)
	fmt.Println(lv)
	if lv < lc {
		fmt.Println("error")
		return errors.New("min-length validation failed")
	}
	return nil
}

func maxLength(val interface{}, param string) error {
	lc, _ := strconv.Atoi(param)
	lv := len(fmt.Sprint(val))
	if lv > lc {
		return errors.New("max-length validation failed")
	}
	return nil
}

func pattern(v interface{}, param string) error {
	return nil
}
