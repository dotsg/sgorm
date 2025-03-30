package golalib

import (
	"database/sql"
	"path/filepath"
	"testing"

	"github.com/olachat/gola/drivers"
	"github.com/olachat/gola/drivers/sqlite3driver"
	"github.com/olachat/gola/ormtpl"
	"github.com/olachat/gola/structs"
	_ "modernc.org/sqlite"
)

func init() {
	memdb, _ := sql.Open("sqlite", ":memory:")
	var err error

	query, _ := fixtures.ReadFile("testdata" + string(filepath.Separator) + "sqlite.sql")
	_, err = memdb.Exec(string(query))
	if err != nil {
		panic(err.Error())
	}

	sqlite3driver.SetMemdb(memdb)
}

func getSQLiteDB() *structs.DBInfo {
	var config drivers.Config = map[string]any{
		"dbname": ":memory:",
		"output": "testdata",
	}

	m := &sqlite3driver.SQLiteDriver{}
	db, err := m.Assemble(config)
	if err != nil {
		panic(err)
	}
	return db
}
func TestSQLiteCodeGen(t *testing.T) {
	db := getSQLiteDB()
	gen := &CodeGen{"mysql"}

	for _, table := range db.Tables {
		testGen(t, func(t ormtpl.TplStruct) map[string][]byte {
			return gen.GenORM(t.(*structs.Table))
		}, table)
	}
}
