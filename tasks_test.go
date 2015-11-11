package gotick

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestGetOpenTasksOnFirstPage(t *testing.T) {
	Convey("Given I want to get the first page of open tasks", t, func() {
		Convey("When the /tasks.json?page=1 URI receives a GET method\n"+
			"And the header contains valid authorization",
			func() {
				var fake FakeSession
				data, _ := GetOpenTasksOnPage(&fake, 1)
				Convey("It should return three tasks", func() {
					So(len(data), ShouldEqual, 3)
				})
				taskNames := []string{"Meetings", "Design", "Test"}
				ordinals := []string{"first", "second", "third"}
				for index, taskName := range taskNames {
					conveyance := fmt.Sprintf("The %s task name should be %s.", ordinals[index], taskName)
					Convey(conveyance, func() {
						So(data[index].Name, ShouldEqual, taskName)
					})

				}
			})
	})
}

func TestGetOpenTasksOnSecondPage(t *testing.T) {
	Convey("Given I want to get the third page of open tasks", t, func() {
		Convey("When the /tasks.json?page=3 URI receives a GET method\n"+
			"And the header contains valid authorization",
			func() {
				var fake FakeSession
				data, _ := GetOpenTasksOnPage(&fake, 3)
				Convey("It should return no tasks", func() {
					expected := Tasks(nil)
					So(data, ShouldResemble, expected)
				})
			})
	})
}

func TestGetOpenTasks(t *testing.T) {
	Convey("Given I want to get the open tasks", t, func() {
		Convey("When the /tasks.json receives a GET method\n"+
			"And the header contains valid authorization",
			func() {
				var fake FakeSession
				data, _ := GetOpenTasks(&fake)
				Convey("It should return four open tasks", func() {
					So(len(data), ShouldEqual, 3)
				})
			})
	})
}

func TestGetSpecificTask(t *testing.T) {
	Convey("Given I want to get a specific task", t, func() {
		Convey("When /tasks/25.json receives a GET method\n"+
			"And the header contains valid authorization",
			func() {
				var fake FakeSession
				data, _ := GetTask(&fake, 25)
				Convey("It should return task with ID 25", func() {
					So(data.ID, ShouldEqual, 25)
				})
				Convey("Task 25 should be named Design", func() {
					So(data.Name, ShouldEqual, "Design")
				})
				Convey("Task 25 should have a project name as well", func() {
					So(data.Project.Name, ShouldEqual, "P.001 First Project")
				})
			})
	})
}
