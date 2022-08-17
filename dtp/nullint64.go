package dtp

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
)

type NullInt64 sql.NullInt64

func (n *NullInt64) Scan(value interface{}) error {
	return (*sql.NullInt64)(n).Scan(value)
}

// Value implements the driver Valuer interface.
func (n NullInt64) Value() (driver.Value, error) {
	if !n.Valid {
		return nil, nil
	}
	return n.Int64, nil
}

func (n NullInt64) MarshalJSON() ([]byte, error) {
	if n.Valid {
		return json.Marshal(n.Int64)
	}
	return json.Marshal(nil)
}

func (n *NullInt64) UnmarshalJSON(b []byte) error {
	if string(b) == NullValue {
		n.Valid = false
		return nil
	}
	err := json.Unmarshal(b, &n.Int64)
	if err == nil {
		n.Valid = true
	}
	return err
}

