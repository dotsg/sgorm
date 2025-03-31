package golalib

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/olachat/gola/drivers"
	"github.com/olachat/gola/drivers/mysqldriver"
	"github.com/olachat/gola/drivers/sqlite3driver"
	"github.com/olachat/gola/structs"
)

func RunSqlite(config drivers.Config) int {
	s := &sqlite3driver.SQLiteDriver{}
	output := config.DefaultString("output", "temp")

	db, err := s.Assemble(config)
	if err != nil {
		panic(err)
	}

	gen := &CodeGen{"sqlite"}
	return genCode(gen, db, output)
}

/*
Run gola to perform code gen for MySql

`output`: output folder path
*/
func RunMySql(config drivers.Config) int {
	dbconfig := drivers.NewDBConfig(config)
	output := config.DefaultString("output", "temp")

	m := &mysqldriver.MySQLDriver{}
	db, err := m.Assemble(dbconfig)
	if err != nil {
		panic(err)
	}

	gen := &CodeGen{"mysql"}
	return genCode(gen, db, output)
}

func genCode(gen *CodeGen, db *structs.DBInfo, output string) int {
	if !strings.HasPrefix(output, "/") {
		// output folder is relative path
		wd, err := os.Getwd()
		if err != nil {
			wd = "."
		}
		output = wd + string(filepath.Separator) + output
	}

	if !strings.HasSuffix(output, string(filepath.Separator)) {
		output = output + string(filepath.Separator)
	}

	for _, t := range db.Tables {
		if len(t.GetPKColumns()) == 0 {
			println(t.Name + " doesn't have primay key")
			continue
		}

		files := gen.GenORM(t)
		needMkdir := true
		for path, data := range files {
			if needMkdir {
				pos := strings.LastIndex(path, string(filepath.Separator))
				expectedFileFolder := output + path[0:pos]
				err := os.Mkdir(expectedFileFolder, os.ModePerm)
				if err != nil && os.IsNotExist(err) {
					println("Failed to create folder, please ensure " + output[:len(output)-1] + " exists")
					return 1
				}
				needMkdir = false
			}

			os.WriteFile(output+path, data, 0644)
		}
	}

	files := gen.GenPackage(db)
	for path, data := range files {
		os.WriteFile(output+path, data, 0644)
	}

	fmt.Printf("code generated in %s\n", output[:len(output)-1])
	return 0
}
