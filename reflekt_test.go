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
	toSlice   []int
}{
	{
		from:      nil,
		recognize: false,
		to:        0,
		toSlice:   []int{},
	},
	{
		from:      int(123),
		recognize: true,
		to:        123,
		toSlice:   []int{123},
	},
	{
		from:      int8(123),
		recognize: true,
		to:        123,
		toSlice:   []int{123},
	},
	{
		from:      int16(123),
		recognize: true,
		to:        123,
		toSlice:   []int{123},
	},
	{
		from:      int32(123),
		recognize: true,
		to:        123,
		toSlice:   []int{123},
	},
	{
		from:      int64(123),
		recognize: true,
		to:        123,
		toSlice:   []int{123},
	},
	{
		from:      uint(123),
		recognize: true,
		to:        123,
		toSlice:   []int{123},
	},
	{
		from:      uint8(123),
		recognize: true,
		to:        123,
		toSlice:   []int{123},
	},
	{
		from:      uint16(123),
		recognize: true,
		to:        123,
		toSlice:   []int{123},
	},
	{
		from:      uint32(123),
		recognize: true,
		to:        123,
		toSlice:   []int{123},
	},
	{
		from:      uint64(123),
		recognize: true,
		to:        123,
		toSlice:   []int{123},
	},
	{
		from:      float64(123.123),
		recognize: false,
		to:        123,
		toSlice:   []int{123},
	},
	{
		from:      float32(123.123),
		recognize: false,
		to:        123,
		toSlice:   []int{123},
	},
	{
		from:      true,
		recognize: false,
		to:        1,
		toSlice:   []int{1},
	},
	{
		from:      false,
		recognize: false,
		to:        0,
		toSlice:   []int{0},
	},
	{
		from:      "123",
		recognize: false,
		to:        123,
		toSlice:   []int{123},
	},
	{
		from:      "-123",
		recognize: false,
		to:        -123,
		toSlice:   []int{-123},
	},
	{
		from:      "123.345",
		recognize: false,
		to:        123,
		toSlice:   []int{123},
	},
	{
		from:      "true",
		recognize: false,
		to:        1,
		toSlice:   []int{1},
	},
	{
		from:      "TRUE",
		recognize: false,
		to:        1,
		toSlice:   []int{1},
	},
	{
		from:      "false",
		recognize: false,
		to:        0,
		toSlice:   []int{0},
	},
	{
		from:      reflect.ValueOf(6),
		recognize: true,
		to:        6,
		toSlice:   []int{6},
	},
	{
		from:      reflect.ValueOf(6.3),
		recognize: false,
		to:        6,
		toSlice:   []int{6},
	},
	{
		from:      reflect.ValueOf("6.3"),
		recognize: false,
		to:        6,
		toSlice:   []int{6},
	},
	{
		from:      "foo",
		recognize: false,
		to:        0,
		toSlice:   []int{0},
	},
	{
		from:      map[string]interface{}{},
		recognize: false,
		to:        0,
		toSlice:   []int{0},
	},
	{
		from:      []bool{true, false, true},
		recognize: false,
		to:        int(0),
		toSlice:   []int{1, 0, 1},
	},
	{
		from:      []int{1, 2, 3},
		recognize: false,
		to:        0,
		toSlice:   []int{1, 2, 3},
	},
	{
		from:      []float64{1.1, 2.2, 3.3},
		recognize: false,
		to:        0,
		toSlice:   []int{1, 2, 3},
	},
	{
		from:      []string{"1", "2", "3"},
		recognize: false,
		to:        0,
		toSlice:   []int{1, 2, 3},
	},
}

func TestIsInt(t *testing.T) {
	Convey("Determine whether value is int", t, func() {
		for i, test := range testsInt {
			Convey(fmt.Sprintf("%d) From %s (%v) expected %v (is int)", i, typeName(test.from), test.from, test.recognize), func() {
				So(test.recognize, ShouldEqual, IsInt(test.from))
			})
		}
	})
}

func TestAsInt(t *testing.T) {
	Convey("Try casting any value to int", t, func() {
		for i, test := range testsInt {
			Convey(fmt.Sprintf("%d) From %s (%v) expected %d (as int)", i, typeName(test.from), test.from, test.to), func() {
				So(test.to, ShouldEqual, AsInt(test.from))
			})
		}
	})
}

func TestAsInts(t *testing.T) {
	Convey("Try casting any value to slice of ints", t, func() {
		for i, test := range testsInt {
			Convey(fmt.Sprintf("%d) From %s (%v) expected %v (as ints)", i, typeName(test.from), test.from, test.toSlice), func() {
				So(test.toSlice, ShouldResemble, AsInts(test.from))
			})
		}
	})
}

var testsFloat = []struct {
	from      interface{}
	recognize bool
	to        float64
	toSlice   []float64
}{
	{
		from:      nil,
		recognize: false,
		to:        float64(0),
		toSlice:   []float64{},
	},
	{
		from:      int(123),
		recognize: false,
		to:        float64(123.0),
		toSlice:   []float64{123.0},
	},
	{
		from:      int8(123),
		recognize: false,
		to:        float64(123.0),
		toSlice:   []float64{123.0},
	},
	{
		from:      int16(123),
		recognize: false,
		to:        float64(123.0),
		toSlice:   []float64{123.0},
	},
	{
		from:      int32(123),
		recognize: false,
		to:        float64(123.0),
		toSlice:   []float64{123.0},
	},
	{
		from:      int64(123),
		recognize: false,
		to:        float64(123.0),
		toSlice:   []float64{123.0},
	},
	{
		from:      uint(123),
		recognize: false,
		to:        float64(123.0),
		toSlice:   []float64{123.0},
	},
	{
		from:      uint8(123),
		recognize: false,
		to:        float64(123.0),
		toSlice:   []float64{123.0},
	},
	{
		from:      uint16(123),
		recognize: false,
		to:        float64(123.0),
		toSlice:   []float64{123.0},
	},
	{
		from:      uint32(123),
		recognize: false,
		to:        float64(123.0),
		toSlice:   []float64{123.0},
	},
	{
		from:      uint64(123),
		recognize: false,
		to:        float64(123.0),
		toSlice:   []float64{123.0},
	},
	{
		from:      float64(123.123),
		recognize: true,
		to:        float64(123.123),
		toSlice:   []float64{123.123},
	},
	{
		from:      float32(123.123),
		recognize: true,
		to:        float64(123.123),
		toSlice:   []float64{123.123},
	},
	{
		from:      "123",
		recognize: false,
		to:        float64(123.0),
		toSlice:   []float64{123.0},
	},
	{
		from:      "-123",
		recognize: false,
		to:        float64(-123.0),
		toSlice:   []float64{-123.0},
	},
	{
		from:      "123.345",
		recognize: false,
		to:        float64(123.345),
		toSlice:   []float64{123.345},
	},
	{
		from:      reflect.ValueOf(6),
		recognize: false,
		to:        float64(6.0),
		toSlice:   []float64{6.0},
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
		toSlice:   []float64{6.3},
	},
	{
		from:      "foo",
		recognize: false,
		to:        float64(0),
		toSlice:   []float64{0},
	},
	{
		from:      map[string]interface{}{},
		recognize: false,
		to:        float64(0),
		toSlice:   []float64{0},
	},
	{
		from:      []bool{true, false, true},
		recognize: false,
		to:        float64(0),
		toSlice:   []float64{1, 0, 1},
	},
	{
		from:      []int{1, 2, 3},
		recognize: false,
		to:        float64(0),
		toSlice:   []float64{1, 2, 3},
	},
	{
		from:      []float64{1.1, 2.2, 3.3},
		recognize: false,
		to:        0,
		toSlice:   []float64{1.1, 2.2, 3.3},
	},
	{
		from:      []string{"1", "2", "3"},
		recognize: false,
		to:        0,
		toSlice:   []float64{1, 2, 3},
	},
}

func TestIsFloat(t *testing.T) {
	Convey("Determine whether value is float", t, func() {
		for i, test := range testsFloat {
			Convey(fmt.Sprintf("%d) From %s (%v) expected %v", i, typeName(test.from), test.from, test.recognize), func() {
				So(test.recognize, ShouldEqual, IsFloat(test.from))
			})
		}
	})
}

func TestAsFloat(t *testing.T) {
	Convey("Try casting any value to float", t, func() {
		for i, test := range testsFloat {
			Convey(fmt.Sprintf("%d) From %s (%v) expected %v", i, typeName(test.from), test.from, test.to), func() {
				So(fmt.Sprintf(`%0.3f`, test.to), ShouldEqual, fmt.Sprintf(`%0.3f`, AsFloat(test.from)))
			})
		}
	})
}

func TestAsFloats(t *testing.T) {
	Convey("Try casting any value to slice of floats", t, func() {
		for i, test := range testsFloat {
			Convey(fmt.Sprintf("%d) From %s (%v) expected %v", i, typeName(test.from), test.from, test.to), func() {
				So(fmt.Sprintf(`%0.3f`, test.to), ShouldEqual, fmt.Sprintf(`%0.3f`, AsFloat(test.from)))
			})
		}
	})
}

var testsBool = []struct {
	from    interface{}
	to      bool
	toSlice []bool
}{
	{
		from:    nil,
		to:      false,
		toSlice: []bool{},
	},
	{
		from:    1,
		to:      true,
		toSlice: []bool{true},
	},
	{
		from:    1.1,
		to:      true,
		toSlice: []bool{true},
	},
	{
		from:    uint(1),
		to:      true,
		toSlice: []bool{true},
	},
	{
		from:    "true",
		to:      true,
		toSlice: []bool{true},
	},
	{
		from:    "TRUE",
		to:      true,
		toSlice: []bool{true},
	},
	{
		from:    true,
		to:      true,
		toSlice: []bool{true},
	},
	{
		from:    "1",
		to:      true,
		toSlice: []bool{true},
	},
	{
		from:    reflect.ValueOf(true),
		to:      true,
		toSlice: []bool{true},
	},
	{
		from:    reflect.ValueOf("1"),
		to:      true,
		toSlice: []bool{true},
	},
	{
		from:    reflect.ValueOf(1),
		to:      true,
		toSlice: []bool{true},
	},
	{
		from:    0,
		to:      false,
		toSlice: []bool{false},
	},
	{
		from:    float64(0.0),
		to:      false,
		toSlice: []bool{false},
	},
	{
		from:    uint(0),
		to:      false,
		toSlice: []bool{false},
	},
	{
		from:    "false",
		to:      false,
		toSlice: []bool{false},
	},
	{
		from:    "FALSE",
		to:      false,
		toSlice: []bool{false},
	},
	{
		from:    false,
		to:      false,
		toSlice: []bool{false},
	},
	{
		from:    "0",
		to:      false,
		toSlice: []bool{false},
	},
	{
		from:    reflect.ValueOf(false),
		to:      false,
		toSlice: []bool{false},
	},
	{
		from:    []bool{true, false, true},
		to:      false,
		toSlice: []bool{true, false, true},
	},
	{
		from:    []int{1, 0, 1},
		to:      false,
		toSlice: []bool{true, false, true},
	},
	{
		from:    []float64{1.1, 0.0, 2.2},
		to:      false,
		toSlice: []bool{true, false, true},
	},
	{
		from:    []string{"1", "0", "TRUE"},
		to:      false,
		toSlice: []bool{true, false, true},
	},
}

func TestAsBool(t *testing.T) {
	Convey("Try casting any value to bool", t, func() {
		for i, test := range testsBool {
			Convey(fmt.Sprintf("%d) From %s (%v) expected %v", i, typeName(test.from), test.from, test.to), func() {
				So(test.to, ShouldEqual, AsBool(test.from))
			})
		}
	})
}

func TestAsBools(t *testing.T) {
	Convey("Try casting any value to slice of bools", t, func() {
		for i, test := range testsBool {
			Convey(fmt.Sprintf("%d) From %s (%v) expected %v", i, typeName(test.from), test.from, test.to), func() {
				So(test.to, ShouldEqual, AsBool(test.from))
			})
		}
	})
}

var testsString = []struct {
	from    interface{}
	to      string
	toSlice []string
}{
	{
		from:    nil,
		to:      "",
		toSlice: []string{},
	},
	{
		from:    1,
		to:      "1",
		toSlice: []string{"1"},
	},
	{
		from:    1.1,
		to:      "1.1",
		toSlice: []string{"1.1"},
	},
	{
		from:    uint(1),
		to:      "1",
		toSlice: []string{"1"},
	},
	{
		from:    true,
		to:      "true",
		toSlice: []string{"true"},
	},
	{
		from:    false,
		to:      "false",
		toSlice: []string{"false"},
	},
	{
		from:    "foo",
		to:      "foo",
		toSlice: []string{"foo"},
	},
	{
		from:    reflect.ValueOf("foo"),
		to:      "foo",
		toSlice: []string{"foo"},
	},
	{
		from:    reflect.ValueOf(123),
		to:      "123",
		toSlice: []string{"123"},
	},
	{
		from:    reflect.ValueOf(true),
		to:      "true",
		toSlice: []string{"true"},
	},
	{
		from:    map[string]interface{}{"foo": "bar"},
		to:      "",
		toSlice: []string{""},
	},
	{
		from:    []bool{true, false},
		to:      "",
		toSlice: []string{"true", "false"},
	},
	{
		from:    []int{1, 2, 3},
		to:      "",
		toSlice: []string{"1", "2", "3"},
	},
	{
		from:    []float64{1.1, 2.2, 3.3},
		to:      "",
		toSlice: []string{"1.1", "2.2", "3.3"},
	},
	{
		from:    []string{"foo", "bar"},
		to:      "",
		toSlice: []string{"foo", "bar"},
	},
}

func TestAsString(t *testing.T) {
	Convey("Try casting any value to string", t, func() {
		for i, test := range testsString {
			Convey(fmt.Sprintf("%d) From %s (%v) expected %v", i, typeName(test.from), test.from, test.to), func() {
				So(AsString(test.from), ShouldEqual, test.to)
			})
		}
	})
}

func TestAsStrings(t *testing.T) {
	Convey("Try casting any value to slice of strings", t, func() {
		for i, test := range testsString {
			Convey(fmt.Sprintf("%d) From %s (%v) expected %v", i, typeName(test.from), test.from, test.to), func() {
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
		from: nil,
		to:   nil,
	},
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
			Convey(fmt.Sprintf("%d) From %s", idx, typeName(test.from)), func() {
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
		from: nil,
		to:   nil,
	},
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
			Convey(fmt.Sprintf("%d) From %s", idx, typeName(test.from)), func() {
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
		from: nil,
		to:   nil,
	},
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
			Convey(fmt.Sprintf("%d) From %s", idx, typeName(test.from)), func() {
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
		from: nil,
		to:   nil,
	},
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
			Convey(fmt.Sprintf("%d) From %s", idx, typeName(test.from)), func() {
				So(AsStringMap(test.from), ShouldResemble, test.to)
			})
		}
	})
}

var testsInterfaceMap = []struct {
	from interface{}
	to   map[string]interface{}
}{
	{
		from: nil,
		to:   nil,
	},
	{
		from: 1,
		to:   map[string]interface{}{},
	},
	{
		from: map[interface{}]interface{}{"foo": 1, "bar": "2.1", "baz": 3.4, "zoing": true},
		to:   map[string]interface{}{"foo": 1, "bar": "2.1", "baz": 3.4, "zoing": true},
	},
}

func TestAsInterfaceMap(t *testing.T) {
	Convey("Try casting any map to map[string]interface{}", t, func() {
		for idx, test := range testsInterfaceMap {
			Convey(fmt.Sprintf("%d) From %s", idx, typeName(test.from)), func() {
				So(AsInterfaceMap(test.from), ShouldResemble, test.to)
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

func typeName(v interface{}) string {
	if v == nil {
		return "nil"
	} else {
		return reflect.TypeOf(v).String()
	}
}
