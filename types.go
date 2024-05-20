package val

type valError struct {
	field string
	msg   string
	err   error
}

func (ve valError) Field() string {
	return ve.field
}

func (ve valError) Error() string {
	return ve.msg
}

func (ve valError) Unwrap() error {
	return ve.err
}
