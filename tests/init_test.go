package tests

import (
	"database/sql"
	"fmt"

	sqle "github.com/dolthub/go-mysql-server"
	"github.com/dolthub/go-mysql-server/memory"
	"github.com/dolthub/go-mysql-server/server"
	gsql "github.com/dolthub/go-mysql-server/sql"
	"github.com/dolthub/go-mysql-server/sql/information_schema"
	"github.com/olachat/gola/coredb"
	"github.com/olachat/gola/drivers/mysqldriver"
	"github.com/olachat/gola/golalib/testdata"
	_ "modernc.org/sqlite"
)

const (
	testDBPort int    = 33067
	testDBName string = "testdb"
)

var tableNames = []string{"users", "blogs", "songs", "song_user_favourites", "profile", "account"}

func init() {
	engine := sqle.NewDefault(gsql.NewDatabaseProvider(
		memory.NewDatabase(testDBName),
		information_schema.NewInformationSchemaDatabase(),
	))

	config := server.Config{
		Protocol: "tcp",
		Address:  fmt.Sprintf("localhost:%d", testDBPort),
	}

	s, err := server.NewDefaultServer(config, engine)
	if err != nil {
		panic(err)
	}

	go s.Start()

	// setup sqlite memdb
	memdb, err := sql.Open("sqlite", ":memory:")
	if err != nil {
		panic(err)
	}
	query, _ := testdata.Fixtures.ReadFile("sqlite.sql")
	_, err = memdb.Exec(string(query))
	if err != nil {
		panic(err.Error())
	}

	// connect to mysql memdb
	connstr := mysqldriver.MySQLBuildQueryString("root", "", testDBName, "localhost", testDBPort, "false")
	db, err := sql.Open("mysql", connstr)
	if err != nil {
		panic(err)
	}

	coredb.Setup(func(dbname string, mode coredb.DBMode) *sql.DB {
		if dbname == "memory" {
			return memdb
		}

		return db
	})

	//create tables
	for _, tableName := range tableNames {
		query, _ := testdata.Fixtures.ReadFile(tableName + ".sql")
		db.Exec(string(query))
	}

	//add data
	db.Exec(`
insert into users (name, email, created_at, updated_at, float_type, double_type, hobby, hobby_no_default, sports_no_default, sports) values
("John Doe", "john@doe.com", NOW(), NOW(), 1.55555, 1.8729, 'running','swimming', ('SWIM,TENNIS'), ("TENNIS")),
("John Doe", "johnalt@doe.com", NOW(), NOW(), 2.5, 2.8239, 'swimming','running', ('BASKETBALL'), ("FOOTBALL")),
("Jane Doe", "jane@doe.com", NOW(), NOW(), 3.5, 334.8593, 'singing','swimming', ('SQUASH,BADMINTON'), ("SQUASH,TENNIS")),
("Evil Bob", "evilbob@gmail.com", NOW(), NOW(), 4.5, 42234.83, 'singing','running', ('TENNIS'), ('BADMINTON,BASKETBALL'))
	`)

}
