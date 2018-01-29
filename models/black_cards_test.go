// Code generated by SQLBoiler (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/randomize"
	"github.com/volatiletech/sqlboiler/strmangle"
)

func testBlackCards(t *testing.T) {
	t.Parallel()

	query := BlackCards(nil)

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}
func testBlackCardsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	blackCard := &BlackCard{}
	if err = randomize.Struct(seed, blackCard, blackCardDBTypes, true, blackCardColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BlackCard struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = blackCard.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = blackCard.Delete(tx); err != nil {
		t.Error(err)
	}

	count, err := BlackCards(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testBlackCardsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	blackCard := &BlackCard{}
	if err = randomize.Struct(seed, blackCard, blackCardDBTypes, true, blackCardColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BlackCard struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = blackCard.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = BlackCards(tx).DeleteAll(); err != nil {
		t.Error(err)
	}

	count, err := BlackCards(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testBlackCardsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	blackCard := &BlackCard{}
	if err = randomize.Struct(seed, blackCard, blackCardDBTypes, true, blackCardColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BlackCard struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = blackCard.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := BlackCardSlice{blackCard}

	if err = slice.DeleteAll(tx); err != nil {
		t.Error(err)
	}

	count, err := BlackCards(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}
func testBlackCardsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	blackCard := &BlackCard{}
	if err = randomize.Struct(seed, blackCard, blackCardDBTypes, true, blackCardColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BlackCard struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = blackCard.Insert(tx); err != nil {
		t.Error(err)
	}

	e, err := BlackCardExists(tx, blackCard.ID)
	if err != nil {
		t.Errorf("Unable to check if BlackCard exists: %s", err)
	}
	if !e {
		t.Errorf("Expected BlackCardExistsG to return true, but got false.")
	}
}
func testBlackCardsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	blackCard := &BlackCard{}
	if err = randomize.Struct(seed, blackCard, blackCardDBTypes, true, blackCardColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BlackCard struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = blackCard.Insert(tx); err != nil {
		t.Error(err)
	}

	blackCardFound, err := FindBlackCard(tx, blackCard.ID)
	if err != nil {
		t.Error(err)
	}

	if blackCardFound == nil {
		t.Error("want a record, got nil")
	}
}
func testBlackCardsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	blackCard := &BlackCard{}
	if err = randomize.Struct(seed, blackCard, blackCardDBTypes, true, blackCardColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BlackCard struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = blackCard.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = BlackCards(tx).Bind(blackCard); err != nil {
		t.Error(err)
	}
}

func testBlackCardsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	blackCard := &BlackCard{}
	if err = randomize.Struct(seed, blackCard, blackCardDBTypes, true, blackCardColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BlackCard struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = blackCard.Insert(tx); err != nil {
		t.Error(err)
	}

	if x, err := BlackCards(tx).One(); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testBlackCardsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	blackCardOne := &BlackCard{}
	blackCardTwo := &BlackCard{}
	if err = randomize.Struct(seed, blackCardOne, blackCardDBTypes, false, blackCardColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BlackCard struct: %s", err)
	}
	if err = randomize.Struct(seed, blackCardTwo, blackCardDBTypes, false, blackCardColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BlackCard struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = blackCardOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = blackCardTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := BlackCards(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testBlackCardsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	blackCardOne := &BlackCard{}
	blackCardTwo := &BlackCard{}
	if err = randomize.Struct(seed, blackCardOne, blackCardDBTypes, false, blackCardColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BlackCard struct: %s", err)
	}
	if err = randomize.Struct(seed, blackCardTwo, blackCardDBTypes, false, blackCardColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BlackCard struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = blackCardOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = blackCardTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := BlackCards(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}
func blackCardBeforeInsertHook(e boil.Executor, o *BlackCard) error {
	*o = BlackCard{}
	return nil
}

func blackCardAfterInsertHook(e boil.Executor, o *BlackCard) error {
	*o = BlackCard{}
	return nil
}

func blackCardAfterSelectHook(e boil.Executor, o *BlackCard) error {
	*o = BlackCard{}
	return nil
}

func blackCardBeforeUpdateHook(e boil.Executor, o *BlackCard) error {
	*o = BlackCard{}
	return nil
}

func blackCardAfterUpdateHook(e boil.Executor, o *BlackCard) error {
	*o = BlackCard{}
	return nil
}

func blackCardBeforeDeleteHook(e boil.Executor, o *BlackCard) error {
	*o = BlackCard{}
	return nil
}

func blackCardAfterDeleteHook(e boil.Executor, o *BlackCard) error {
	*o = BlackCard{}
	return nil
}

func blackCardBeforeUpsertHook(e boil.Executor, o *BlackCard) error {
	*o = BlackCard{}
	return nil
}

func blackCardAfterUpsertHook(e boil.Executor, o *BlackCard) error {
	*o = BlackCard{}
	return nil
}

func testBlackCardsHooks(t *testing.T) {
	t.Parallel()

	var err error

	empty := &BlackCard{}
	o := &BlackCard{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, blackCardDBTypes, false); err != nil {
		t.Errorf("Unable to randomize BlackCard object: %s", err)
	}

	AddBlackCardHook(boil.BeforeInsertHook, blackCardBeforeInsertHook)
	if err = o.doBeforeInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	blackCardBeforeInsertHooks = []BlackCardHook{}

	AddBlackCardHook(boil.AfterInsertHook, blackCardAfterInsertHook)
	if err = o.doAfterInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	blackCardAfterInsertHooks = []BlackCardHook{}

	AddBlackCardHook(boil.AfterSelectHook, blackCardAfterSelectHook)
	if err = o.doAfterSelectHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	blackCardAfterSelectHooks = []BlackCardHook{}

	AddBlackCardHook(boil.BeforeUpdateHook, blackCardBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	blackCardBeforeUpdateHooks = []BlackCardHook{}

	AddBlackCardHook(boil.AfterUpdateHook, blackCardAfterUpdateHook)
	if err = o.doAfterUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	blackCardAfterUpdateHooks = []BlackCardHook{}

	AddBlackCardHook(boil.BeforeDeleteHook, blackCardBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	blackCardBeforeDeleteHooks = []BlackCardHook{}

	AddBlackCardHook(boil.AfterDeleteHook, blackCardAfterDeleteHook)
	if err = o.doAfterDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	blackCardAfterDeleteHooks = []BlackCardHook{}

	AddBlackCardHook(boil.BeforeUpsertHook, blackCardBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	blackCardBeforeUpsertHooks = []BlackCardHook{}

	AddBlackCardHook(boil.AfterUpsertHook, blackCardAfterUpsertHook)
	if err = o.doAfterUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	blackCardAfterUpsertHooks = []BlackCardHook{}
}
func testBlackCardsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	blackCard := &BlackCard{}
	if err = randomize.Struct(seed, blackCard, blackCardDBTypes, true, blackCardColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BlackCard struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = blackCard.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := BlackCards(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testBlackCardsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	blackCard := &BlackCard{}
	if err = randomize.Struct(seed, blackCard, blackCardDBTypes, true); err != nil {
		t.Errorf("Unable to randomize BlackCard struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = blackCard.Insert(tx, blackCardColumnsWithoutDefault...); err != nil {
		t.Error(err)
	}

	count, err := BlackCards(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testBlackCardToOneSetUsingSet(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var local BlackCard
	var foreign Set

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, blackCardDBTypes, false, blackCardColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BlackCard struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, setDBTypes, false, setColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Set struct: %s", err)
	}

	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	local.SetID = foreign.ID
	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.Set(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.ID != foreign.ID {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := BlackCardSlice{&local}
	if err = local.L.LoadSet(tx, false, (*[]*BlackCard)(&slice)); err != nil {
		t.Fatal(err)
	}
	if local.R.Set == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Set = nil
	if err = local.L.LoadSet(tx, true, &local); err != nil {
		t.Fatal(err)
	}
	if local.R.Set == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testBlackCardToOneSetOpSetUsingSet(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a BlackCard
	var b, c Set

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, blackCardDBTypes, false, strmangle.SetComplement(blackCardPrimaryKeyColumns, blackCardColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, setDBTypes, false, strmangle.SetComplement(setPrimaryKeyColumns, setColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, setDBTypes, false, strmangle.SetComplement(setPrimaryKeyColumns, setColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Set{&b, &c} {
		err = a.SetSet(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Set != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.BlackCards[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.SetID != x.ID {
			t.Error("foreign key was wrong value", a.SetID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.SetID))
		reflect.Indirect(reflect.ValueOf(&a.SetID)).Set(zero)

		if err = a.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.SetID != x.ID {
			t.Error("foreign key was wrong value", a.SetID, x.ID)
		}
	}
}
func testBlackCardsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	blackCard := &BlackCard{}
	if err = randomize.Struct(seed, blackCard, blackCardDBTypes, true, blackCardColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BlackCard struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = blackCard.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = blackCard.Reload(tx); err != nil {
		t.Error(err)
	}
}

func testBlackCardsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	blackCard := &BlackCard{}
	if err = randomize.Struct(seed, blackCard, blackCardDBTypes, true, blackCardColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BlackCard struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = blackCard.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := BlackCardSlice{blackCard}

	if err = slice.ReloadAll(tx); err != nil {
		t.Error(err)
	}
}
func testBlackCardsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	blackCard := &BlackCard{}
	if err = randomize.Struct(seed, blackCard, blackCardDBTypes, true, blackCardColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BlackCard struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = blackCard.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := BlackCards(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	blackCardDBTypes = map[string]string{`CreatedAt`: `timestamp without time zone`, `DeletedAt`: `timestamp without time zone`, `ID`: `integer`, `Pick`: `integer`, `SetID`: `integer`, `Text`: `text`, `UpdatedAt`: `timestamp without time zone`}
	_                = bytes.MinRead
)

func testBlackCardsUpdate(t *testing.T) {
	t.Parallel()

	if len(blackCardColumns) == len(blackCardPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	blackCard := &BlackCard{}
	if err = randomize.Struct(seed, blackCard, blackCardDBTypes, true, blackCardColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BlackCard struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = blackCard.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := BlackCards(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, blackCard, blackCardDBTypes, true, blackCardColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BlackCard struct: %s", err)
	}

	if err = blackCard.Update(tx); err != nil {
		t.Error(err)
	}
}

func testBlackCardsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(blackCardColumns) == len(blackCardPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	blackCard := &BlackCard{}
	if err = randomize.Struct(seed, blackCard, blackCardDBTypes, true, blackCardColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BlackCard struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = blackCard.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := BlackCards(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, blackCard, blackCardDBTypes, true, blackCardPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize BlackCard struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(blackCardColumns, blackCardPrimaryKeyColumns) {
		fields = blackCardColumns
	} else {
		fields = strmangle.SetComplement(
			blackCardColumns,
			blackCardPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(blackCard))
	updateMap := M{}
	for _, col := range fields {
		updateMap[col] = value.FieldByName(strmangle.TitleCase(col)).Interface()
	}

	slice := BlackCardSlice{blackCard}
	if err = slice.UpdateAll(tx, updateMap); err != nil {
		t.Error(err)
	}
}
func testBlackCardsUpsert(t *testing.T) {
	t.Parallel()

	if len(blackCardColumns) == len(blackCardPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	blackCard := BlackCard{}
	if err = randomize.Struct(seed, &blackCard, blackCardDBTypes, true); err != nil {
		t.Errorf("Unable to randomize BlackCard struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = blackCard.Upsert(tx, false, nil, nil); err != nil {
		t.Errorf("Unable to upsert BlackCard: %s", err)
	}

	count, err := BlackCards(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &blackCard, blackCardDBTypes, false, blackCardPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize BlackCard struct: %s", err)
	}

	if err = blackCard.Upsert(tx, true, nil, nil); err != nil {
		t.Errorf("Unable to upsert BlackCard: %s", err)
	}

	count, err = BlackCards(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
