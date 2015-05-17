// Copyright 2015 Ondrej Donek. All rights reserved.
// See LICENSE file for more informations about licensing.

package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os/user"
	"path"
	"strconv"
	"strings"
	"time"
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
	UpdateActivity:        sqliteStorage_UpdateActivity,
	UpdateProject:         sqliteStorage_UpdateProject,
}

// Initialize storage.
func sqliteStorage_Init() (db *sql.DB, err error) {
	dbPath, err := databasePath()
	//log.Println(dbPath)
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
	//log.Printf("Current database schema version: %d\n", ver)
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
func sqliteStorage_InsertActivity(db *sql.DB, pid int64, name string, desc string, tags string) (a Activity, err error) {
	sqlStmt := `
	INSERT INTO Activities 
	(ProjectId, Name, Description, Tags, Started) 
	VALUES (?, ?, ?, ?, ?)
	`
	stmt, err := db.Prepare(sqlStmt)
	if err != nil {
		log.Fatal(err)
	}

	started := time.Now().Format(time.RFC3339)
	res, err := stmt.Exec(pid, name, desc, tags, started)
	if err != nil {
		log.Fatal(err)
	}

	aid, err := res.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}

	a.ActivityId = aid
	a.ProjectId = pid
	a.Name = name
	a.Description = desc
	a.Tags = tags
	a.Started = started

	return a, nil
}

// Insert new project.
func sqliteStorage_InsertProject(db *sql.DB, name string, desc string) (p Project, err error) {
	sqlStmt := `
	INSERT INTO Projects 
	(Name, Description, Created) 
	VALUES (?, ?, ?)
	`
	stmt, err := db.Prepare(sqlStmt)
	if err != nil {
		log.Fatal(err)
	}

	created := time.Now().Format(time.RFC3339)
	res, err := stmt.Exec(name, desc, created)
	if err != nil {
		log.Fatal(err)
	}

	pid, err := res.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}

	p.ProjectId = pid
	p.Name = name
	p.Description = desc
	p.Created = created

	return p, nil
}

// Remove activity(-ies) with given Id(s) form the database.
func sqliteStorage_RemoveActivity(db *sql.DB, id ...int64) (int, error) {
	// ...
	return 0, nil
}

// Remove project(s) with given Id(s) form the database.
func sqliteStorage_RemoveProject(db *sql.DB, id ...int64) (int, error) {
	// ...
	return 0, nil
}

// Return activities.
func sqliteStorage_SelectActivities(db *sql.DB, limit int) (activities []Activity, err error) {
	stmtSql := `SELECT * FROM Activities ORDER BY Started DESC LIMIT ?`
	stmt, err := db.Prepare(stmtSql)
	if err != nil {
		return activities, err
	}

	rows, err := stmt.Query(limit)
	if err != nil {
		return activities, err
	}

	defer rows.Close()
	for rows.Next() {
		var a Activity
		rows.Scan(&a.ActivityId, &a.ProjectId, &a.Name, &a.Description, &a.Tags, &a.Started, &a.Stopped)
		activities = append(activities, a)
	}
	rows.Close()

	return activities, nil
}

// Return activity(-ies) by given ID(s).
func sqliteStorage_SelectActivityById(db *sql.DB, id ...int64) (a []Activity, err error) {
	// ...
	return a, nil
}

// Return currently running activity.
func sqliteStorage_SelectActivityRunning(db *sql.DB) (a Activity, err error) {
	row := db.QueryRow(`SELECT * FROM Activities WHERE Stopped IS "" LIMIT 1`)
	err = row.Scan(&a.ActivityId, &a.ProjectId, &a.Name, &a.Description, &a.Tags, &a.Started, &a.Stopped)
	if err != nil {
		return a, err
	}

	return a, nil
}

// Return projects.
func sqliteStorage_SelectProjects(db *sql.DB, limit int) (p []Project, err error) {
	stmtSql := `SELECT * FROM Projects ORDER BY Name ASC LIMIT ?`
	stmt, err := db.Prepare(stmtSql)
	if err != nil {
		return p, err
	}

	rows, err := stmt.Query(limit)
	if err != nil {
		return p, err
	}

	defer rows.Close()
	return parseProjectsFromRows(rows)
}

// Return project(s) by given ID(s).
func sqliteStorage_SelectProjectById(db *sql.DB, id ...int64) (p []Project, err error) {
	var ids []string
	for _, id := range id {
		ids = append(ids, strconv.FormatInt(id, 10))
	}
	idsStr := strings.Join(ids, ", ")

	sqlStmt := "SELECT * FROM Projects WHERE Id IN (" + idsStr + ")"
	rows, err := db.Query(sqlStmt)
	if err != nil {
		return p, err
	}

	defer rows.Close()
	return parseProjectsFromRows(rows)
}

// Return single project by given name(s).
func sqliteStorage_SelectProjectByName(db *sql.DB, name ...string) (projects []Project, err error) {
	namesStr := strings.Join(name, "\", \"")
	namesStr = "\"" + namesStr + "\""

	sqlStmt := "SELECT * FROM Projects WHERE Name IN (" + namesStr + ")"
	rows, err := db.Query(sqlStmt)
	if err != nil {
		return projects, err
	}

	defer rows.Close()
	return parseProjectsFromRows(rows)
}

// Update activity in the database
// Return 1 if update was successfull otherwise 0.
func sqliteStorage_UpdateActivity(db *sql.DB, a Activity) (cnt int64, err error) {
	sqlStmt := `
	UPDATE Activities 
	SET
	ProjectId = ?, 
	Name = ?, 
	Description = ?, 
	Tags = ?, 
	Started = ?, 
	Stopped = ? 
	WHERE ActivityId = ? 
	`
	stmt, err := db.Prepare(sqlStmt)
	if err != nil {
		log.Fatal(err)
	}

	res, err := stmt.Exec(a.ProjectId, a.Name, a.Description, a.Tags,
		a.Started, a.Stopped, a.ActivityId)
	if err != nil {
		log.Fatal(err)
	}

	return res.RowsAffected()
}

// Update project in the database.
// Return 1 if update was successfull otherwise 0.
func sqliteStorage_UpdateProject(db *sql.DB, p Project) (cnt int64, err error) {
	sqlStmt := `
	UPDATE Projects 
	SET 
	Name = ?, 
	Description = ?, 
	Created = ? 
	WHERE ProjectId = ? 
	`
	stmt, err := db.Prepare(sqlStmt)
	if err != nil {
		log.Fatal(err)
	}

	res, err := stmt.Exec(p.Name, p.Description, p.Created, p.ProjectId)
	if err != nil {
		log.Fatal(err)
	}

	return res.RowsAffected()
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

// Helper method that converting results rows
// into regular instances of Project object.
func parseProjectsFromRows(rows *sql.Rows) (projects []Project, err error) {
	for rows.Next() {
		var p Project
		rows.Scan(&p.ProjectId, &p.Name, &p.Description, &p.Created)
		projects = append(projects, p)
	}
	rows.Close()

	return projects, nil
}
