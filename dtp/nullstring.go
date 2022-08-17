package dtp

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
)

type NullString sql.NullString

func (n *NullString) Scan(value interface{}) error {
	return (*sql.NullString)(n).Scan(value)
}

// Value implements the driver Valuer interface.
func (n NullString) Value() (driver.Value, error) {
	if !n.Valid {
		return nil, nil
	}
	return n.String, nil
}

func (n NullString) MarshalJSON() ([]byte, error) {
	if n.Valid {
		return json.Marshal(n.String)
	}
	return json.Marshal(nil)
}

func (n *NullString) UnmarshalJSON(b []byte) error {
	if string(b) == NullValue {
		n.Valid = false
		return nil
	}
	err := json.Unmarshal(b, &n.String)
	if err == nil {
		n.Valid = true
	}
	return err
}
