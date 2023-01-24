package mysql

import (
	"database/sql"
	"encoding/json"
)

type NullInt struct {
	sql.NullInt64
}

func (v *NullInt) MarshalJSON() ([]byte, error) {
	if !v.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(v.Int64)
}
