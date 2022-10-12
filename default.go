package validator

import "errors"

type ErrMsg struct {
	Err error
}

func (e ErrMsg) Error() string {
	return e.Err.Error()
}

var (
	ErrNotNull = "notnull validation failed for field %s"

	ErrInvalidValidationForField = "invalid validation applied to the field %s"

	ErrMin = "min value validation failed for field %s"

	ErrMax = "max value validation failed for field %s"

	ErrExclusiveMin = "exclusive min validation failed for field %s"

	ErrExclusiveMax = "exclusive max validation failed for field %s"

	ErrMultipleOf = ErrMsg{errors.New("multipleOf validation failed")}

	ErrMinLength = ErrMsg{errors.New("min-length validation failed")}

	ErrMaxLength = ErrMsg{errors.New("max-length validation failed")}

	ErrPattern = ErrMsg{errors.New("pattern validation failed")}

	ErrBadConstraint = ErrMsg{errors.New("invalid constraint value")}

	ErrNotSupported = ErrMsg{errors.New("unsupported constraint on type")}

	ErrEnums = ErrMsg{errors.New("enum validation failed")}
)
