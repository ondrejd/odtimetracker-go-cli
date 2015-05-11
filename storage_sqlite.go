// Copyright 2015 Ondrej Donek. All rights reserved.
// See LICENSE file for more informations about licensing.

package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os/user"
	"path"
)

const (
	DefaultDatabaseName = ".odtimetracker.sqlite"
)

// Schema for our SQLite database
var schemaSql = `
CREATE TABLE Projects (
	ProjectId INTEGER PRIMARY KEY, 
	Name TEXT,
	Description TEXT,
	Created TEXT NOT NULL
);
CREATE TABLE Activities (
	ActivityId INTEGER PRIMARY KEY,
	ProjectId INTEGER NOT NULL,
	Name TEXT,
	Description TEXT,
	Tags TEXT,
	Started TEXT NOT NULL,
	Stopped TEXT NOT NULL DEFAULT '',
	FOREIGN KEY(ProjectId) REFERENCES Projects(ProjectId) 
);
PRAGMA user_version = 1;
`

// SQLite storage
var SqliteStorage = &Storage{
	Init:                  sqliteStorage_Init,
	InsertActivity:        sqliteStorage_InsertActivity,
	InsertProject:         sqliteStorage_InsertProject,
	RemoveActivity:        sqliteStorage_RemoveActivity,
	RemoveProject:         sqliteStorage_RemoveProject,
	SelectActivities:      sqliteStorage_SelectActivities,
	SelectActivityById:    sqliteStorage_SelectActivityById,
	SelectActivityRunning: sqliteStorage_SelectActivityRunning,
	SelectProjects:        sqliteStorage_SelectProjects,
	SelectProjectById:     sqliteStorage_SelectProjectById,
	SelectProjectByName:   sqliteStorage_SelectProjectByName,
}

// Initialize storage.
func sqliteStorage_Init() (db *sql.DB, err error) {
	dbPath, err := databasePath()
	log.Println(dbPath)
	if err != nil {
		return nil, err
	}

	db, err = sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, err
	}
	//defer db.Close() // Moved to `main.go`

	// Check if we are need to create schema
	ver, err := schemaVersion(db)
	log.Printf("Current database schema version: %d\n", ver)
	if err != nil {
		return db, err
	}
	if ver < 1 {
		err := schemaCreate(db)
		if err != nil {
			return db, err
		}
	}

	return db, nil
}

// Insert new activity.
func sqliteStorage_InsertActivity(Name string, Project string, Tags string, Description string) (Activity, error) {
	var a Activity
	// ...
	return a, nil
}

// Insert new project.
func sqliteStorage_InsertProject(Name string, Description string) (Project, error) {
	var p Project
	// ...
	return p, nil
}

// Remove activity(-ies) with given Id(s) form the database.
func sqliteStorage_RemoveActivity(Id ...int64) (int, error) {
	// ...
	return 0, nil
}

// Remove project(s) with given Id(s) form the database.
func sqliteStorage_RemoveProject(Id ...int64) (int, error) {
	// ...
	return 0, nil
}

// Return activities.
func sqliteStorage_SelectActivities(db *sql.DB) ([]Activity, error) {
	var a []Activity
	// ...
	return a, nil
}

// Return activity(-ies) by given ID(s).
func sqliteStorage_SelectActivityById(Id ...int64) ([]Activity, error) {
	var a []Activity
	// ...
	return a, nil
}

// Return currently running activity.
func sqliteStorage_SelectActivityRunning(db *sql.DB) (Activity, error) {
	var a Activity
	row := db.QueryRow(`SELECT * FROM Activities WHERE Stopped IS "" LIMIT 1`)
	err := row.Scan(&a.ActivityId, &a.ProjectId, &a.Name, &a.Description, &a.Tags, &a.Started, &a.Stopped)
	if err != nil {
		return a, err
	}

	return a, nil
}

// Return projects.
func sqliteStorage_SelectProjects(db *sql.DB) ([]Project, error) {
	var p []Project
	// ...
	return p, nil
}

// Return project(s) by given ID(s).
func sqliteStorage_SelectProjectById(Id ...int64) ([]Project, error) {
	var p []Project
	// ...
	return p, nil
}

// Return single project by given name(s).
func sqliteStorage_SelectProjectByName(Name ...string) ([]Project, error) {
	var p []Project
	// ...
	return p, nil
}

// ==========================================================================
// Some "internal" functions used in code above

// Returns path to the SQLite database file.
func databasePath() (string, error) {
	usr, err := user.Current()
	if err != nil {
		return "", err
	}
	return path.Join(usr.HomeDir, DefaultDatabaseName), nil
}

// Returns schema version.
func schemaVersion(db *sql.DB) (int, error) {
	var user_version int = 0
	row := db.QueryRow("PRAGMA user_version;")
	err := row.Scan(&user_version)
	if err != nil {
		return 0, err
	}
	return user_version, nil
}

// Creates database schema.
func schemaCreate(db *sql.DB) error {
	_, err := db.Exec(schemaSql)
	if err != nil {
		return err
	}
	return nil
}
