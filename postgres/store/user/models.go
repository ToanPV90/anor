// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package user

import (
	"database/sql/driver"
	"fmt"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/shopspring/decimal"
)

type ProductStatus string

const (
	ProductStatusDraft           ProductStatus = "Draft"
	ProductStatusPendingApproval ProductStatus = "PendingApproval"
	ProductStatusActive          ProductStatus = "Active"
)

func (e *ProductStatus) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = ProductStatus(s)
	case string:
		*e = ProductStatus(s)
	default:
		return fmt.Errorf("unsupported scan type for ProductStatus: %T", src)
	}
	return nil
}

type NullProductStatus struct {
	ProductStatus ProductStatus
	Valid         bool // Valid is true if ProductStatus is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullProductStatus) Scan(value interface{}) error {
	if value == nil {
		ns.ProductStatus, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.ProductStatus.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullProductStatus) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.ProductStatus), nil
}

type UserStatus string

const (
	UserStatusBlocked             UserStatus = "Blocked"
	UserStatusRegistrationPending UserStatus = "RegistrationPending"
	UserStatusActive              UserStatus = "Active"
	UserStatusInactice            UserStatus = "Inactice"
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

type Category struct {
	ID       int32
	Category string
	Slug     string
	ParentID *int32
}

type Product struct {
	ID         int64
	StoreID    int32
	CategoryID int32
	Name       string
	Brand      *string
	Slug       string
	ShortInfo  []string
	ImageUrls  []byte
	Specs      []byte
	Status     ProductStatus
	CreatedAt  pgtype.Timestamptz
	UpdatedAt  pgtype.Timestamptz
}

type ProductAttribute struct {
	ID        int64
	ProductID int64
	Attribute string
}

type ProductPricing struct {
	ProductID        int64
	BasePrice        decimal.Decimal
	CurrencyCode     string
	DiscountLevel    decimal.Decimal
	DiscountedAmount decimal.Decimal
	IsOnSale         bool
}

type ProductRating struct {
	ID        int64
	ProductID int64
	UserID    *int64
	Rating    int16
	Review    *string
	ImageUrls []string
	CreatedAt pgtype.Timestamptz
}

type Role struct {
	ID   int16
	Role string
}

type SellerStore struct {
	ID          int32
	PublicID    string
	SellerID    int64
	Name        string
	Description string
	CreatedAt   pgtype.Timestamptz
	UpdatedAt   pgtype.Timestamptz
}

type Session struct {
	Token  string
	Data   []byte
	Expiry pgtype.Timestamptz
}

type Sku struct {
	ID              int64
	ProductID       int64
	Sku             string
	QuantityInStock int32
	HasSepPricing   bool
	ImageRefs       []int16
	CreatedAt       pgtype.Timestamptz
	UpdatedAt       pgtype.Timestamptz
}

type SkuPricing struct {
	SkuID            int64
	BasePrice        decimal.Decimal
	CurrencyCode     string
	DiscountLevel    decimal.Decimal
	DiscountedAmount decimal.Decimal
	IsOnSale         bool
}

type SkuProductAttributeValue struct {
	SkuID              int64
	ProductAttributeID int64
	AttributeValue     string
}

type User struct {
	ID          int64
	Email       string
	Password    string
	PhoneNumber *string
	FullName    string
	Status      UserStatus
	Otp         *string
	OtpExpiry   *int64
	CreatedAt   pgtype.Timestamptz
	UpdatedAt   pgtype.Timestamptz
}

type UserRole struct {
	UserID int64
	RoleID int16
}
