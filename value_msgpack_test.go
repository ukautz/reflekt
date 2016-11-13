package reflekt

import (
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
	"gopkg.in/vmihailenco/msgpack.v2"
	"github.com/davecgh/go-spew/spew"
)

var testsTestValue_Msgpack = []struct {
	from interface{}
	get  func(v *Value) interface{}
	to   interface{}
}{
	{
		from: 123,
		get: func(v *Value) interface{} {
			return v.Int()
		},
		to: 123,
	},
	{
		from: `xxx`,
		get: func(v *Value) interface{} {
			return v.String()
		},
		to: "xxx",
	},
	{
		from: map[string]interface{}{
			"foo": 1,
		},
		get: func(v *Value) interface{} {
			return v.InterfaceMap()
		},
		to: map[string]interface{}{
			"foo": uint64(1),
		},
	},
	{
		from: map[string]interface{}{
			"foo": "bar",
		},
		get: func(v *Value) interface{} {
			return v.StringMap()
		},
		to: map[string]string{
			"foo": "bar",
		},
	},
	{
		from: []int{1, 2, 3},
		get: func(v *Value) interface{} {
			return v.Ints()
		},
		to: []int{1, 2, 3},
	},
}

func _testMakeMsgpack(v interface{}) []byte {
	if r, err := msgpack.Marshal(v); err != nil {
		panic(fmt.Sprintf("Failed to MsgPack \"%s\": %s", spew.Sdump(v), err))
	} else {
		return r
	}
}

func TestValue_Msgpack(t *testing.T) {
	Convey("Marshal and unmarshal Msgpack", t, func() {
		for i, test := range testsTestValue_Msgpack {
			out := fmt.Sprintf("(%d) From \"%###v\"", i+1, test.from)
			Convey(out, func() {
				to := NewValue(nil)
				from := _testMakeMsgpack(test.from)
				err := msgpack.Unmarshal(from, to)
				So(err, ShouldBeNil)
				So(test.get(to), ShouldResemble, test.to)
				raw, err := msgpack.Marshal(to)
				So(string(raw), ShouldEqual, string(from))
			})
		}
	})
}
