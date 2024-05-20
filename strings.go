package val

import "fmt"

type stringsValidator = func(string) (string, error)

type strings struct {
	field      string
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
