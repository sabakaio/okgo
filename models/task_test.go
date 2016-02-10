package models

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestJobModel(t *testing.T) {
	Convey("Job model", t, func() {
		job := NewJob("a_job", "echo ok", "")
		So(job.Once, ShouldBeTrue)
		data, err := job.Marshal()
		So(err, ShouldBeNil)

		res, err := UnmarshalJob(data)
		So(err, ShouldBeNil)
		So(res.Command, ShouldEqual, job.Command)
		So(res.Once, ShouldBeTrue)
	})

	Convey("Job model operations", t, func() {
		// TODO Use test storage, reset data before running tests

		jobName := "a_test_job"

		job, err := CreateJob(jobName, "command", "")
		So(err, ShouldBeNil)
		So(job.Name, ShouldEqual, jobName)
		So(job.Command, ShouldEqual, "command")
		So(job.Once, ShouldBeTrue)

		Convey("Get", func() {
			job, err := GetJob(jobName)
			So(err, ShouldBeNil)
			So(job.Command, ShouldEqual, "command")
			So(job.Once, ShouldBeTrue)
		})

		Convey("List", func() {
			jobs, err := ListJobs()
			So(err, ShouldBeNil)
			// Greater than 1, because we do not reset storage yet
			// TODO reset storage and assert 1 job in a list
			So(len(*jobs), ShouldBeGreaterThanOrEqualTo, 1)
		})

		Convey("Remove", func() {
			err := RemoveJob(jobName)
			So(err, ShouldBeNil)
			job, err := kv.Get("test")
			So(job, ShouldBeNil)
			So(err, ShouldNotBeNil)
		})
	})
}
