// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package cart

import (
	"database/sql/driver"
	"fmt"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/shopspring/decimal"
)

type CartStatus string

const (
	CartStatusOpen      CartStatus = "Open"
	CartStatusCompleted CartStatus = "Completed"
	CartStatusMerged    CartStatus = "Merged"
	CartStatusExpired   CartStatus = "Expired"
	CartStatusAbandoned CartStatus = "Abandoned"
)

func (e *CartStatus) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = CartStatus(s)
	case string:
		*e = CartStatus(s)
	default:
		return fmt.Errorf("unsupported scan type for CartStatus: %T", src)
	}
	return nil
}

type NullCartStatus struct {
	CartStatus CartStatus
	Valid      bool // Valid is true if CartStatus is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullCartStatus) Scan(value interface{}) error {
	if value == nil {
		ns.CartStatus, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.CartStatus.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullCartStatus) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.CartStatus), nil
}

type Cart struct {
	ID           int64
	UserID       *int64
	Status       CartStatus
	LastActivity pgtype.Timestamptz
	ExpiresAt    pgtype.Timestamptz
	CreatedAt    pgtype.Timestamptz
	UpdatedAt    pgtype.Timestamptz
}

type CartItem struct {
	ID                int64
	CartID            int64
	VariantID         int64
	Qty               int32
	Price             decimal.Decimal
	Currency          string
	Thumbnail         string
	ProductName       string
	ProductPath       string
	VariantAttributes []byte
	IsRemoved         bool
	CreatedAt         pgtype.Timestamptz
	UpdatedAt         pgtype.Timestamptz
}
