// Copyright 2015 Ondrej Donek. All rights reserved.
// See LICENSE file for more informations about licensing.

package main

import (
	"errors"
	"time"
)

// Project definition
type Project struct {
	ProjectId int64 // Numeric identifier of the project.
	Name string // Name of the project.
	Description string // Description of the project.
	Created string // Datetime of creation of the project (formatted by RFC3339).
}

// Returns `Created` string as regular instance of `time.Time`.
func (p *Project) CreatedTime() (time.Time, error) {
	return time.Parse(p.Created, time.RFC3339)
}

// Activity definition
type Activity struct {
	ActivityId int64
	ProjectId int64
	Name string
	Description string
	Tags string
	Started string
	Stopped string
}

// Returns `Started` string as regular instance of `time.Time`.
func (a *Activity) StartedTime() (time.Time, error) {
	return time.Parse(a.Started, time.RFC3339)
}

// Returns `Stopped` string as regular instance of `time.Time`.
func (a *Activity) StoppedTime() (time.Time, error) {
	return time.Parse(a.Stopped, time.RFC3339)
}

// Initialize activity from the given string.
func (a *Activity) Parse(activityString string) error {
	if activityString == "" {
		return errors.New("Empty string given!")
	}
	
	return errors.New("Implement `Activity.Parse()`!")
}