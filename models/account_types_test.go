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

func testAccountTypes(t *testing.T) {
	t.Parallel()

	query := AccountTypes(nil)

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}
func testAccountTypesDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	accountType := &AccountType{}
	if err = randomize.Struct(seed, accountType, accountTypeDBTypes, true, accountTypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AccountType struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = accountType.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = accountType.Delete(tx); err != nil {
		t.Error(err)
	}

	count, err := AccountTypes(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testAccountTypesQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	accountType := &AccountType{}
	if err = randomize.Struct(seed, accountType, accountTypeDBTypes, true, accountTypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AccountType struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = accountType.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = AccountTypes(tx).DeleteAll(); err != nil {
		t.Error(err)
	}

	count, err := AccountTypes(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testAccountTypesSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	accountType := &AccountType{}
	if err = randomize.Struct(seed, accountType, accountTypeDBTypes, true, accountTypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AccountType struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = accountType.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := AccountTypeSlice{accountType}

	if err = slice.DeleteAll(tx); err != nil {
		t.Error(err)
	}

	count, err := AccountTypes(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}
func testAccountTypesExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	accountType := &AccountType{}
	if err = randomize.Struct(seed, accountType, accountTypeDBTypes, true, accountTypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AccountType struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = accountType.Insert(tx); err != nil {
		t.Error(err)
	}

	e, err := AccountTypeExists(tx, accountType.ID)
	if err != nil {
		t.Errorf("Unable to check if AccountType exists: %s", err)
	}
	if !e {
		t.Errorf("Expected AccountTypeExistsG to return true, but got false.")
	}
}
func testAccountTypesFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	accountType := &AccountType{}
	if err = randomize.Struct(seed, accountType, accountTypeDBTypes, true, accountTypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AccountType struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = accountType.Insert(tx); err != nil {
		t.Error(err)
	}

	accountTypeFound, err := FindAccountType(tx, accountType.ID)
	if err != nil {
		t.Error(err)
	}

	if accountTypeFound == nil {
		t.Error("want a record, got nil")
	}
}
func testAccountTypesBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	accountType := &AccountType{}
	if err = randomize.Struct(seed, accountType, accountTypeDBTypes, true, accountTypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AccountType struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = accountType.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = AccountTypes(tx).Bind(accountType); err != nil {
		t.Error(err)
	}
}

func testAccountTypesOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	accountType := &AccountType{}
	if err = randomize.Struct(seed, accountType, accountTypeDBTypes, true, accountTypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AccountType struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = accountType.Insert(tx); err != nil {
		t.Error(err)
	}

	if x, err := AccountTypes(tx).One(); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testAccountTypesAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	accountTypeOne := &AccountType{}
	accountTypeTwo := &AccountType{}
	if err = randomize.Struct(seed, accountTypeOne, accountTypeDBTypes, false, accountTypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AccountType struct: %s", err)
	}
	if err = randomize.Struct(seed, accountTypeTwo, accountTypeDBTypes, false, accountTypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AccountType struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = accountTypeOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = accountTypeTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := AccountTypes(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testAccountTypesCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	accountTypeOne := &AccountType{}
	accountTypeTwo := &AccountType{}
	if err = randomize.Struct(seed, accountTypeOne, accountTypeDBTypes, false, accountTypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AccountType struct: %s", err)
	}
	if err = randomize.Struct(seed, accountTypeTwo, accountTypeDBTypes, false, accountTypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AccountType struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = accountTypeOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = accountTypeTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := AccountTypes(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}
func accountTypeBeforeInsertHook(e boil.Executor, o *AccountType) error {
	*o = AccountType{}
	return nil
}

func accountTypeAfterInsertHook(e boil.Executor, o *AccountType) error {
	*o = AccountType{}
	return nil
}

func accountTypeAfterSelectHook(e boil.Executor, o *AccountType) error {
	*o = AccountType{}
	return nil
}

func accountTypeBeforeUpdateHook(e boil.Executor, o *AccountType) error {
	*o = AccountType{}
	return nil
}

func accountTypeAfterUpdateHook(e boil.Executor, o *AccountType) error {
	*o = AccountType{}
	return nil
}

func accountTypeBeforeDeleteHook(e boil.Executor, o *AccountType) error {
	*o = AccountType{}
	return nil
}

func accountTypeAfterDeleteHook(e boil.Executor, o *AccountType) error {
	*o = AccountType{}
	return nil
}

func accountTypeBeforeUpsertHook(e boil.Executor, o *AccountType) error {
	*o = AccountType{}
	return nil
}

func accountTypeAfterUpsertHook(e boil.Executor, o *AccountType) error {
	*o = AccountType{}
	return nil
}

func testAccountTypesHooks(t *testing.T) {
	t.Parallel()

	var err error

	empty := &AccountType{}
	o := &AccountType{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, accountTypeDBTypes, false); err != nil {
		t.Errorf("Unable to randomize AccountType object: %s", err)
	}

	AddAccountTypeHook(boil.BeforeInsertHook, accountTypeBeforeInsertHook)
	if err = o.doBeforeInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	accountTypeBeforeInsertHooks = []AccountTypeHook{}

	AddAccountTypeHook(boil.AfterInsertHook, accountTypeAfterInsertHook)
	if err = o.doAfterInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	accountTypeAfterInsertHooks = []AccountTypeHook{}

	AddAccountTypeHook(boil.AfterSelectHook, accountTypeAfterSelectHook)
	if err = o.doAfterSelectHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	accountTypeAfterSelectHooks = []AccountTypeHook{}

	AddAccountTypeHook(boil.BeforeUpdateHook, accountTypeBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	accountTypeBeforeUpdateHooks = []AccountTypeHook{}

	AddAccountTypeHook(boil.AfterUpdateHook, accountTypeAfterUpdateHook)
	if err = o.doAfterUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	accountTypeAfterUpdateHooks = []AccountTypeHook{}

	AddAccountTypeHook(boil.BeforeDeleteHook, accountTypeBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	accountTypeBeforeDeleteHooks = []AccountTypeHook{}

	AddAccountTypeHook(boil.AfterDeleteHook, accountTypeAfterDeleteHook)
	if err = o.doAfterDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	accountTypeAfterDeleteHooks = []AccountTypeHook{}

	AddAccountTypeHook(boil.BeforeUpsertHook, accountTypeBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	accountTypeBeforeUpsertHooks = []AccountTypeHook{}

	AddAccountTypeHook(boil.AfterUpsertHook, accountTypeAfterUpsertHook)
	if err = o.doAfterUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	accountTypeAfterUpsertHooks = []AccountTypeHook{}
}
func testAccountTypesInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	accountType := &AccountType{}
	if err = randomize.Struct(seed, accountType, accountTypeDBTypes, true, accountTypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AccountType struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = accountType.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := AccountTypes(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testAccountTypesInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	accountType := &AccountType{}
	if err = randomize.Struct(seed, accountType, accountTypeDBTypes, true); err != nil {
		t.Errorf("Unable to randomize AccountType struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = accountType.Insert(tx, accountTypeColumnsWithoutDefault...); err != nil {
		t.Error(err)
	}

	count, err := AccountTypes(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testAccountTypeToManyUserAccountTypeHistories(t *testing.T) {
	var err error
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a AccountType
	var b, c UserAccountTypeHistory

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, accountTypeDBTypes, true, accountTypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AccountType struct: %s", err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}

	randomize.Struct(seed, &b, userAccountTypeHistoryDBTypes, false, userAccountTypeHistoryColumnsWithDefault...)
	randomize.Struct(seed, &c, userAccountTypeHistoryDBTypes, false, userAccountTypeHistoryColumnsWithDefault...)

	b.AccountTypeID = a.ID
	c.AccountTypeID = a.ID
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(tx); err != nil {
		t.Fatal(err)
	}

	userAccountTypeHistory, err := a.UserAccountTypeHistories(tx).All()
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range userAccountTypeHistory {
		if v.AccountTypeID == b.AccountTypeID {
			bFound = true
		}
		if v.AccountTypeID == c.AccountTypeID {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := AccountTypeSlice{&a}
	if err = a.L.LoadUserAccountTypeHistories(tx, false, (*[]*AccountType)(&slice)); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.UserAccountTypeHistories); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.UserAccountTypeHistories = nil
	if err = a.L.LoadUserAccountTypeHistories(tx, true, &a); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.UserAccountTypeHistories); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", userAccountTypeHistory)
	}
}

func testAccountTypeToManyCurrentAccountTypeUsers(t *testing.T) {
	var err error
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a AccountType
	var b, c User

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, accountTypeDBTypes, true, accountTypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AccountType struct: %s", err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}

	randomize.Struct(seed, &b, userDBTypes, false, userColumnsWithDefault...)
	randomize.Struct(seed, &c, userDBTypes, false, userColumnsWithDefault...)

	b.CurrentAccountTypeID = a.ID
	c.CurrentAccountTypeID = a.ID
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(tx); err != nil {
		t.Fatal(err)
	}

	user, err := a.CurrentAccountTypeUsers(tx).All()
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range user {
		if v.CurrentAccountTypeID == b.CurrentAccountTypeID {
			bFound = true
		}
		if v.CurrentAccountTypeID == c.CurrentAccountTypeID {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := AccountTypeSlice{&a}
	if err = a.L.LoadCurrentAccountTypeUsers(tx, false, (*[]*AccountType)(&slice)); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.CurrentAccountTypeUsers); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.CurrentAccountTypeUsers = nil
	if err = a.L.LoadCurrentAccountTypeUsers(tx, true, &a); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.CurrentAccountTypeUsers); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", user)
	}
}

func testAccountTypeToManyAddOpUserAccountTypeHistories(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a AccountType
	var b, c, d, e UserAccountTypeHistory

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, accountTypeDBTypes, false, strmangle.SetComplement(accountTypePrimaryKeyColumns, accountTypeColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*UserAccountTypeHistory{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, userAccountTypeHistoryDBTypes, false, strmangle.SetComplement(userAccountTypeHistoryPrimaryKeyColumns, userAccountTypeHistoryColumnsWithoutDefault)...); err != nil {
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

	foreignersSplitByInsertion := [][]*UserAccountTypeHistory{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddUserAccountTypeHistories(tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.ID != first.AccountTypeID {
			t.Error("foreign key was wrong value", a.ID, first.AccountTypeID)
		}
		if a.ID != second.AccountTypeID {
			t.Error("foreign key was wrong value", a.ID, second.AccountTypeID)
		}

		if first.R.AccountType != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.AccountType != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.UserAccountTypeHistories[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.UserAccountTypeHistories[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.UserAccountTypeHistories(tx).Count()
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}
func testAccountTypeToManyAddOpCurrentAccountTypeUsers(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a AccountType
	var b, c, d, e User

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, accountTypeDBTypes, false, strmangle.SetComplement(accountTypePrimaryKeyColumns, accountTypeColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*User{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, userDBTypes, false, strmangle.SetComplement(userPrimaryKeyColumns, userColumnsWithoutDefault)...); err != nil {
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

	foreignersSplitByInsertion := [][]*User{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddCurrentAccountTypeUsers(tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.ID != first.CurrentAccountTypeID {
			t.Error("foreign key was wrong value", a.ID, first.CurrentAccountTypeID)
		}
		if a.ID != second.CurrentAccountTypeID {
			t.Error("foreign key was wrong value", a.ID, second.CurrentAccountTypeID)
		}

		if first.R.CurrentAccountType != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.CurrentAccountType != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.CurrentAccountTypeUsers[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.CurrentAccountTypeUsers[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.CurrentAccountTypeUsers(tx).Count()
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}

func testAccountTypesReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	accountType := &AccountType{}
	if err = randomize.Struct(seed, accountType, accountTypeDBTypes, true, accountTypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AccountType struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = accountType.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = accountType.Reload(tx); err != nil {
		t.Error(err)
	}
}

func testAccountTypesReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	accountType := &AccountType{}
	if err = randomize.Struct(seed, accountType, accountTypeDBTypes, true, accountTypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AccountType struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = accountType.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := AccountTypeSlice{accountType}

	if err = slice.ReloadAll(tx); err != nil {
		t.Error(err)
	}
}
func testAccountTypesSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	accountType := &AccountType{}
	if err = randomize.Struct(seed, accountType, accountTypeDBTypes, true, accountTypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AccountType struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = accountType.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := AccountTypes(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	accountTypeDBTypes = map[string]string{`APIKeyLimit`: `integer`, `CreatedAt`: `timestamp without time zone`, `ID`: `integer`, `RequestLimit`: `integer`, `Type`: `enum.account_type('basic','admin')`, `UpdatedAt`: `timestamp without time zone`}
	_                  = bytes.MinRead
)

func testAccountTypesUpdate(t *testing.T) {
	t.Parallel()

	if len(accountTypeColumns) == len(accountTypePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	accountType := &AccountType{}
	if err = randomize.Struct(seed, accountType, accountTypeDBTypes, true, accountTypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AccountType struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = accountType.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := AccountTypes(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, accountType, accountTypeDBTypes, true, accountTypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AccountType struct: %s", err)
	}

	if err = accountType.Update(tx); err != nil {
		t.Error(err)
	}
}

func testAccountTypesSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(accountTypeColumns) == len(accountTypePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	accountType := &AccountType{}
	if err = randomize.Struct(seed, accountType, accountTypeDBTypes, true, accountTypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize AccountType struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = accountType.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := AccountTypes(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, accountType, accountTypeDBTypes, true, accountTypePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize AccountType struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(accountTypeColumns, accountTypePrimaryKeyColumns) {
		fields = accountTypeColumns
	} else {
		fields = strmangle.SetComplement(
			accountTypeColumns,
			accountTypePrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(accountType))
	updateMap := M{}
	for _, col := range fields {
		updateMap[col] = value.FieldByName(strmangle.TitleCase(col)).Interface()
	}

	slice := AccountTypeSlice{accountType}
	if err = slice.UpdateAll(tx, updateMap); err != nil {
		t.Error(err)
	}
}
func testAccountTypesUpsert(t *testing.T) {
	t.Parallel()

	if len(accountTypeColumns) == len(accountTypePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	accountType := AccountType{}
	if err = randomize.Struct(seed, &accountType, accountTypeDBTypes, true); err != nil {
		t.Errorf("Unable to randomize AccountType struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = accountType.Upsert(tx, false, nil, nil); err != nil {
		t.Errorf("Unable to upsert AccountType: %s", err)
	}

	count, err := AccountTypes(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &accountType, accountTypeDBTypes, false, accountTypePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize AccountType struct: %s", err)
	}

	if err = accountType.Upsert(tx, true, nil, nil); err != nil {
		t.Errorf("Unable to upsert AccountType: %s", err)
	}

	count, err = AccountTypes(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}