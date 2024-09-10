package val

import "fmt"

type stringsValidator = func(string) (string, error)

type strings struct {
	field      string
	coerce     bool
	validators []stringsValidator
}

func Strings() strings {
	return strings{}
}

func (s strings) MinLen(min int) strings {
	validator := func(val string) (string, error) {
		l := len(val)
		if l < min {
			msg := fmt.Sprintf("This field requires at least %d characters", min)
			return val, &valError{field: s.field, msg: msg}
		}
		return val, nil
	}

	s.validators = append(s.validators, validator)

	return s
}

func (s strings) MaxLen(max int) strings {
	validator := func(val string) (string, error) {
		l := len(val)
		if l < max {
			msg := fmt.Sprintf("This field requires maximum %d characters", max)
			return val, &valError{field: s.field, msg: msg}
		}
		return val, nil
	}

	s.validators = append(s.validators, validator)

	return s
}

func (s strings) Len(min, max int) strings {
	if min != max {
		return s.MinLen(min).MaxLen(max)
	}

	validator := func(val string) (string, error) {
		l := len(val)
		if l != min {
			msg := fmt.Sprintf("This field requires %d characters", min)
			return val, &valError{field: s.field, msg: msg}
		}
		return val, nil
	}

	s.validators = append(s.validators, validator)

	return s
}

func (s *strings) Val(value string) (string, error) {
	var err error

	for _, val := range s.validators {
		value, err = val(value)

		if err != nil {
			return value, err
		}
	}

	return value, nil
}

func (s strings) CoerceAndVal(value any) (string, error) {

	return s.Val(value)
}

func (s strings) ValAny(value string) (string, []error) {
	var err error

	errs := make([]error, 0)

	for _, val := range s.validators {
		value, err = val(value)

		errs = append(errs, err)
	}

	return value, errs
}
