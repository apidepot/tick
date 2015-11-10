package gotick

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestGetOpenProjectsOnFirstPage(t *testing.T) {
	Convey("Given I want to get the first page of open projects", t, func() {
		Convey("When the /projects.json?page=1 URI receives a GET method\n"+
			"And the header contains valid authorization",
			func() {
				var fake FakeSession
				data, _ := GetOpenProjectsOnPage(&fake, 1)
				Convey("It should return two projects", func() {
					So(len(data), ShouldEqual, 2)
				})
				Convey("The first project ID should be 1111111", func() {
					So(data[0].ID, ShouldEqual, 1111111)
				})
				Convey("The first project name should be P.001 First Project", func() {
					So(data[0].Name, ShouldEqual, "P.001 First Project")
				})
				Convey("The second project ID should be 1111111", func() {
					So(data[1].ID, ShouldEqual, 1111112)
				})
				Convey("The second project name should be P.001 First Project", func() {
					So(data[1].Name, ShouldEqual, "P.002 Second Project")
				})
			})
	})
}

func TestGetOpenProjectsOnSecondPage(t *testing.T) {
	Convey("Given I want to get the third page of open projects", t, func() {
		Convey("When the /projects.json?page=3 URI receives a GET method\n"+
			"And the header contains valid authorization",
			func() {
				var fake FakeSession
				data, _ := GetOpenProjectsOnPage(&fake, 3)
				Convey("It should return no projects", func() {
					expected := Projects(nil)
					So(data, ShouldResemble, expected)
				})
			})
	})
}

func TestGetOpenProjects(t *testing.T) {
	Convey("Given I want to get the open projects", t, func() {
		Convey("When the /projects.json receives a GET method\n"+
			"And the header contains valid authorization",
			func() {
				var fake FakeSession
				data, _ := GetOpenProjects(&fake)
				Convey("It should return four open projects", func() {
					So(len(data), ShouldEqual, 4)
				})
			})
	})
}
