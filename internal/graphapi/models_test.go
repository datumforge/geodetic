package graphapi_test

import (
	"context"
	"testing"

	"github.com/brianvoe/gofakeit/v7"

	ent "github.com/datumforge/geodetic/internal/ent/generated"
)

type GroupBuilder struct {
	client *client

	// Fields
	Name     string
	Location string
}

type GroupCleanup struct {
	client *client

	// Fields
	GroupID string
}

// MustNew group builder is used to create groups in the database
func (g *GroupBuilder) MustNew(ctx context.Context, t *testing.T) *ent.Group {
	if g.Name == "" {
		g.Name = gofakeit.AppName()
	}

	if g.Location == "" {
		g.Location = "den"
	}

	group := g.client.db.Group.Create().
		SetName(g.Name).
		SetPrimaryLocation(g.Location).
		SaveX(ctx)

	return group
}

// MustDelete is used to cleanup groups in the database
func (g *GroupCleanup) MustDelete(ctx context.Context, t *testing.T) {
	g.client.db.Group.DeleteOneID(g.GroupID).ExecX(ctx)
}
