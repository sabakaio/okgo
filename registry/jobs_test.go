package registry

import (
	"../models"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestJobRegistry(t *testing.T) {
	Convey("Jobs registry", t, func() {
		reg := NewJobsRegistry()

		job := models.NewJob("a_job", "echo ok", "")
		err := reg.Put(job)
		So(err, ShouldBeNil)

		job = models.NewJob("b_job", "echo still ok", "")
		err = reg.Put(job)
		So(err, ShouldBeNil)

		Convey("Get", func() {
			res, err := reg.Get("a_job")
			So(err, ShouldBeNil)
			So(res, ShouldNotBeNil)
			So(res.Command, ShouldEqual, "echo ok")
		})

		Convey("List", func() {
			res, err := reg.List()
			So(err, ShouldBeNil)
			So(len(res), ShouldEqual, 2)
		})

		Convey("Delete", func() {
			err := reg.Delete("a_job")
			So(err, ShouldBeNil)

			job, err := reg.Get("a_job")
			So(job, ShouldBeNil)
			So(err, ShouldNotBeNil)

			job, err = reg.Get("b_job")
			So(job, ShouldNotBeNil)
			So(err, ShouldBeNil)
		})

		Convey("DeleteAll", func() {
			job := models.NewJob("c_job", "echo ok yet", "")
			err := reg.Put(job)
			So(err, ShouldBeNil)

			err = reg.DeleteAll()
			So(err, ShouldBeNil)

			job, err = reg.Get("b_job")
			So(job, ShouldBeNil)
			So(err, ShouldNotBeNil)

			job, err = reg.Get("c_job")
			So(job, ShouldBeNil)
			So(err, ShouldNotBeNil)
		})
	})
}
