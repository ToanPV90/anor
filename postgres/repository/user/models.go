// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package user

import (
	"database/sql/driver"
	"fmt"

	"github.com/jackc/pgx/v5/pgtype"
)

type UserStatus string

const (
	UserStatusBlocked             UserStatus = "Blocked"
	UserStatusRegistrationPending UserStatus = "RegistrationPending"
	UserStatusActive              UserStatus = "Active"
)

func (e *UserStatus) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = UserStatus(s)
	case string:
		*e = UserStatus(s)
	default:
		return fmt.Errorf("unsupported scan type for UserStatus: %T", src)
	}
	return nil
}

type NullUserStatus struct {
	UserStatus UserStatus
	Valid      bool // Valid is true if UserStatus is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullUserStatus) Scan(value interface{}) error {
	if value == nil {
		ns.UserStatus, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.UserStatus.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullUserStatus) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.UserStatus), nil
}

type User struct {
	ID          int64
	Email       string
	Password    string
	PhoneNumber *string
	FullName    string
	Status      UserStatus
	CreatedAt   pgtype.Timestamptz
	UpdatedAt   pgtype.Timestamptz
}
