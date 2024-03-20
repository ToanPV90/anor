// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package category

import (
	"context"
)

type Querier interface {
	CreateChildCategoryIfNotExists(ctx context.Context, category string, slug string, parentID *int32) (*CreateChildCategoryIfNotExistsRow, error)
	CreateTopCategory(ctx context.Context, category string, slug string) (*Category, error)
	GetCategory(ctx context.Context, id int32) (*GetCategoryRow, error)
	GetCategoryByName(ctx context.Context, category string) (*Category, error)
	GetCategoryWithAncestors(ctx context.Context, id int32) ([]*GetCategoryWithAncestorsRow, error)
	GetCategoryWithSiblings(ctx context.Context, id int32) ([]*Category, error)
	GetChildCategories(ctx context.Context, parentID *int32) ([]*Category, error)
	GetLeafCategoryIDs(ctx context.Context, id int32) ([][]int32, error)
}

var _ Querier = (*Queries)(nil)
