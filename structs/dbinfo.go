package structs

import "strings"

// DBInfo represents information about a database and used for codegen
type DBInfo struct {
	Schema  string   `json:"schema"`
	Tables  []*Table `json:"tables"`
	version string
}

// SetVersion for code gen
func (db *DBInfo) SetVersion(version string) {
	db.version = version
}

// GetVersion for code gen
func (db *DBInfo) GetVersion() string {
	return db.version
}

// GetName for code gen
func (db *DBInfo) GetName() string {
	return db.Schema
}

// Package returns package name
func (db *DBInfo) Package() string {
	return strings.ToLower(db.Schema)
}
