package gotick

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
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

type FakeSession struct {
}

func (fake *FakeSession) GetJSON(url string) ([]byte, error) {
	var data []byte
	switch url {
	case "/projects.json":
		data = []byte(projectsPage1)
	case "/projects.json?page=1":
		data = []byte(projectsPage1)
	case "/projects.json?page=2":
		data = []byte(projectsPage2)
	case "/users.json":
		data = []byte(usersPage1)
	default:
		data = []byte("[]")
	}
	return data, nil
}

func TestCreatingNewTickSession(t *testing.T) {
	Convey("Given I want to create a new Tick session", t, func() {
		Convey("When the NewTickSession is called with valid variables", func() {
			tick, _ := NewTickSession("mytoken", "subID", "thisUserAgent")
			Convey("The session should contain the APIToken", func() {
				So(tick.APIToken, ShouldEqual, "mytoken")
			})
			Convey("The session should contain the SubscriptionID", func() {
				So(tick.SubscriptionID, ShouldEqual, "subID")
			})
			Convey("The session should contain the UserAgent", func() {
				So(tick.UserAgent, ShouldEqual, "thisUserAgent")
			})
		})
	})
}
