package convert

import (
	"fmt"
	"strconv"
)

//nolint
var errConvertionError = func(v interface{}) error {
	return fmt.Errorf("cannot convert value %v (type %T) to integer", v, v)
}

// ToInteger converts given value to integer.
func ToInteger(v interface{}) (int, error) {
	i, err := strconv.Atoi(v.(string))
	if err != nil {
		return i, errConvertionError(v)
	}

	return i, nil
}
