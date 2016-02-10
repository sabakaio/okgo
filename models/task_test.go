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
}
