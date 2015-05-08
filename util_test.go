package reflekt
import (
	"testing"
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
)

var testSnakeCases = []struct {
	from string
	to   string
}{
	{
		from: "foo",
		to:   "foo",
	},
	{
		from: "F",
		to:   "f",
	},
	{
		from: "FF",
		to:   "ff",
	},
	{
		from: "Foo",
		to:   "foo",
	},
	{
		from: "fooBar",
		to:   "foo_bar",
	},
	{
		from: "FooBar",
		to:   "foo_bar",
	},
	{
		from: "fooBarBaz",
		to:   "foo_bar_baz",
	},
}

func TestSnakeCase(t *testing.T) {
	Convey("Case snaking", t, func() {
		for _, test := range testSnakeCases {
			Convey(fmt.Sprintf("Snaking %s -> %s", test.from, test.to), func() {
				to := snakeCase(test.from)
				So(to, ShouldEqual, test.to)
			})
		}
	})
}