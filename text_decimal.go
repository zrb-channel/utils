package utils

import (
	"fmt"
	"strings"

	"github.com/shopspring/decimal"
)

type TextDecimal struct {
	decimal.Decimal
}

// UnmarshalJSON
// @param data
// @date 2022-09-21 18:00:57
func (d *TextDecimal) UnmarshalJSON(data []byte) error {

	if string(data) == "null" {
		return nil
	}

	str, err := unquoteIfQuoted(data)
	if err != nil {
		return fmt.Errorf("error decoding string '%s': %s", data, err)
	}
	str = strings.ReplaceAll(str, ",", "")
	if str == `""` || str == `— —` {
		str = "0"
	}

	value, err := decimal.NewFromString(str)
	*d = TextDecimal{value}
	if err != nil {
		return fmt.Errorf("error decoding string '%s': %s", str, err)
	}

	return nil
}

// unquoteIfQuoted
// @param value
// @date 2022-09-21 18:00:55
func unquoteIfQuoted(value interface{}) (string, error) {
	var bytes []byte

	switch v := value.(type) {
	case string:
		bytes = []byte(v)
	case []byte:
		bytes = v
	default:
		return "", fmt.Errorf("could not convert value '%+v' to byte array of type '%T'",
			value, value)
	}

	// If the amount is quoted, strip the quotes
	if len(bytes) > 2 && bytes[0] == '"' && bytes[len(bytes)-1] == '"' {
		bytes = bytes[1 : len(bytes)-1]
	}
	return string(bytes), nil
}
