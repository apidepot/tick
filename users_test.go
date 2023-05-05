// Copyright (c) 2016-2022 The tick developers. All rights reserved.
// Project site: https://github.com/apidepot/tick
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

package tick

import (
	"io"
	"net/http"
	"strings"
	"testing"
)

type fakeService func(req *http.Request) (*http.Response, error)

func (f fakeService) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req)
}

func TestGetUsers(t *testing.T) {

	client := &http.Client{
		Transport: fakeService(func(req *http.Request) (*http.Response, error) {
			// Test request parameters
			assertString(t, "url", req.URL.String(), "https://www.tickspot.com/1234/api/v2/users.json")

			return &http.Response{
				StatusCode: http.StatusOK,
				Header: http.Header{
					"Content-Type": []string{"application/json"},
				},
				Body: io.NopCloser(strings.NewReader(`{"foo":"bar"}`)),
			}, nil
		}),
	}
	_, err := NewClient("mytoken", "subID", "thusUserAgent", WithHTTPClient(client))
	if err != nil {
		t.Fatalf("Error creating client for user testing: %s", err)
	}

}
