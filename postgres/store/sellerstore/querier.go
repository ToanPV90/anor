// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package sellerstore

import (
	"context"
)

type Querier interface {
	CreateStore(ctx context.Context, publicID string, sellerID int64, name string, description string) (int32, error)
}

var _ Querier = (*Queries)(nil)
