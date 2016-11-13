package reflekt

import (
	"encoding/json"
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

var testsTestValue_JSON = []struct {
	from []byte
	get  func(v *Value) interface{}
	to   interface{}
}{
	{
		from: []byte(`123`),
		get: func(v *Value) interface{} {
			return v.Int()
		},
		to: 123,
	},
	{
		from: []byte(`"xxx"`),
		get: func(v *Value) interface{} {
			return v.String()
		},
		to: "xxx",
	},
	{
		from: []byte(`{"foo":1}`),
		get: func(v *Value) interface{} {
			return v.InterfaceMap()
		},
		to: map[string]interface{}{
			"foo": float64(1),
		},
	},
	{
		from: []byte(`{"foo":"bar"}`),
		get: func(v *Value) interface{} {
			return v.StringMap()
		},
		to: map[string]string{
			"foo": "bar",
		},
	},
	{
		from: []byte(`[1,2,3]`),
		get: func(v *Value) interface{} {
			return v.Ints()
		},
		to: []int{1, 2, 3},
	},
}

func TestValue_JSON(t *testing.T) {
	Convey("Marshal and unmarshal JSON", t, func() {
		for i, test := range testsTestValue_JSON {
			out := fmt.Sprintf("(%d) From \"%v\"", i+1, test.from)
			Convey(out, func() {
				to := NewValue(nil)
				err := json.Unmarshal(test.from, to)
				So(err, ShouldBeNil)
				So(test.get(to), ShouldResemble, test.to)
				raw, err := json.Marshal(to)
				So(string(raw), ShouldEqual, string(test.from))
			})
		}
	})
}
