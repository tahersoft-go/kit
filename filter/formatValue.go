package filter

import (
	"fmt"
	"strconv"
)

func FormatValue(value string) string {
	if _, err := strconv.ParseFloat(value, 64); err == nil {
		// Value is a number, do not add quotes
		return value
	}
	// Value is not a number, add quotes
	return fmt.Sprintf("'%s'", value)
}
