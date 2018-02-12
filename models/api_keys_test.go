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

func testAPIKeys(t *testing.T) {
	t.Parallel()

	query := APIKeys(nil)

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}
func testAPIKeysDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	apiKey := &APIKey{}
	if err = randomize.Struct(seed, apiKey, apiKeyDBTypes, true, apiKeyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize APIKey struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = apiKey.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = apiKey.Delete(tx); err != nil {
		t.Error(err)
	}

	count, err := APIKeys(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testAPIKeysQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	apiKey := &APIKey{}
	if err = randomize.Struct(seed, apiKey, apiKeyDBTypes, true, apiKeyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize APIKey struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = apiKey.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = APIKeys(tx).DeleteAll(); err != nil {
		t.Error(err)
	}

	count, err := APIKeys(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testAPIKeysSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	apiKey := &APIKey{}
	if err = randomize.Struct(seed, apiKey, apiKeyDBTypes, true, apiKeyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize APIKey struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = apiKey.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := APIKeySlice{apiKey}

	if err = slice.DeleteAll(tx); err != nil {
		t.Error(err)
	}

	count, err := APIKeys(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}
func testAPIKeysExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	apiKey := &APIKey{}
	if err = randomize.Struct(seed, apiKey, apiKeyDBTypes, true, apiKeyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize APIKey struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = apiKey.Insert(tx); err != nil {
		t.Error(err)
	}

	e, err := APIKeyExists(tx, apiKey.ID)
	if err != nil {
		t.Errorf("Unable to check if APIKey exists: %s", err)
	}
	if !e {
		t.Errorf("Expected APIKeyExistsG to return true, but got false.")
	}
}
func testAPIKeysFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	apiKey := &APIKey{}
	if err = randomize.Struct(seed, apiKey, apiKeyDBTypes, true, apiKeyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize APIKey struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = apiKey.Insert(tx); err != nil {
		t.Error(err)
	}

	apiKeyFound, err := FindAPIKey(tx, apiKey.ID)
	if err != nil {
		t.Error(err)
	}

	if apiKeyFound == nil {
		t.Error("want a record, got nil")
	}
}
func testAPIKeysBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	apiKey := &APIKey{}
	if err = randomize.Struct(seed, apiKey, apiKeyDBTypes, true, apiKeyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize APIKey struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = apiKey.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = APIKeys(tx).Bind(apiKey); err != nil {
		t.Error(err)
	}
}

func testAPIKeysOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	apiKey := &APIKey{}
	if err = randomize.Struct(seed, apiKey, apiKeyDBTypes, true, apiKeyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize APIKey struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = apiKey.Insert(tx); err != nil {
		t.Error(err)
	}

	if x, err := APIKeys(tx).One(); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testAPIKeysAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	apiKeyOne := &APIKey{}
	apiKeyTwo := &APIKey{}
	if err = randomize.Struct(seed, apiKeyOne, apiKeyDBTypes, false, apiKeyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize APIKey struct: %s", err)
	}
	if err = randomize.Struct(seed, apiKeyTwo, apiKeyDBTypes, false, apiKeyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize APIKey struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = apiKeyOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = apiKeyTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := APIKeys(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testAPIKeysCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	apiKeyOne := &APIKey{}
	apiKeyTwo := &APIKey{}
	if err = randomize.Struct(seed, apiKeyOne, apiKeyDBTypes, false, apiKeyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize APIKey struct: %s", err)
	}
	if err = randomize.Struct(seed, apiKeyTwo, apiKeyDBTypes, false, apiKeyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize APIKey struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = apiKeyOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = apiKeyTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := APIKeys(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}
func apiKeyBeforeInsertHook(e boil.Executor, o *APIKey) error {
	*o = APIKey{}
	return nil
}

func apiKeyAfterInsertHook(e boil.Executor, o *APIKey) error {
	*o = APIKey{}
	return nil
}

func apiKeyAfterSelectHook(e boil.Executor, o *APIKey) error {
	*o = APIKey{}
	return nil
}

func apiKeyBeforeUpdateHook(e boil.Executor, o *APIKey) error {
	*o = APIKey{}
	return nil
}

func apiKeyAfterUpdateHook(e boil.Executor, o *APIKey) error {
	*o = APIKey{}
	return nil
}

func apiKeyBeforeDeleteHook(e boil.Executor, o *APIKey) error {
	*o = APIKey{}
	return nil
}

func apiKeyAfterDeleteHook(e boil.Executor, o *APIKey) error {
	*o = APIKey{}
	return nil
}

func apiKeyBeforeUpsertHook(e boil.Executor, o *APIKey) error {
	*o = APIKey{}
	return nil
}

func apiKeyAfterUpsertHook(e boil.Executor, o *APIKey) error {
	*o = APIKey{}
	return nil
}

func testAPIKeysHooks(t *testing.T) {
	t.Parallel()

	var err error

	empty := &APIKey{}
	o := &APIKey{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, apiKeyDBTypes, false); err != nil {
		t.Errorf("Unable to randomize APIKey object: %s", err)
	}

	AddAPIKeyHook(boil.BeforeInsertHook, apiKeyBeforeInsertHook)
	if err = o.doBeforeInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	apiKeyBeforeInsertHooks = []APIKeyHook{}

	AddAPIKeyHook(boil.AfterInsertHook, apiKeyAfterInsertHook)
	if err = o.doAfterInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	apiKeyAfterInsertHooks = []APIKeyHook{}

	AddAPIKeyHook(boil.AfterSelectHook, apiKeyAfterSelectHook)
	if err = o.doAfterSelectHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	apiKeyAfterSelectHooks = []APIKeyHook{}

	AddAPIKeyHook(boil.BeforeUpdateHook, apiKeyBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	apiKeyBeforeUpdateHooks = []APIKeyHook{}

	AddAPIKeyHook(boil.AfterUpdateHook, apiKeyAfterUpdateHook)
	if err = o.doAfterUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	apiKeyAfterUpdateHooks = []APIKeyHook{}

	AddAPIKeyHook(boil.BeforeDeleteHook, apiKeyBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	apiKeyBeforeDeleteHooks = []APIKeyHook{}

	AddAPIKeyHook(boil.AfterDeleteHook, apiKeyAfterDeleteHook)
	if err = o.doAfterDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	apiKeyAfterDeleteHooks = []APIKeyHook{}

	AddAPIKeyHook(boil.BeforeUpsertHook, apiKeyBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	apiKeyBeforeUpsertHooks = []APIKeyHook{}

	AddAPIKeyHook(boil.AfterUpsertHook, apiKeyAfterUpsertHook)
	if err = o.doAfterUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	apiKeyAfterUpsertHooks = []APIKeyHook{}
}
func testAPIKeysInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	apiKey := &APIKey{}
	if err = randomize.Struct(seed, apiKey, apiKeyDBTypes, true, apiKeyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize APIKey struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = apiKey.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := APIKeys(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testAPIKeysInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	apiKey := &APIKey{}
	if err = randomize.Struct(seed, apiKey, apiKeyDBTypes, true); err != nil {
		t.Errorf("Unable to randomize APIKey struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = apiKey.Insert(tx, apiKeyColumnsWithoutDefault...); err != nil {
		t.Error(err)
	}

	count, err := APIKeys(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testAPIKeyToManyUserAPIKeys(t *testing.T) {
	var err error
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a APIKey
	var b, c UserAPIKey

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, apiKeyDBTypes, true, apiKeyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize APIKey struct: %s", err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}

	randomize.Struct(seed, &b, userAPIKeyDBTypes, false, userAPIKeyColumnsWithDefault...)
	randomize.Struct(seed, &c, userAPIKeyDBTypes, false, userAPIKeyColumnsWithDefault...)

	b.APIKeyID = a.ID
	c.APIKeyID = a.ID
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(tx); err != nil {
		t.Fatal(err)
	}

	userAPIKey, err := a.UserAPIKeys(tx).All()
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range userAPIKey {
		if v.APIKeyID == b.APIKeyID {
			bFound = true
		}
		if v.APIKeyID == c.APIKeyID {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := APIKeySlice{&a}
	if err = a.L.LoadUserAPIKeys(tx, false, (*[]*APIKey)(&slice)); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.UserAPIKeys); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.UserAPIKeys = nil
	if err = a.L.LoadUserAPIKeys(tx, true, &a); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.UserAPIKeys); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", userAPIKey)
	}
}

func testAPIKeyToManyAddOpUserAPIKeys(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a APIKey
	var b, c, d, e UserAPIKey

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, apiKeyDBTypes, false, strmangle.SetComplement(apiKeyPrimaryKeyColumns, apiKeyColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*UserAPIKey{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, userAPIKeyDBTypes, false, strmangle.SetComplement(userAPIKeyPrimaryKeyColumns, userAPIKeyColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(tx); err != nil {
		t.Fatal(err)
	}

	foreignersSplitByInsertion := [][]*UserAPIKey{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddUserAPIKeys(tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.ID != first.APIKeyID {
			t.Error("foreign key was wrong value", a.ID, first.APIKeyID)
		}
		if a.ID != second.APIKeyID {
			t.Error("foreign key was wrong value", a.ID, second.APIKeyID)
		}

		if first.R.APIKey != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.APIKey != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.UserAPIKeys[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.UserAPIKeys[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.UserAPIKeys(tx).Count()
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}

func testAPIKeysReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	apiKey := &APIKey{}
	if err = randomize.Struct(seed, apiKey, apiKeyDBTypes, true, apiKeyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize APIKey struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = apiKey.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = apiKey.Reload(tx); err != nil {
		t.Error(err)
	}
}

func testAPIKeysReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	apiKey := &APIKey{}
	if err = randomize.Struct(seed, apiKey, apiKeyDBTypes, true, apiKeyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize APIKey struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = apiKey.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := APIKeySlice{apiKey}

	if err = slice.ReloadAll(tx); err != nil {
		t.Error(err)
	}
}
func testAPIKeysSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	apiKey := &APIKey{}
	if err = randomize.Struct(seed, apiKey, apiKeyDBTypes, true, apiKeyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize APIKey struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = apiKey.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := APIKeys(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	apiKeyDBTypes = map[string]string{`APIKey`: `character varying`, `CreatedAt`: `timestamp without time zone`, `DeletedAt`: `timestamp without time zone`, `ID`: `integer`}
	_             = bytes.MinRead
)

func testAPIKeysUpdate(t *testing.T) {
	t.Parallel()

	if len(apiKeyColumns) == len(apiKeyPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	apiKey := &APIKey{}
	if err = randomize.Struct(seed, apiKey, apiKeyDBTypes, true, apiKeyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize APIKey struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = apiKey.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := APIKeys(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, apiKey, apiKeyDBTypes, true, apiKeyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize APIKey struct: %s", err)
	}

	if err = apiKey.Update(tx); err != nil {
		t.Error(err)
	}
}

func testAPIKeysSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(apiKeyColumns) == len(apiKeyPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	apiKey := &APIKey{}
	if err = randomize.Struct(seed, apiKey, apiKeyDBTypes, true, apiKeyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize APIKey struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = apiKey.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := APIKeys(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, apiKey, apiKeyDBTypes, true, apiKeyPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize APIKey struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(apiKeyColumns, apiKeyPrimaryKeyColumns) {
		fields = apiKeyColumns
	} else {
		fields = strmangle.SetComplement(
			apiKeyColumns,
			apiKeyPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(apiKey))
	updateMap := M{}
	for _, col := range fields {
		updateMap[col] = value.FieldByName(strmangle.TitleCase(col)).Interface()
	}

	slice := APIKeySlice{apiKey}
	if err = slice.UpdateAll(tx, updateMap); err != nil {
		t.Error(err)
	}
}
func testAPIKeysUpsert(t *testing.T) {
	t.Parallel()

	if len(apiKeyColumns) == len(apiKeyPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	apiKey := APIKey{}
	if err = randomize.Struct(seed, &apiKey, apiKeyDBTypes, true); err != nil {
		t.Errorf("Unable to randomize APIKey struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = apiKey.Upsert(tx, false, nil, nil); err != nil {
		t.Errorf("Unable to upsert APIKey: %s", err)
	}

	count, err := APIKeys(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &apiKey, apiKeyDBTypes, false, apiKeyPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize APIKey struct: %s", err)
	}

	if err = apiKey.Upsert(tx, true, nil, nil); err != nil {
		t.Errorf("Unable to upsert APIKey: %s", err)
	}

	count, err = APIKeys(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
