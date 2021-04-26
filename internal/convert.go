package convert

import (
	"fmt"
)

var (
	errConvertionError = func(v interface{}) error {
		return fmt.Errorf("cannot convert %v to integer", v)
	}
)

// ConvertToInteger converts given value to integer
func ConvertToInteger(v interface{}) (int, error) {
	i, ok := v.(int)
	if !ok {
		return i, errConvertionError(v)
	}

	return i, nil
}
