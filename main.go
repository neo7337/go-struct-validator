package main

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

func Validate(v interface{}) error {

	value := reflect.ValueOf(v)
	typ := value.Type()

	for i := 0; i < typ.NumField(); i++ {
		f := typ.Field(i)

		tag := f.Tag.Get("constraints")
		constraints := parseTag(tag)
		fieldValue := value.Field(i)
		// fmt.Println("fieldName", f.Name)
		// fmt.Println("fieldValue", value.Field(i))
		// fmt.Println("fieldType", f.Type.Kind())
		// fmt.Println("tagNames", constraints)

		if err := executeValidators(fieldValue, f.Type, constraints); err != nil {
			return err
		}
	}

	return nil
}

func parseTag(tag string) map[string]string {
	m := make(map[string]string)
	split := strings.Split(tag, ",")
	for _, str := range split {
		constraintName := strings.Split(str, ":")[0]
		constraintValue := strings.Split(str, ":")[1]
		m[constraintName] = constraintValue
	}
	return m
}

// type validatorFunc func(value reflect.Value, constraint map[string]string)

func executeValidators(value reflect.Value, typ reflect.Type, constraint map[string]string) error {
	switch typ.Kind() {
	case reflect.Bool:
		return boolValidator(value, constraint)
	case reflect.String:
		return stringValidator(value, constraint)
	default:
		return invalidValidator(value, constraint)
	}
}

func stringValidator(value reflect.Value, constraint map[string]string) error {

	// constraints to be predefined and mapped to a particular validation
	fmt.Println("executing validator", value)
	for i, val := range constraint {
		fmt.Println(i, val)
		le, _ := strconv.Atoi(val)
		lenF := len(value.String())

		if lenF < le {
			return errors.New("validation failed")
		}

	}
	return nil
}

func boolValidator(value reflect.Value, constraint map[string]string) error {
	return nil
}

func invalidValidator(value reflect.Value, constraint map[string]string) error {
	return nil
}