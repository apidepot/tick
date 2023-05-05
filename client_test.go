// Copyright (c) 2016-2023 The tick developers. All rights reserved.
// Project site: https://github.com/apidepot/tick
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

package tick

import (
	"fmt"
	"testing"
)

const (
	projectsPage1 = `[
	{
		"id":1111111,
		"name":"P.001 First Project",
		"budget":null,
		"date_closed":null,
		"notifications":false,
		"billable":true,
		"recurring":false,
		"client_id":555555,
		"owner_id":444444,
		"url":"https://www.tickspot.com/22222/api/v2/projects/1111111.json",
		"created_at":"2014-08-06T09:39:44.000-04:00",
		"updated_at":"2014-08-26T18:19:31.000-04:00"
	},
	{
		"id":1111112,
		"name":"P.002 Second Project",
		"budget":null,
		"date_closed":null,
		"notifications":false,
		"billable":true,
		"recurring":false,
		"client_id":555555,
		"owner_id":444444,
		"url":"https://www.tickspot.com/22222/api/v2/projects/1111112.json",
		"created_at":"2014-08-07T17:43:39.000-04:00",
		"updated_at":"2014-08-13T12:22:35.000-04:00"
	}
]`
	projectsPage2 = `[
	{
		"id":1111113,
		"name":"P.003 Third Project",
		"budget":null,
		"date_closed":null,
		"notifications":false,
		"billable":true,
		"recurring":false,
		"client_id":555555,
		"owner_id":444444,
		"url":"https://www.tickspot.com/22222/api/v2/projects/1111113.json",
		"created_at":"2014-08-06T09:39:44.000-04:00",
		"updated_at":"2014-08-26T18:19:31.000-04:00"
	},
	{
		"id":1111114,
		"name":"P.004 Fourth Project",
		"budget":null,
		"date_closed":null,
		"notifications":false,
		"billable":true,
		"recurring":false,
		"client_id":555555,
		"owner_id":444444,
		"url":"https://www.tickspot.com/22222/api/v2/projects/1111114.json",
		"created_at":"2014-08-07T17:43:39.000-04:00",
		"updated_at":"2014-08-13T12:22:35.000-04:00"
	}
]`

	usersPage1 = `[
  {
    "id": 444444,
    "first_name": "Adam",
    "last_name": "Smith",
    "email": "adam.smith@example.com",
    "timezone": "Central Time (US & Canada)",
    "created_at": "2014-08-28T13:03:28.000-04:00",
    "updated_at": "2014-09-02T13:54:07.000-04:00"
  },
  {
    "id": 444445,
    "first_name": "Thomas",
    "last_name": "Sowell",
    "email": "thomas.sowell@example.com",
    "timezone": "Central Time (US & Canada)",
    "created_at": "2015-05-18T13:37:49.000-04:00",
    "updated_at": "2015-05-22T09:58:05.000-04:00"
  },
  {
    "id": 444446,
    "first_name": "Friedrich",
    "last_name": "von Hayek",
    "email": "friedrich.von.hayek@example.com",
    "timezone": "Central Time (US & Canada)",
    "created_at": "2014-08-28T13:04:26.000-04:00",
    "updated_at": "2014-08-28T13:04:26.000-04:00"
  },
  {
    "id": 444447,
    "first_name": "Benjamin",
    "last_name": "Graham",
    "email": "benjamin.graham@example.com",
    "timezone": "Central Time (US & Canada)",
    "created_at": "2014-08-28T13:03:28.000-04:00",
    "updated_at": "2014-09-02T13:54:07.000-04:00"
  },
  {
    "id": 444448,
    "first_name": "Milton",
    "last_name": "Friedman",
    "email": "milton.friedman@example.com",
    "timezone": "Central Time (US & Canada)",
    "created_at": "2015-05-18T13:37:49.000-04:00",
    "updated_at": "2015-05-22T09:58:05.000-04:00"
  }
]`
)

func TestCreatingNewTickSession(t *testing.T) {
	testCases := []struct {
		token string
		subID string
		agent string
	}{
		{"mytoken", "subID", "thisUserAgent"},
	}
	for i, tc := range testCases {
		name := fmt.Sprintf("new_session_%d", i)
		t.Run(name, func(t *testing.T) {
			tick, _ := NewClient("mytoken", "subID", "thisUserAgent")
			if tick.apiToken != tc.token {
				assertString(t, "token", tick.apiToken, tc.token)
			}
			if tick.subscriptionID != tc.subID {
				assertString(t, "subscriptionID", tick.subscriptionID, tc.subID)
			}
			if tick.userAgent != tc.agent {
				assertString(t, "userAgent", tick.userAgent, tc.agent)
			}
		})
	}
}
