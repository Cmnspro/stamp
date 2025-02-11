// Code generated by SQLBoiler 4.6.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"bytes"
	"context"
	"reflect"
	"testing"

	"github.com/volatiletech/randomize"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/strmangle"
)

var (
	// Relationships sometimes use the reflection helper queries.Equal/queries.Assign
	// so force a package dependency in case they don't.
	_ = queries.Equal
)

func testDomainStamps(t *testing.T) {
	t.Parallel()

	query := DomainStamps()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testDomainStampsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DomainStamp{}
	if err = randomize.Struct(seed, o, domainStampDBTypes, true, domainStampColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DomainStamp struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := o.Delete(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := DomainStamps().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testDomainStampsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DomainStamp{}
	if err = randomize.Struct(seed, o, domainStampDBTypes, true, domainStampColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DomainStamp struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := DomainStamps().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := DomainStamps().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testDomainStampsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DomainStamp{}
	if err = randomize.Struct(seed, o, domainStampDBTypes, true, domainStampColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DomainStamp struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := DomainStampSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := DomainStamps().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testDomainStampsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DomainStamp{}
	if err = randomize.Struct(seed, o, domainStampDBTypes, true, domainStampColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DomainStamp struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := DomainStampExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if DomainStamp exists: %s", err)
	}
	if !e {
		t.Errorf("Expected DomainStampExists to return true, but got false.")
	}
}

func testDomainStampsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DomainStamp{}
	if err = randomize.Struct(seed, o, domainStampDBTypes, true, domainStampColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DomainStamp struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	domainStampFound, err := FindDomainStamp(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if domainStampFound == nil {
		t.Error("want a record, got nil")
	}
}

func testDomainStampsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DomainStamp{}
	if err = randomize.Struct(seed, o, domainStampDBTypes, true, domainStampColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DomainStamp struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = DomainStamps().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testDomainStampsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DomainStamp{}
	if err = randomize.Struct(seed, o, domainStampDBTypes, true, domainStampColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DomainStamp struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := DomainStamps().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testDomainStampsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	domainStampOne := &DomainStamp{}
	domainStampTwo := &DomainStamp{}
	if err = randomize.Struct(seed, domainStampOne, domainStampDBTypes, false, domainStampColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DomainStamp struct: %s", err)
	}
	if err = randomize.Struct(seed, domainStampTwo, domainStampDBTypes, false, domainStampColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DomainStamp struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = domainStampOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = domainStampTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := DomainStamps().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testDomainStampsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	domainStampOne := &DomainStamp{}
	domainStampTwo := &DomainStamp{}
	if err = randomize.Struct(seed, domainStampOne, domainStampDBTypes, false, domainStampColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DomainStamp struct: %s", err)
	}
	if err = randomize.Struct(seed, domainStampTwo, domainStampDBTypes, false, domainStampColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DomainStamp struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = domainStampOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = domainStampTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := DomainStamps().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func testDomainStampsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DomainStamp{}
	if err = randomize.Struct(seed, o, domainStampDBTypes, true, domainStampColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DomainStamp struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := DomainStamps().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testDomainStampsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DomainStamp{}
	if err = randomize.Struct(seed, o, domainStampDBTypes, true); err != nil {
		t.Errorf("Unable to randomize DomainStamp struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(domainStampColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := DomainStamps().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testDomainStampToManyVotes(t *testing.T) {
	var err error
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a DomainStamp
	var b, c Vote

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, domainStampDBTypes, true, domainStampColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DomainStamp struct: %s", err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = randomize.Struct(seed, &b, voteDBTypes, false, voteColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, voteDBTypes, false, voteColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}

	b.DomainStampID = a.ID
	c.DomainStampID = a.ID

	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := a.Votes().All(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range check {
		if v.DomainStampID == b.DomainStampID {
			bFound = true
		}
		if v.DomainStampID == c.DomainStampID {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := DomainStampSlice{&a}
	if err = a.L.LoadVotes(ctx, tx, false, (*[]*DomainStamp)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.Votes); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.Votes = nil
	if err = a.L.LoadVotes(ctx, tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.Votes); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", check)
	}
}

func testDomainStampToManyAddOpVotes(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a DomainStamp
	var b, c, d, e Vote

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, domainStampDBTypes, false, strmangle.SetComplement(domainStampPrimaryKeyColumns, domainStampColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Vote{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, voteDBTypes, false, strmangle.SetComplement(votePrimaryKeyColumns, voteColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	foreignersSplitByInsertion := [][]*Vote{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddVotes(ctx, tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.ID != first.DomainStampID {
			t.Error("foreign key was wrong value", a.ID, first.DomainStampID)
		}
		if a.ID != second.DomainStampID {
			t.Error("foreign key was wrong value", a.ID, second.DomainStampID)
		}

		if first.R.DomainStamp != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.DomainStamp != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.Votes[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.Votes[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.Votes().Count(ctx, tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}
func testDomainStampToOneDomainUsingDomain(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local DomainStamp
	var foreign Domain

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, domainStampDBTypes, false, domainStampColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DomainStamp struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, domainDBTypes, false, domainColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Domain struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	local.DomainID = foreign.ID
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.Domain().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.ID != foreign.ID {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := DomainStampSlice{&local}
	if err = local.L.LoadDomain(ctx, tx, false, (*[]*DomainStamp)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Domain == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Domain = nil
	if err = local.L.LoadDomain(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Domain == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testDomainStampToOneStampUsingStamp(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local DomainStamp
	var foreign Stamp

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, domainStampDBTypes, false, domainStampColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DomainStamp struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, stampDBTypes, false, stampColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Stamp struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	local.StampID = foreign.ID
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.Stamp().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.ID != foreign.ID {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := DomainStampSlice{&local}
	if err = local.L.LoadStamp(ctx, tx, false, (*[]*DomainStamp)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Stamp == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Stamp = nil
	if err = local.L.LoadStamp(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Stamp == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testDomainStampToOneSetOpDomainUsingDomain(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a DomainStamp
	var b, c Domain

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, domainStampDBTypes, false, strmangle.SetComplement(domainStampPrimaryKeyColumns, domainStampColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, domainDBTypes, false, strmangle.SetComplement(domainPrimaryKeyColumns, domainColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, domainDBTypes, false, strmangle.SetComplement(domainPrimaryKeyColumns, domainColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Domain{&b, &c} {
		err = a.SetDomain(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Domain != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.DomainStamps[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.DomainID != x.ID {
			t.Error("foreign key was wrong value", a.DomainID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.DomainID))
		reflect.Indirect(reflect.ValueOf(&a.DomainID)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.DomainID != x.ID {
			t.Error("foreign key was wrong value", a.DomainID, x.ID)
		}
	}
}
func testDomainStampToOneSetOpStampUsingStamp(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a DomainStamp
	var b, c Stamp

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, domainStampDBTypes, false, strmangle.SetComplement(domainStampPrimaryKeyColumns, domainStampColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, stampDBTypes, false, strmangle.SetComplement(stampPrimaryKeyColumns, stampColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, stampDBTypes, false, strmangle.SetComplement(stampPrimaryKeyColumns, stampColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Stamp{&b, &c} {
		err = a.SetStamp(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Stamp != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.DomainStamps[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.StampID != x.ID {
			t.Error("foreign key was wrong value", a.StampID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.StampID))
		reflect.Indirect(reflect.ValueOf(&a.StampID)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.StampID != x.ID {
			t.Error("foreign key was wrong value", a.StampID, x.ID)
		}
	}
}

func testDomainStampsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DomainStamp{}
	if err = randomize.Struct(seed, o, domainStampDBTypes, true, domainStampColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DomainStamp struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = o.Reload(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testDomainStampsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DomainStamp{}
	if err = randomize.Struct(seed, o, domainStampDBTypes, true, domainStampColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DomainStamp struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := DomainStampSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testDomainStampsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DomainStamp{}
	if err = randomize.Struct(seed, o, domainStampDBTypes, true, domainStampColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DomainStamp struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := DomainStamps().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	domainStampDBTypes = map[string]string{`ID`: `uuid`, `DomainID`: `uuid`, `StampID`: `uuid`, `CreatedAt`: `timestamp with time zone`, `UpdatedAt`: `timestamp with time zone`}
	_                  = bytes.MinRead
)

func testDomainStampsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(domainStampPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(domainStampAllColumns) == len(domainStampPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &DomainStamp{}
	if err = randomize.Struct(seed, o, domainStampDBTypes, true, domainStampColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DomainStamp struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := DomainStamps().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, domainStampDBTypes, true, domainStampPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize DomainStamp struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testDomainStampsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(domainStampAllColumns) == len(domainStampPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &DomainStamp{}
	if err = randomize.Struct(seed, o, domainStampDBTypes, true, domainStampColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DomainStamp struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := DomainStamps().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, domainStampDBTypes, true, domainStampPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize DomainStamp struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(domainStampAllColumns, domainStampPrimaryKeyColumns) {
		fields = domainStampAllColumns
	} else {
		fields = strmangle.SetComplement(
			domainStampAllColumns,
			domainStampPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	typ := reflect.TypeOf(o).Elem()
	n := typ.NumField()

	updateMap := M{}
	for _, col := range fields {
		for i := 0; i < n; i++ {
			f := typ.Field(i)
			if f.Tag.Get("boil") == col {
				updateMap[col] = value.Field(i).Interface()
			}
		}
	}

	slice := DomainStampSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testDomainStampsUpsert(t *testing.T) {
	t.Parallel()

	if len(domainStampAllColumns) == len(domainStampPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := DomainStamp{}
	if err = randomize.Struct(seed, &o, domainStampDBTypes, true); err != nil {
		t.Errorf("Unable to randomize DomainStamp struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert DomainStamp: %s", err)
	}

	count, err := DomainStamps().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, domainStampDBTypes, false, domainStampPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize DomainStamp struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert DomainStamp: %s", err)
	}

	count, err = DomainStamps().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
