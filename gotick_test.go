package gotick

import (
	"testing"

	c "github.com/smartystreets/goconvey/convey"
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

	tasksPage1 = `[
  {
    "id": 6464641,
    "name": "Meetings",
    "budget": null,
    "position": 1,
    "project_id": 1111111,
    "date_closed": null,
    "billable": true,
    "url": "https://www.tickspot.com/22222/api/v2/tasks/6464641.json",
    "created_at": "2014-09-18T12:59:09.000-04:00",
    "updated_at": "2014-09-18T12:59:09.000-04:00"
  },
  {
    "id": 6464642,
    "name": "Design",
    "budget": null,
    "position": 18,
    "project_id": 1111111,
    "date_closed": null,
    "billable": true,
    "url": "https://www.tickspot.com/22222/api/v2/tasks/6464642.json",
    "created_at": "2014-08-26T18:39:23.000-04:00",
    "updated_at": "2014-08-26T18:39:23.000-04:00"
  },
  {
    "id": 6464643,
    "name": "Test",
    "budget": null,
    "position": 18,
    "project_id": 1111111,
    "date_closed": null,
    "billable": true,
    "url": "https://www.tickspot.com/22222/api/v2/tasks/6464643.json",
    "created_at": "2015-09-29T13:29:28.000-04:00",
    "updated_at": "2015-09-29T13:29:28.000-04:00"
  }
]`

	singleTask = `{
  "id": 25,
  "name": "Design",
  "budget": null,
  "position": 7,
  "project_id": 1111111,
  "date_closed": "2015-10-08",
  "billable": true,
  "created_at": "2015-09-21T11:56:20.000-04:00",
  "updated_at": "2015-10-08T14:18:00.000-04:00",
  "total_hours": 36.0,
  "entries": {
    "count": 7,
    "url": "https://www.tickspot.com/22222/api/v2/tasks/25/entries.json",
    "updated_at": "2015-10-02T12:04:17.000-04:00"
  },
  "project": {
    "id": 1111111,
    "name": "P.001 First Project",
    "budget": null,
    "date_closed": null,
    "notifications": false,
    "billable": true,
    "recurring": false,
    "client_id": 222244,
    "owner_id": 222243,
    "url": "https://www.tickspot.com/22222/api/v2/projects/1111111.json",
    "created_at": "2015-05-18T10:51:45.000-04:00",
    "updated_at": "2015-05-18T10:51:45.000-04:00"
  }
}`
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
	case "/tasks.json":
		data = []byte(tasksPage1)
	case "/tasks.json?page=1":
		data = []byte(tasksPage1)
	case "/tasks/25.json":
		data = []byte(singleTask)
	default:
		data = []byte("[]")
	}
	return data, nil
}

func TestCreatingNewTickSession(t *testing.T) {
	c.Convey("Given I want to create a new Tick session", t, func() {
		c.Convey("When the NewTickSession is called with valid variables", func() {
			tick, _ := NewTickSession("mytoken", "subID", "thisUserAgent")
			c.Convey("The session should contain the APIToken", func() {
				c.So(tick.APIToken, c.ShouldEqual, "mytoken")
			})
			c.Convey("The session should contain the SubscriptionID", func() {
				c.So(tick.SubscriptionID, c.ShouldEqual, "subID")
			})
			c.Convey("The session should contain the UserAgent", func() {
				c.So(tick.UserAgent, c.ShouldEqual, "thisUserAgent")
			})
		})
	})
}
