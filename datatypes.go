// Copyright 2015 Ondrej Donek. All rights reserved.
// See LICENSE file for more informations about licensing.

package main

import (
	"database/sql"
	"errors"
	"log"
	"strings"
	"time"
)

const (
	ProjectNameSep = "@"
	TagsSep        = ";"
	DescriptionSep = "#"
)

// Project definition
type Project struct {
	ProjectId   int64  // Numeric identifier of the project.
	Name        string // Name of the project.
	Description string // Description of the project.
	Created     string // Datetime of creation of the project (formatted by RFC3339).
}

// Returns `Created` string as regular instance of `time.Time`.
func (p *Project) CreatedTime() (time.Time, error) {
	return time.Parse(p.Created, time.RFC3339)
}

// Activity definition
type Activity struct {
	ActivityId  int64
	ProjectId   int64
	Name        string
	Description string
	Tags        string
	Started     string
	Stopped     string
	project     Project
}

func (a *Activity) GetProject() Project {
	return a.project
}

func (a *Activity) SetProject(p Project) {
	a.project = p
}

// Returns `Started` string as regular instance of `time.Time`.
func (a *Activity) StartedTime() (time.Time, error) {
	return time.Parse(time.RFC3339, a.Started)
}

// Returns `Stopped` string as regular instance of `time.Time`.
func (a *Activity) StoppedTime() (time.Time, error) {
	return time.Parse(time.RFC3339, a.Stopped)
}

// Returns duration of the activity.
func (a *Activity) Duration() time.Duration {
	stopped, err := a.StoppedTime()
	if err != nil {
		stopped = time.Now()
	}

	started, err := a.StartedTime()
	if err != nil {
		started = time.Now()
	}

	return stopped.Sub(started)
}

// Initialize activity from the given string.
func (a *Activity) Parse(activityString string, db *sql.DB) error {
	aStr := strings.Trim(activityString, " \n\t")
	if aStr == "" {
		return errors.New("Empty string given!")
	}

	if strings.Count(aStr, ProjectNameSep) > 1 || strings.Count(aStr, TagsSep) > 1 || strings.Count(aStr, DescriptionSep) > 1 {
		return errors.New("Given activity string is not well formed!")
	}

	hasProjectName := strings.Contains(aStr, ProjectNameSep)
	hasTags := strings.Contains(aStr, TagsSep)
	hasDescription := strings.Contains(aStr, DescriptionSep)
	projectName := ""

	if hasDescription == true {
		parts := strings.Split(aStr, DescriptionSep)
		aStr = parts[0]
		a.Description = parts[1]
	}

	if hasTags == true {
		parts := strings.Split(aStr, TagsSep)
		aStr = parts[0]
		a.Tags = parts[1]
	}

	if hasProjectName == true {
		parts := strings.Split(aStr, ProjectNameSep)
		aStr = parts[0]
		projectName = parts[1]
	}

	a.Name = aStr

	if hasProjectName == true && projectName != "" {
		projects, err := SqliteStorage.SelectProjectByName(db, projectName)
		if err != nil {
			log.Fatal(err)
		}

		if len(projects) == 1 {
			a.ProjectId = projects[0].ProjectId
			a.SetProject(projects[0])
		} else if len(projects) == 0 {
			p, err := SqliteStorage.InsertProject(db, projectName, "")
			if err != nil {
				log.Fatal(err)
			}
			a.ProjectId = p.ProjectId
		}
	}

	return nil
}
