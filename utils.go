package validator

import (
	"strconv"
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
