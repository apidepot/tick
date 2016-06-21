// Copyright (c) 2016 The gotick developers. All rights reserved.
// Project site: https://github.com/questrail/gotick
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

package gotick

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type Entry struct {
	ID        uint     `json:"id"`
	Date      TickDate `json:"date"`
	Hours     float64  `json:"hours"`
	Notes     string   `json:"notes"`
	TaskID    uint     `json:"task_id"`
	UserID    uint     `json:"user_id"`
	URL       string   `json:"url"`
	CreatedAt string   `json:"created_at"`
	UpdatedAt string   `json:"updated_at"`
}

type Entries []Entry

func GetEntriesBetweenDates(tickData JSONGetter, startDate, endDate string) (Entries, error) {
	var allEntries Entries
	foundLastPage := false
	currentPage := 1
	for !foundLastPage {
		entries, err := GetEntriesBetweenDatesOnPage(tickData, startDate, endDate, currentPage)
		if err != nil {
			return nil, err
		}
		if entries == Entries(nil) {
			foundLastPage = true
		} else {
			allEntries = append(allEntries, entries...)
			currentPage++
		}
	}
	return allEntries, nil
}

func GetEntriesBetweenDatesOnPage(tickData JSONGetter, startDate, endDate string, page int) (Entries, error) {
	var entries Entries
	// /entries.json?start_date=2014-12-28&end_date=2015-01-24&page=15
	url := fmt.Sprintf(
		"/entries.json?page=%d&start_date='%s'&end_date='%s'",
		page,
		startDate,
		endDate,
	)
	data, err := tickData.GetJSON(url)
	if err != nil {
		return nil, err
	}
	if bytes.Equal(data, []byte("[]")) {
		return nil, nil
	}
	err = json.Unmarshal(data, &entries)
	if err != nil {
		return nil, err
	}
	return entries, nil
}

func GetProjectEntriesBetweenDates(
	tickData JSONGetter, projectID uint, startDate, endDate string,
) (Entries, error) {
	var allEntries Entries
	foundLastPage := false
	currentPage := 1
	for !foundLastPage {
		entries, err := GetProjectEntriesBetweenDatesOnPage(
			tickData, projectID, startDate, endDate, currentPage,
		)
		if err != nil {
			return nil, err
		}
		if entries == Entries(nil) {
			foundLastPage = true
		} else {
			allEntries = append(allEntries, entries...)
			currentPage++
		}
	}
	return allEntries, nil
}

func GetProjectEntriesBetweenDatesOnPage(
	tickData JSONGetter, projectID uint, startDate, endDate string, page int,
) (Entries, error) {
	var entries Entries
	// /entries.json?start_date=2014-12-28&end_date=2015-01-24&page=15
	url := fmt.Sprintf(
		"/projects/%d/entries.json?page=%d&start_date='%s'&end_date='%s'",
		projectID,
		page,
		startDate,
		endDate,
	)
	data, err := tickData.GetJSON(url)
	if err != nil {
		return nil, err
	}
	if bytes.Equal(data, []byte("[]")) {
		return nil, nil
	}
	err = json.Unmarshal(data, &entries)
	if err != nil {
		return nil, err
	}
	return entries, nil
}
