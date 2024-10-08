// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: address.sql

package address

import (
	"context"
)

const createAddress = `-- name: CreateAddress :one
INSERT INTO addresses (
    user_id, default_for, name, address_line_1, address_line_2, city, state_province, postal_code, country
) VALUES  (
  $1, $2, $3, $4, $5,$6,$7,$8, $9
) RETURNING id, user_id, default_for, name, address_line_1, address_line_2, city, state_province, postal_code, country, phone, created_at, updated_at
`

func (q *Queries) CreateAddress(ctx context.Context, userID int64, defaultFor NullAddressDefaultType, name string, addressLine1 string, addressLine2 *string, city string, stateProvince *string, postalCode *string, country *string) (*Address, error) {
	row := q.db.QueryRow(ctx, createAddress,
		userID,
		defaultFor,
		name,
		addressLine1,
		addressLine2,
		city,
		stateProvince,
		postalCode,
		country,
	)
	var i Address
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.DefaultFor,
		&i.Name,
		&i.AddressLine1,
		&i.AddressLine2,
		&i.City,
		&i.StateProvince,
		&i.PostalCode,
		&i.Country,
		&i.Phone,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return &i, err
}

const getAddressByID = `-- name: GetAddressByID :one
SELECT id, user_id, default_for, name, address_line_1, address_line_2, city, state_province, postal_code, country, phone, created_at, updated_at FROM addresses
WHERE id = $1
`

func (q *Queries) GetAddressByID(ctx context.Context, id int64) (*Address, error) {
	row := q.db.QueryRow(ctx, getAddressByID, id)
	var i Address
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.DefaultFor,
		&i.Name,
		&i.AddressLine1,
		&i.AddressLine2,
		&i.City,
		&i.StateProvince,
		&i.PostalCode,
		&i.Country,
		&i.Phone,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return &i, err
}

const getUserDefaultAddress = `-- name: GetUserDefaultAddress :one
SELECT id, user_id, default_for, name, address_line_1, address_line_2, city, state_province, postal_code, country, phone, created_at, updated_at FROM addresses
WHERE user_id = $1
    AND default_for = $2
ORDER BY created_at
OFFSET 1
`

func (q *Queries) GetUserDefaultAddress(ctx context.Context, userID int64, defaultFor NullAddressDefaultType) (*Address, error) {
	row := q.db.QueryRow(ctx, getUserDefaultAddress, userID, defaultFor)
	var i Address
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.DefaultFor,
		&i.Name,
		&i.AddressLine1,
		&i.AddressLine2,
		&i.City,
		&i.StateProvince,
		&i.PostalCode,
		&i.Country,
		&i.Phone,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return &i, err
}

const listAddressesByUserID = `-- name: ListAddressesByUserID :many
SELECT id, user_id, default_for, name, address_line_1, address_line_2, city, state_province, postal_code, country, phone, created_at, updated_at FROM addresses
WHERE user_id = $1
ORDER BY created_at
OFFSET $2
    LIMIT $3
`

func (q *Queries) ListAddressesByUserID(ctx context.Context, userID int64, offset int32, limit int32) ([]*Address, error) {
	rows, err := q.db.Query(ctx, listAddressesByUserID, userID, offset, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []*Address
	for rows.Next() {
		var i Address
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.DefaultFor,
			&i.Name,
			&i.AddressLine1,
			&i.AddressLine2,
			&i.City,
			&i.StateProvince,
			&i.PostalCode,
			&i.Country,
			&i.Phone,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, &i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
