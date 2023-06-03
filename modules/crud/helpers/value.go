package helpers

func IsInteger(value any) bool {
	switch value.(type) {
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
		return true
	}
	return false
}

func IsString(value any) bool {
	switch value.(type) {
	case string:
		return true
	}
	return false
}

func IsFloat(value any) bool {
	switch value.(type) {
	case float32:
	case float64:
		return true
	}
	return false
}
