package validator

import "errors"

type ErrMsg struct {
	Err error
}

func (e ErrMsg) Error() string {
	return e.Err.Error()
}

var (
	ErrRequired = ErrMsg{errors.New("required validation failed")}

	ErrNillable = ErrMsg{errors.New("nillable validation failed")}

	ErrDefault = ErrMsg{errors.New("default validation failed")}

	ErrMin = ErrMsg{errors.New("min value validation failed")}

	ErrMax = ErrMsg{errors.New("max value validation failed")}

	ErrExclusiveMin = ErrMsg{errors.New("exclusive min validation failed")}

	ErrExclusiveMax = ErrMsg{errors.New("exclusive max validation failed")}

	ErrMultipleOf = ErrMsg{errors.New("multipleOf validation failed")}

	ErrMinLength = ErrMsg{errors.New("min-length validation failed")}

	ErrMaxLength = ErrMsg{errors.New("max-length validation failed")}

	ErrPattern = ErrMsg{errors.New("pattern validation failed")}

	ErrBadConstraint = ErrMsg{errors.New("invalid constraint value")}
)
