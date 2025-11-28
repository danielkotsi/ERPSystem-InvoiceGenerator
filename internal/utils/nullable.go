package utils

import (
	"database/sql"
	"fmt"
)

func NullableString(ns sql.NullString) *string {
	if ns.Valid {
		return &ns.String
	}
	return nil
}

func CheckIfSomethingNotNull(nullables ...sql.NullString) bool {
	for _, string := range nullables {
		fmt.Println("hello this is the string", string.String)
		if string.Valid {
			return true
		}
	}
	return false
}
