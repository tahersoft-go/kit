package filter

import (
	"fmt"
	"strings"
)

type OperatorValue struct {
	Op    string
	Value interface{}
}

type FilterValue[T string | int] struct {
	Op    string `json:"op,omitempty"`
	Value T      `json:"value,omitempty"`
}

const (
	AND = "AND"
	OR  = "OR"
)

const (
	EQUALS                     = "equals"
	CONTAINS                   = "contains"
	STARTS_WITH                = "startsWith"
	ENDS_WITH                  = "endsWith"
	IS                         = "is"         // boolean is and date is
	IS_EMPTY                   = "isEmpty"    // works for string, date, number
	IS_NOT_EMPTY               = "isNotEmpty" // works for string, date, number
	IS_ANY_OF                  = "isAnyOf"    // works for string, date, number
	NUMBER_EQUALS              = "="
	NUMBER_NOT_EQUALS          = "!="
	NUMBER_GREATER_THAN        = ">"
	NUMBER_GREATER_THAN_EQUALS = ">="
	NUMBER_LESS_THAN           = "<"
	NUMBER_LESS_THAN_EQUALS    = "<="
	DATE_IS_NOT                = "not"
	DATE_IS_AFTER              = "after"
	DATE_IS_ON_OR_AFTER        = "onOrAfter"
	DATE_IS_BEFORE             = "before"
	DATE_IS_ON_OR_BEFORE       = "onOrBefore"
)

var MapSqlOperators = map[string]string{
	EQUALS:                     "=",
	CONTAINS:                   "ILIKE",
	STARTS_WITH:                "ILIKE",
	ENDS_WITH:                  "ILIKE",
	IS_EMPTY:                   "IS NULL",
	IS_NOT_EMPTY:               "IS NOT NULL",
	IS_ANY_OF:                  "IN",
	NUMBER_EQUALS:              "=",
	NUMBER_NOT_EQUALS:          "!=",
	NUMBER_GREATER_THAN:        ">",
	NUMBER_GREATER_THAN_EQUALS: ">=",
	NUMBER_LESS_THAN:           "<",
	NUMBER_LESS_THAN_EQUALS:    "<=",
	IS:                         "=",
	DATE_IS_NOT:                "!=",
	DATE_IS_AFTER:              ">",
	DATE_IS_BEFORE:             "<",
	DATE_IS_ON_OR_AFTER:        ">=",
	DATE_IS_ON_OR_BEFORE:       "<=",
}

type DBOperatorAndValue struct {
	Operator string
	Value    string
}

// GetDBOperatorAndValue returns the database operator and value for a given operator and value
func GetDBOperatorAndValue(op, value string) DBOperatorAndValue {

	switch op {
	case CONTAINS:
		return DBOperatorAndValue{
			Operator: MapSqlOperators[CONTAINS],
			Value:    "'" + "%" + value + "%" + "'",
		}
	case STARTS_WITH:
		return DBOperatorAndValue{
			Operator: MapSqlOperators[STARTS_WITH],
			Value:    "'" + value + "%" + "'",
		}
	case ENDS_WITH:
		return DBOperatorAndValue{
			Operator: MapSqlOperators[ENDS_WITH],
			Value:    "'" + "%" + value + "'",
		}
	case EQUALS:
		return DBOperatorAndValue{
			Operator: MapSqlOperators[EQUALS],
			Value:    "'" + value + "'",
		}
	case IS_EMPTY:
		return DBOperatorAndValue{
			Operator: MapSqlOperators[IS_EMPTY],
			Value:    "",
		}
	case IS_NOT_EMPTY:
		return DBOperatorAndValue{
			Operator: MapSqlOperators[IS_NOT_EMPTY],
			Value:    "",
		}
	case IS_ANY_OF:
		strSlice := strings.Split(value, ",")
		for i, v := range strSlice {
			strSlice[i] = "'" + v + "'"
		}
		return DBOperatorAndValue{
			Operator: MapSqlOperators[IS_ANY_OF],
			Value:    fmt.Sprintf("(%s)", strings.Join(strSlice, ",")),
		}
	case NUMBER_EQUALS:
		return DBOperatorAndValue{
			Operator: MapSqlOperators[NUMBER_EQUALS],
			Value:    "'" + value + "'",
		}
	case NUMBER_NOT_EQUALS:
		return DBOperatorAndValue{
			Operator: MapSqlOperators[NUMBER_NOT_EQUALS],
			Value:    "'" + value + "'",
		}
	case NUMBER_GREATER_THAN:
		return DBOperatorAndValue{
			Operator: MapSqlOperators[NUMBER_GREATER_THAN],
			Value:    "'" + value + "'",
		}
	case NUMBER_GREATER_THAN_EQUALS:
		return DBOperatorAndValue{
			Operator: MapSqlOperators[NUMBER_GREATER_THAN_EQUALS],
			Value:    "'" + value + "'",
		}
	case NUMBER_LESS_THAN:
		return DBOperatorAndValue{
			Operator: MapSqlOperators[NUMBER_LESS_THAN],
			Value:    "'" + value + "'",
		}
	case NUMBER_LESS_THAN_EQUALS:
		return DBOperatorAndValue{
			Operator: MapSqlOperators[NUMBER_LESS_THAN_EQUALS],
			Value:    "'" + value + "'",
		}
	case IS:
		return DBOperatorAndValue{
			Operator: MapSqlOperators[IS],
			Value:    "'" + value + "'",
		}
	case DATE_IS_NOT:
		return DBOperatorAndValue{
			Operator: MapSqlOperators[DATE_IS_NOT],
			Value:    "'" + value + "'",
		}
	case DATE_IS_AFTER:
		return DBOperatorAndValue{
			Operator: MapSqlOperators[DATE_IS_AFTER],
			Value:    "'" + value + "'",
		}
	case DATE_IS_BEFORE:
		return DBOperatorAndValue{
			Operator: MapSqlOperators[DATE_IS_BEFORE],
			Value:    "'" + value + "'",
		}
	case DATE_IS_ON_OR_AFTER:
		return DBOperatorAndValue{
			Operator: MapSqlOperators[DATE_IS_ON_OR_AFTER],
			Value:    "'" + value + "'",
		}
	case DATE_IS_ON_OR_BEFORE:
		return DBOperatorAndValue{
			Operator: MapSqlOperators[DATE_IS_ON_OR_BEFORE],
			Value:    "'" + value + "'",
		}
	}
	return DBOperatorAndValue{
		Operator: MapSqlOperators[EQUALS],
		Value:    "'" + value + "'",
	}
}
