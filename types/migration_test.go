package types

import (
	"testing"

	"sort"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

func TestMigration(t *testing.T) {
	Convey("Should compare migrations", t, func() {
		tpl := Migration{
			Number:     1,
			Name:       "test",
			UpScript:   "SELECT 1",
			DownScript: "SELECT 2",
		}
		a := tpl
		b := tpl
		So(a.Compare(&b), ShouldBeTrue)
		now := time.Now()
		b.AppliedAt = &now
		So(a.Compare(&b), ShouldBeTrue)
		b = tpl
		b.Number = 10
		So(a.Compare(&b), ShouldBeFalse)

		b = tpl
		b.Name = "test 1"
		So(a.Compare(&b), ShouldBeFalse)

		b = tpl
		b.UpScript = "SELECT 2"
		So(a.Compare(&b), ShouldBeFalse)

		b = tpl
		b.DownScript = "SELECT 1"
		So(a.Compare(&b), ShouldBeFalse)

	})
	Convey("Should test Migration.Compare", t, func() {
		a := []Migration{
			{Number: 3, Name: "3"},
			{Number: 1, Name: "1"},
			{Number: 2, Name: "2"},
		}
		sort.Sort(ByNumber(a))
		So(a[0].Name, ShouldEqual, "1")
		So(a[1].Name, ShouldEqual, "2")
		So(a[2].Name, ShouldEqual, "3")
	})
	Convey("Should merge Migrations AppliedAt", t, func() {
		a := []Migration{
			{Number: 3, Name: "3"},
			{Number: 1, Name: "1"},
			{Number: 2, Name: "2"},
		}
		oneTime := time.Now()
		threeTime := time.Now()
		b := []Migration{
			{Number: 3, Name: "3", AppliedAt: &oneTime},
			{Number: 1, Name: "1"},
			{Number: 2, Name: "2", AppliedAt: &threeTime},
		}
		c := MergeMigrationsAppliedAt(a, b)
		for _, m := range c {
			switch m.Number {
			case 1:
				So(m.AppliedAt, ShouldBeNil)
			case 2:
				So(m.AppliedAt.Equal(threeTime), ShouldBeTrue)
			case 3:
				So(m.AppliedAt.Equal(oneTime), ShouldBeTrue)
			}
		}
		sort.Sort(ByNumber(a))
		So(a[0].Name, ShouldEqual, "1")
		So(a[1].Name, ShouldEqual, "2")
		So(a[2].Name, ShouldEqual, "3")
	})
}
