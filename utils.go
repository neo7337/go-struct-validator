package validator

import (
	"strconv"
)

/**
utils
*/
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
	i, err := strconv.ParseUint(param, 0, 64)
	if err != nil {
		return 0, ErrBadConstraint
	}
	return i, nil
}

func convertFloat(param string, bit int) (float64, error) {
	i, err := strconv.ParseFloat(param, 64)
	if err != nil {
		return 0, ErrBadConstraint
	}
	return i, nil
}
