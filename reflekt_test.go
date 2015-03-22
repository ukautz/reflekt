package reflekt

import (
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"reflect"
	"testing"
)

var testsInt = []struct {
	from      interface{}
	recognize bool
	to        int
}{
	{
		from:      int(123),
		recognize: true,
		to:        123,
	},
	{
		from:      int8(123),
		recognize: true,
		to:        123,
	},
	{
		from:      int16(123),
		recognize: true,
		to:        123,
	},
	{
		from:      int32(123),
		recognize: true,
		to:        123,
	},
	{
		from:      int64(123),
		recognize: true,
		to:        123,
	},
	{
		from:      uint(123),
		recognize: true,
		to:        123,
	},
	{
		from:      uint8(123),
		recognize: true,
		to:        123,
	},
	{
		from:      uint16(123),
		recognize: true,
		to:        123,
	},
	{
		from:      uint32(123),
		recognize: true,
		to:        123,
	},
	{
		from:      uint64(123),
		recognize: true,
		to:        123,
	},
	{
		from:      float64(123.123),
		recognize: false,
		to:        123,
	},
	{
		from:      float32(123.123),
		recognize: false,
		to:        123,
	},
	{
		from:      true,
		recognize: false,
		to:        1,
	},
	{
		from:      false,
		recognize: false,
		to:        0,
	},
	{
		from:      "123",
		recognize: false,
		to:        123,
	},
	{
		from:      "-123",
		recognize: false,
		to:        -123,
	},
	{
		from:      "123.345",
		recognize: false,
		to:        123,
	},
	{
		from:      "true",
		recognize: false,
		to:        1,
	},
	{
		from:      "TRUE",
		recognize: false,
		to:        1,
	},
	{
		from:      "false",
		recognize: false,
		to:        0,
	},
	{
		from:      "foo",
		recognize: false,
		to:        0,
	},
	{
		from:      map[string]interface{}{},
		recognize: false,
		to:        0,
	},
	{
		from:      []int{1, 2, 3},
		recognize: false,
		to:        0,
	},
}

func TestIsInt(t *testing.T) {
	Convey("Determine whether value is int", t, func() {
		for _, test := range testsInt {
			r := reflect.ValueOf(test.from)
			Convey(fmt.Sprintf("From %s (%v) expected %v", r.Type().String(), test.from, test.recognize), func() {
				So(test.recognize, ShouldEqual, IsInt(test.from))
			})
		}
	})
}

func TestAsInt(t *testing.T) {
	Convey("Try casting any value to int", t, func() {
		for _, test := range testsInt {
			r := reflect.ValueOf(test.from)
			Convey(fmt.Sprintf("From %s (%v) expected %d", r.Type().String(), test.from, test.to), func() {
				So(test.to, ShouldEqual, AsInt(test.from))
			})
		}
	})
}

var testsFloat = []struct {
	from      interface{}
	recognize bool
	to        float64
}{
	{
		from:      int(123),
		recognize: false,
		to:        float64(123.0),
	},
	{
		from:      int8(123),
		recognize: false,
		to:        float64(123.0),
	},
	{
		from:      int16(123),
		recognize: false,
		to:        float64(123.0),
	},
	{
		from:      int32(123),
		recognize: false,
		to:        float64(123.0),
	},
	{
		from:      int64(123),
		recognize: false,
		to:        float64(123.0),
	},
	{
		from:      uint(123),
		recognize: false,
		to:        float64(123.0),
	},
	{
		from:      uint8(123),
		recognize: false,
		to:        float64(123.0),
	},
	{
		from:      uint16(123),
		recognize: false,
		to:        float64(123.0),
	},
	{
		from:      uint32(123),
		recognize: false,
		to:        float64(123.0),
	},
	{
		from:      uint64(123),
		recognize: false,
		to:        float64(123.0),
	},
	{
		from:      float64(123.123),
		recognize: true,
		to:        float64(123.123),
	},
	{
		from:      float32(123.123),
		recognize: true,
		to:        float64(123.123),
	},
	{
		from:      "123",
		recognize: false,
		to:        float64(123.0),
	},
	{
		from:      "-123",
		recognize: false,
		to:        float64(-123.0),
	},
	{
		from:      "123.345",
		recognize: false,
		to:        float64(123.345),
	},
	{
		from:      "foo",
		recognize: false,
		to:        float64(0),
	},
	{
		from:      map[string]interface{}{},
		recognize: false,
		to:        float64(0),
	},
	{
		from:      []int{1, 2, 3},
		recognize: false,
		to:        float64(0),
	},
}

func TestIsFloat(t *testing.T) {
	Convey("Determine whether value is float", t, func() {
		for _, test := range testsFloat {
			r := reflect.ValueOf(test.from)
			Convey(fmt.Sprintf("From %s (%v) expected %v", r.Type().String(), test.from, test.recognize), func() {
				So(test.recognize, ShouldEqual, IsFloat(test.from))
			})
		}
	})
}

func TestAsFloat(t *testing.T) {
	Convey("Try casting any value to float", t, func() {
		for _, test := range testsFloat {
			r := reflect.ValueOf(test.from)
			Convey(fmt.Sprintf("From %s (%v) expected %v", r.Type().String(), test.from, test.to), func() {
				So(fmt.Sprintf(`%0.3f`, test.to), ShouldEqual, fmt.Sprintf(`%0.3f`, AsFloat(test.from)))
			})
		}
	})
}

var testsBool = []struct {
	from interface{}
	to   bool
}{
	{
		from: 1,
		to:   true,
	},
	{
		from: 1.1,
		to:   true,
	},
	{
		from: uint(1),
		to:   true,
	},
	{
		from: "true",
		to:   true,
	},
	{
		from: "TRUE",
		to:   true,
	},
	{
		from: true,
		to:   true,
	},
	{
		from: "1",
		to:   true,
	},
	{
		from: 0,
		to:   false,
	},
	{
		from: float64(0.0),
		to:   false,
	},
	{
		from: uint(0),
		to:   false,
	},
	{
		from: "false",
		to:   false,
	},
	{
		from: "FALSE",
		to:   false,
	},
	{
		from: false,
		to:   false,
	},
	{
		from: "0",
		to:   false,
	},
}

func TestAsBool(t *testing.T) {
	Convey("Try casting any value to bool", t, func() {
		for _, test := range testsBool {
			r := reflect.ValueOf(test.from)
			Convey(fmt.Sprintf("From %s (%v) expected %v", r.Type().String(), test.from, test.to), func() {
				So(test.to, ShouldEqual, AsBool(test.from))
			})
		}
	})
}

var testsString = []struct {
	from interface{}
	to   string
}{
	{
		from: 1,
		to:   "1",
	},
	{
		from: 1.1,
		to:   "1.100000",
	},
	{
		from: uint(1),
		to:   "1",
	},
	{
		from: true,
		to:   "true",
	},
	{
		from: false,
		to:   "false",
	},
	{
		from: "foo",
		to:   "foo",
	},
	{
		from: map[string]interface{}{"foo": "bar"},
		to:   "",
	},
}

func TestAsString(t *testing.T) {
	Convey("Try casting any value to string", t, func() {
		for _, test := range testsString {
			r := reflect.ValueOf(test.from)
			Convey(fmt.Sprintf("From %s (%v) expected %v", r.Type().String(), test.from, test.to), func() {
				So(AsString(test.from), ShouldEqual, test.to)
			})
		}
	})
}
