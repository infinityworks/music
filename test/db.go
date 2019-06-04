package test

import (
	"database/sql"
	_ "github.com/infinityworks/music/migrations"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"github.com/pressly/goose"
	"github.com/steinfletcher/apitest"
	"github.com/steinfletcher/apitest/x/db"
)

const DBAddr = "./music.db"
const driver = "sqlite3"

var Recorder *apitest.Recorder

func init() {
	Recorder = apitest.NewTestRecorder()
	wrappedDriver := db.WrapWithRecorder(driver, Recorder)
	sql.Register("wrappedSqlite", wrappedDriver)
}

func DBSetup(setup func(tx *sqlx.Tx)) *sqlx.DB {
	d, err := sqlx.Connect(driver, DBAddr)
	if err != nil {
		panic(err)
	}

	goose.SetDialect(driver)
	err = goose.Up(d.DB, ".")
	if err != nil {
		panic(err)
	}

	tx, err := d.Beginx()
	if err != nil {
		panic(err)
	}

	setup(tx)

	err = tx.Commit()
	if err != nil {
		panic(err)
	}
	return d
}

func DBConnect() *sqlx.DB {
	testDB, err := sqlx.Connect("wrappedSqlite", DBAddr)
	if err != nil {
		panic(err)
	}
	return testDB
}
