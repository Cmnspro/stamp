package db_test

import (
	"testing"

	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"stamp/internal/models"
	"stamp/internal/test"
	"stamp/internal/util/db"
)

func TestILike(t *testing.T) {
	q := models.NewQuery(
		qm.Select("*"),
		qm.From("users"),
		db.InnerJoin("users", "id", "app_user_profiles", "user_id"),
		db.ILike("%Max.Muster%", "users", "username"),
		db.ILike("Max", "users", "app_user_profiles", "first_name"),
	)

	sql, args := queries.BuildQuery(q)

	test.Snapshoter.Label("SQL").Save(t, sql)
	test.Snapshoter.Label("Args").Save(t, args)
}
