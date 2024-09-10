package val

import "reflect"

type Stringer interface {
	String() string
}

func coerceToString(value any) string {
	s, ok := value.(Stringer)

	if ok {
		return s.String()
	}

	switch t := reflect.TypeOf(value); t.Kind() {
	case reflect.Bool:
		v := value.(bool)
		if v {
			return "true"
		}
		return "false"
	default:
		return ""
	}

	// Bool
	// Int
	// Int8
	// Int16
	// Int32
	// Int64
	// Uint
	// Uint8
	// Uint16
	// Uint32
	// Uint64
	// Uintptr
	// Float32
	// Float64
	// Complex64
	// Complex128
	// Array
	// Chan
	// Func
	// Interface
	// Map
	// Pointer
	// Slice
	// String
	// Struct
}
