package mysql

import "database/sql"

func NewNullInt64(i int64) *NullInt {
	if i == 0 {
		return &NullInt{}
	}

	return &NullInt{
		NullInt64: sql.NullInt64{
			Int64: i,
			Valid: true,
		},
	}
}
