package hemnetparser

// PtrFloat64 get pointer of value float64
func PtrFloat64(val float64) *float64 {
	return &val
}

// GetFloat64 get value of pointer or empty value
func GetFloat64(val *float64) float64 {
	if val == nil {
		return 0.0
	} else {
		return *val
	}
}

// PtrString get pointer of value string
func PtrString(val string) *string {
	return &val
}

// GetString get value of pointer or empty value
func GetString(val *string) string {
	if val == nil {
		return ""
	} else {
		return *val
	}
}

// StringValueOrNil convert empty if it's nil
func StringValueOrNil(val string) *string {
	if len(val) == 0 {
		return nil
	}
	return &val
}
