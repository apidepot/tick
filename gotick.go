// Copyright (c) 2016 The gotick developers. All rights reserved.
// Project site: https://github.com/questrail/gotick
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

package gotick

import (
	"fmt"
	"io/ioutil"
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
	apiTokenString := fmt.Sprintf("Token token=%s", tickSession.APIToken)
	req.Header.Add("Authorization", apiTokenString)
	req.Header.Add("User-Agent", tickSession.UserAgent)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln("Problem with GET request: %s", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln("Problem reading body response: %s", err)
	}
	return body, nil
}
