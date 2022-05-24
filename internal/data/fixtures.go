package data

import (
	"context"
	"stamp/internal/models"

	"github.com/volatiletech/sqlboiler/v4/boil"
)

// Live Service fixtures to be applied by manually running the CLI "app db seed"
// Note that these fixtures are not available while testing
// see the separate internal/test/fixtures.go file

type Upsertable interface {
	Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error
}

type FixtureMap struct {
	AnalyticsStamp *models.Stamp
}

func Fixtures() FixtureMap {
	f := FixtureMap{}

	f.AnalyticsStamp = &models.Stamp{
		ID:   "1c8ceb14-b369-4cc8-9a85-8c0445588ee2",
		Name: "Analytics",
	}

	return f
}

func Upserts() []Upsertable {
	fixtures := Fixtures()

	return []Upsertable{
		fixtures.AnalyticsStamp,
	}
}
