package utils

import (
	"database/sql"
)

func NullableString(ns sql.NullString) *string {
	if ns.Valid {
		return &ns.String
	}
	return nil
}

func CheckIfSomethingNotNull(nullables ...sql.NullString) bool {
	for _, string := range nullables {
		if string.Valid {
			return true
		}
	}
	return false
}
