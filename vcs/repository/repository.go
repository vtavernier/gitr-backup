package repository

import (
	"context"
)

type Ref struct {
	Name    string `diff:"name, identifier"`
	Sha     string `diff:"sha"`
	RefName string `diff:"-"`
}

type Repository interface {
	GetName() string
	GetDescription() string
	AddLabel(ctx context.Context, label string) error
	RemoveLabel(ctx context.Context, label string) error
	ListRefs(ctx context.Context) ([]Ref, error)
	GetHttpsCloneUrl() (string, error)
	GetUrl() string
	GetDefaultBranch() string
	SetDefaultBranch(ctx context.Context, branch string) error
}
