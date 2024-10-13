package common

import (
	"database/sql"
	"time"
)

func ConvertNullString(ns sql.NullString) string {
	if ns.Valid {
		return ns.String
	}
	return ""
}

func ConvertNullTime(nt sql.NullTime) time.Time {
	if nt.Valid {
		return nt.Time
	}
	return time.Time{}
}

func ConvertNullInt32(ni sql.NullInt32) int32 {
	if ni.Valid {
		return ni.Int32
	}
	return 0
}
