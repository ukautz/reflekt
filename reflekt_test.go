package reflekt

import (
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"reflect"
	"sort"
	"strings"
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
		from:      reflect.ValueOf(6),
		recognize: true,
		to:        6,
	},
	{
		from:      reflect.ValueOf(6.3),
		recognize: false,
		to:        6,
	},
	{
		from:      reflect.ValueOf("6.3"),
		recognize: false,
		to:        6,
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
		from:      reflect.ValueOf(6),
		recognize: false,
		to:        float64(6.0),
	},
	{
		from:      reflect.ValueOf(6.3),
		recognize: true,
		to:        float64(6.3),
	},
	{
		from:      reflect.ValueOf("6.3"),
		recognize: false,
		to:        float64(6.3),
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
		from: reflect.ValueOf(true),
		to:   true,
	},
	{
		from: reflect.ValueOf("1"),
		to:   true,
	},
	{
		from: reflect.ValueOf(1),
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
	{
		from: reflect.ValueOf(false),
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
		to:   "1.1",
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
		from: reflect.ValueOf("foo"),
		to:   "foo",
	},
	{
		from: reflect.ValueOf(123),
		to:   "123",
	},
	{
		from: reflect.ValueOf(true),
		to:   "true",
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

var testsIntMap = []struct {
	from interface{}
	to   map[string]int
}{
	{
		from: 1,
		to:   map[string]int{},
	},
	{
		from: map[interface{}]interface{}{"foo": 1, "bar": "2", "baz": 3.4, "zoing": true},
		to:   map[string]int{"foo": 1, "bar": 2, "baz": 3, "zoing": 1},
	},
	{
		from: map[string]float64{"foo": 1, "bar": 2},
		to:   map[string]int{"foo": 1, "bar": 2},
	},
	{
		from: map[string]string{"foo": "1", "bar": "2", "baz": "3.4", "zoing": "true"},
		to:   map[string]int{"foo": 1, "bar": 2, "baz": 3, "zoing": 1},
	},
	{
		from: map[string]int{"foo": 1, "bar": 2, "baz": 3, "zoing": 1},
		to:   map[string]int{"foo": 1, "bar": 2, "baz": 3, "zoing": 1},
	},
}

func TestAsIntMap(t *testing.T) {
	Convey("Try casting any map to map[string]int", t, func() {
		for idx, test := range testsIntMap {
			r := reflect.ValueOf(test.from)
			Convey(fmt.Sprintf("%d) From %s", idx, r.Type().String()), func() {
				So(AsIntMap(test.from), ShouldResemble, test.to)
			})
		}
	})
}

var testsFloatMap = []struct {
	from interface{}
	to   map[string]float64
}{
	{
		from: 1,
		to:   map[string]float64{},
	},
	{
		from: map[interface{}]interface{}{"foo": 1, "bar": "2.1", "baz": 3.4, "zoing": true},
		to:   map[string]float64{"foo": 1, "bar": 2.1, "baz": 3.4, "zoing": 1},
	},
	{
		from: map[string]int{"foo": 1, "bar": 2},
		to:   map[string]float64{"foo": 1, "bar": 2},
	},
	{
		from: map[string]string{"foo": "1", "bar": "3.4", "baz": "true"},
		to:   map[string]float64{"foo": 1, "bar": 3.4, "baz": 1},
	},
	{
		from: map[string]float64{"foo": 1, "bar": 2, "baz": 1},
		to:   map[string]float64{"foo": 1, "bar": 2, "baz": 1},
	},
}

func TestAsFloatMap(t *testing.T) {
	Convey("Try casting any map to map[string]float64", t, func() {
		for idx, test := range testsFloatMap {
			r := reflect.ValueOf(test.from)
			Convey(fmt.Sprintf("%d) From %s", idx, r.Type().String()), func() {
				So(AsFloatMap(test.from), ShouldResemble, test.to)
			})
		}
	})
}

var testsBoolMap = []struct {
	from interface{}
	to   map[string]bool
}{
	{
		from: 1,
		to:   map[string]bool{},
	},
	{
		from: map[interface{}]interface{}{"foo": 1, "bar": "2.1", "baz": 3.4, "zoing": true},
		to:   map[string]bool{"foo": true, "bar": true, "baz": true, "zoing": true},
	},
	{
		from: map[string]int{"foo": 1, "bar": 0},
		to:   map[string]bool{"foo": true, "bar": false},
	},
	{
		from: map[string]float64{"foo": 1.3, "bar": 0},
		to:   map[string]bool{"foo": true, "bar": false},
	},
	{
		from: map[string]string{"foo": "1", "bar": "3.4", "baz": "true", "zoing": "false"},
		to:   map[string]bool{"foo": true, "bar": true, "baz": true, "zoing": false},
	},
	{
		from: map[string]bool{"foo": true, "bar": false},
		to:   map[string]bool{"foo": true, "bar": false},
	},
}

func TestAsBoolMap(t *testing.T) {
	Convey("Try casting any map to map[string]bool", t, func() {
		for idx, test := range testsBoolMap {
			r := reflect.ValueOf(test.from)
			Convey(fmt.Sprintf("%d) From %s", idx, r.Type().String()), func() {
				So(AsBoolMap(test.from), ShouldResemble, test.to)
			})
		}
	})
}

var testsStringMap = []struct {
	from interface{}
	to   map[string]string
}{
	{
		from: 1,
		to:   map[string]string{},
	},
	{
		from: map[interface{}]interface{}{"foo": 1, "bar": "2.1", "baz": 3.4, "zoing": true},
		to:   map[string]string{"foo": "1", "bar": "2.1", "baz": "3.4", "zoing": "true"},
	},
	{
		from: map[string]int{"foo": 1, "bar": 0},
		to:   map[string]string{"foo": "1", "bar": "0"},
	},
	{
		from: map[string]float64{"foo": 1.3, "bar": 0},
		to:   map[string]string{"foo": "1.3", "bar": "0"},
	},
	{
		from: map[string]bool{"foo": true, "bar": false},
		to:   map[string]string{"foo": "true", "bar": "false"},
	},
	{
		from: map[string]string{"foo": "1", "bar": "3.4", "baz": "true", "zoing": "false"},
		to:   map[string]string{"foo": "1", "bar": "3.4", "baz": "true", "zoing": "false"},
	},
}

func TestAsStringMap(t *testing.T) {
	Convey("Try casting any map to map[string]string", t, func() {
		for idx, test := range testsStringMap {
			r := reflect.ValueOf(test.from)
			Convey(fmt.Sprintf("%d) From %s", idx, r.Type().String()), func() {
				So(AsStringMap(test.from), ShouldResemble, test.to)
			})
		}
	})
}

func serializeMap(v interface{}) string {
	r := reflect.ValueOf(v)
	if r.Kind() != reflect.Map {
		return fmt.Sprintf("%v", r.Interface())
	}
	keys := []string{}
	orig := make(map[string]reflect.Value)
	for _, k := range r.MapKeys() {
		kv := fmt.Sprintf("%s", k.Interface())
		keys = append(keys, kv)
		orig[kv] = k
	}
	sort.Strings(keys)

	out := []string{}
	for _, k := range keys {
		out = append(out, fmt.Sprintf("%s: %v", k, r.MapIndex(orig[k]).Interface()))
	}

	return fmt.Sprintf("%s{%s}", r.Type(), strings.Join(out, ", "))
}
