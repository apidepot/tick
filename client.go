// Copyright (c) 2016-2023 The tick developers. All rights reserved.
// Project site: https://github.com/apidepot/tick
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

package tick

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

	"golang.org/x/time/rate"
)

// Client models a client to consume the Tick API v2.
type Client struct {
	apiToken       string
	subscriptionID string
	userAgent      string
	baseURL        *url.URL
	httpClient     *http.Client
	rateLimiter    *rate.Limiter
}

// ClientOption applies an option to the client.
type ClientOption func(*Client)

// NewClient creates a client with the given API Token, Subscription ID, and
// UserAgent. The UserAgent must include an email address per the Tick API
// requirements.
func NewClient(token, id, agent string, opts ...ClientOption) (*Client, error) {

	url, err := url.Parse(fmt.Sprintf("https://www.tickspot.com/%s/api/v2", id))
	if err != nil {
		return nil, err
	}
	c := &Client{
		apiToken:       token,
		subscriptionID: id,
		baseURL:        url,

		// Set default values, which may be overriden by user options.
		userAgent:   agent,
		httpClient:  &http.Client{Timeout: time.Second * 60},
		rateLimiter: rate.NewLimiter(rate.Every(time.Second), 100),
	}

	// Apply options using the functional option pattern.
	for _, opt := range opts {
		opt(c)
	}

	// Validate the client
	if err := c.validate(); err != nil {
		return nil, err
	}
	return c, nil
}

// validate ensures the client has validate options.
func (c *Client) validate() error {
	// TODO(mdr): Confirm that the userAgent contains a valid email address.
	return nil
}

// WithHTTPClient configures the Tick client to use the given http client.
func WithHTTPClient(httpClient *http.Client) ClientOption {
	return func(c *Client) {
		c.httpClient = httpClient
	}
}

// WithBaseURL configures the Tick client to use the given base URL instead of
// the default base URL.
func WithBaseURL(baseURL string) ClientOption {
	return func(c *Client) {
		url, err := url.Parse(baseURL)
		if err != nil {
			c.baseURL = nil
		} else {
			c.baseURL = url
		}
	}
}

// WithRateLimiter sets the rate limiter.
func WithRateLimiter(duration time.Duration, numRequests int) ClientOption {
	return func(c *Client) {
		c.rateLimiter = rate.NewLimiter(rate.Every(duration), numRequests)
	}
}

func (c *Client) newRequest(ctx context.Context, method, path string, body interface{}) (*http.Request, error) {
	rel := &url.URL{Path: path}
	u := c.baseURL.ResolveReference(rel)
	var buf io.ReadWriter
	if body != nil {
		buf = new(bytes.Buffer)
		err := json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}
	req, err := http.NewRequestWithContext(ctx, method, u.String(), buf)
	if err != nil {
		return nil, err
	}
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	apiTokenString := fmt.Sprintf("Token token=%s", c.apiToken)
	req.Header.Set("Authorization", apiTokenString)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", c.userAgent)
	return req, nil
}

// get issues a GET to the specified path and decodes the JSON to the given
// empty interface.
func (c *Client) get(ctx context.Context, path string, v interface{}) error {
	req, err := c.newRequest(ctx, "GET", path, nil)
	if err != nil {
		return fmt.Errorf("error creating get request for path %s: %w", path, err)
	}
	resp, err := c.httpClient.Do(req)

	if err != nil {
		return fmt.Errorf("error performing c.httpClient.Do: %w", err)
	}
	defer resp.Body.Close()

	return json.NewDecoder(resp.Body).Decode(v)
}
