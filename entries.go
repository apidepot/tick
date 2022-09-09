// Copyright (c) 2016-2022 The tick developers. All rights reserved.
// Project site: https://github.com/apidepot/tick
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

package tick

import (
	"context"
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

func (c Client) GetEntry(ctx context.Context, entryID int) (Entry, error) {
	var entry Entry
	path := fmt.Sprintf("/entries/%d.json", entryID)
	err := c.get(ctx, path, &entry)
	return entry, err
}

type EntryOptions struct {
}

func (c Client) GetEntries(ctx context.Context, options EntryOptions) (Entries, error) {
	return nil, nil
}

func (c Client) getEntriesBetweenDates(ctx context.Context, startDate, endDate string) (Entries, error) {
	var allEntries Entries
	foundLastPage := false
	currentPage := 1
	for !foundLastPage {
		entries, err := c.getEntriesBetweenDatesOnPage(ctx, startDate, endDate, currentPage)
		if err != nil {
			return nil, err
		}
		if entries == nil {
			foundLastPage = true
		} else {
			allEntries = append(allEntries, entries...)
			currentPage++
		}
	}
	return allEntries, nil
}

func (c Client) getEntriesBetweenDatesOnPage(ctx context.Context, startDate, endDate string, page int) (Entries, error) {
	var entries Entries
	// /entries.json?start_date=2014-12-28&end_date=2015-01-24&page=15
	path := fmt.Sprintf(
		"/entries.json?page=%d&start_date='%s'&end_date='%s'",
		page,
		startDate,
		endDate,
	)
	err := c.get(ctx, path, &entries)
	return entries, err
}

func (c Client) getProjectEntriesBetweenDates(
	ctx context.Context, projectID int, startDate, endDate string,
) (Entries, error) {
	var allEntries Entries
	foundLastPage := false
	currentPage := 1
	for !foundLastPage {
		entries, err := c.getProjectEntriesBetweenDatesOnPage(
			ctx, projectID, startDate, endDate, currentPage)
		if err != nil {
			return nil, err
		}
		if entries == nil {
			foundLastPage = true
		} else {
			allEntries = append(allEntries, entries...)
			currentPage++
		}
	}
	return allEntries, nil
}

func (c Client) getProjectEntriesBetweenDatesOnPage(
	ctx context.Context, projectID int, startDate, endDate string, page int,
) (Entries, error) {
	var entries Entries
	// /entries.json?start_date=2014-12-28&end_date=2015-01-24&page=15
	path := fmt.Sprintf(
		"/projects/%d/entries.json?page=%d&start_date='%s'&end_date='%s'",
		projectID,
		page,
		startDate,
		endDate,
	)
	err := c.get(ctx, path, &entries)
	return entries, err
}
