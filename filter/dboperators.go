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
	IS_EMPTY                   = "isEmpty"
	IS_NOT_EMPTY               = "isNotEmpty"
	IS_ANY_OF                  = "isAnyOf"
	NUMBER_EQUALS              = "="
	NUMBER_NOT_EQUALS          = "!="
	NUMBER_GREATER_THAN        = ">"
	NUMBER_GREATER_THAN_EQUALS = ">="
	NUMBER_LESS_THAN           = "<"
	NUMBER_LESS_THAN_EQUALS    = "<="
)

var SQL = map[string]string{
	EQUALS:                     "=",
	CONTAINS:                   "LIKE",
	STARTS_WITH:                "LIKE",
	ENDS_WITH:                  "LIKE",
	IS_EMPTY:                   "IS NULL",
	IS_NOT_EMPTY:               "IS NOT NULL",
	IS_ANY_OF:                  "IN",
	NUMBER_EQUALS:              "=",
	NUMBER_NOT_EQUALS:          "!=",
	NUMBER_GREATER_THAN:        ">",
	NUMBER_GREATER_THAN_EQUALS: ">=",
	NUMBER_LESS_THAN:           "<",
	NUMBER_LESS_THAN_EQUALS:    "<=",
}

type DBOperatorAndValue struct {
	Operator string
	Value    string
}

// GetDBOperatorAndValue returns the database operator and value for a given operator and value
func GetDBOperatorAndValue(op, value string) DBOperatorAndValue {
	caseSensitive := false

	likeOperator := "LIKE"
	equalsOperator := "="

	if !caseSensitive {
		// Use ILIKE for PostgreSQL or use LOWER for other databases
		likeOperator = "ILIKE"   // or use LOWER(column) LIKE LOWER(value) in the calling function
		equalsOperator = "ILIKE" // or use LOWER(column) = LOWER(value) in the calling function
	}

	switch op {
	case CONTAINS:
		return DBOperatorAndValue{
			Operator: likeOperator,
			Value:    "%" + value + "%",
		}
	case STARTS_WITH:
		return DBOperatorAndValue{
			Operator: likeOperator,
			Value:    value + "%",
		}
	case ENDS_WITH:
		return DBOperatorAndValue{
			Operator: likeOperator,
			Value:    "%" + value,
		}
	case EQUALS:
		return DBOperatorAndValue{
			Operator: equalsOperator,
			Value:    value,
		}
	case IS_EMPTY:
		return DBOperatorAndValue{
			Operator: "IS NULL",
			Value:    "",
		}
	case IS_NOT_EMPTY:
		return DBOperatorAndValue{
			Operator: "IS NOT NULL",
			Value:    "",
		}
	case IS_ANY_OF:
		strSlice := strings.Split(value, ",")
		for i, v := range strSlice {
			strSlice[i] = "'" + v + "'"
		}
		return DBOperatorAndValue{
			Operator: "IN",
			Value:    fmt.Sprintf("(%s)", strings.Join(strSlice, ",")),
		}
	case NUMBER_EQUALS:
		return DBOperatorAndValue{
			Operator: "=",
			Value:    value,
		}
	case NUMBER_NOT_EQUALS:
		return DBOperatorAndValue{
			Operator: "!=",
			Value:    value,
		}
	case NUMBER_GREATER_THAN:
		return DBOperatorAndValue{
			Operator: ">",
			Value:    value,
		}
	case NUMBER_GREATER_THAN_EQUALS:
		return DBOperatorAndValue{
			Operator: ">=",
			Value:    value,
		}
	case NUMBER_LESS_THAN:
		return DBOperatorAndValue{
			Operator: "<",
			Value:    value,
		}
	case NUMBER_LESS_THAN_EQUALS:
		return DBOperatorAndValue{
			Operator: "<=",
			Value:    value,
		}

	}
	return DBOperatorAndValue{
		Operator: "=",
		Value:    value,
	}
}
