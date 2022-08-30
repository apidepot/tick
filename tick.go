// Copyright (c) 2016-2022 The tick developers. All rights reserved.
// Project site: https://github.com/apidepot/tick
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

package gotick

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

type JSONGetter interface {
	GetJSON(string) ([]byte, error)
}

type TickSession struct {
	APIToken       string
	SubscriptionID string
	UserAgent      string
}

func NewTickSession(apiToken, subscriptionID, userAgent string) (*TickSession, error) {
	tickSession := TickSession{
		APIToken:       apiToken,
		SubscriptionID: subscriptionID,
		UserAgent:      userAgent,
	}
	return &tickSession, nil
}

func (tickSession *TickSession) GetJSON(url string) ([]byte, error) {
	client := &http.Client{}
	fullURL := fmt.Sprintf(
		"https://www.tickspot.com/%s/api/v2%s",
		tickSession.SubscriptionID,
		url,
	)
	log.Printf("URL: %s", fullURL)
	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return nil, err
	}
	apiTokenString := fmt.Sprintf("Token token=%s", tickSession.APIToken)
	req.Header.Add("Authorization", apiTokenString)
	req.Header.Add("User-Agent", tickSession.UserAgent)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return io.ReadAll(resp.Body)
}
