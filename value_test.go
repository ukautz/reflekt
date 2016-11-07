package reflekt

import (
	"testing"
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/davecgh/go-spew/spew"
	"strings"
)

var testsValue = []struct {
	from           interface{}
	toInterface    interface{}
	toInterfaces   []interface{}
	toInterfaceMap map[string]interface{}
	toInt          int
	toInts         []int
	toIntMap       map[string]int
	toFloat        float64
	toFloats       []float64
	toFloatMap     map[string]float64
	toBool         bool
	toBools        []bool
	toBoolMap      map[string]bool
	toString       string
	toStrings      []string
	toStringMap    map[string]string
}{
	/*
	------------------------------------
	  SCALAR
	------------------------------------
	*/
	{
		from:           nil,
		toInterface:    nil,
		toInterfaces:   []interface{}{},
		toInterfaceMap: nil,
		toInt:          0,
		toInts:         []int{},
		toIntMap:       nil,
		toFloat:        0,
		toFloats:       []float64{},
		toFloatMap:     nil,
		toBool:         false,
		toBools:        []bool{},
		toBoolMap:      nil,
		toString:       "",
		toStrings:      []string{},
		toStringMap:    nil,
	},
	{
		from:           1,
		toInterface:    1,
		toInterfaces:   []interface{}{1},
		toInterfaceMap: map[string]interface{}{},
		toInt:          1,
		toInts:         []int{1},
		toIntMap:       map[string]int{},
		toFloat:        1,
		toFloats:       []float64{1},
		toFloatMap:     map[string]float64{},
		toBool:         true,
		toBools:        []bool{true},
		toBoolMap:      map[string]bool{},
		toString:       "1",
		toStrings:      []string{"1"},
		toStringMap:    map[string]string{},
	},
	{
		from:           1.234,
		toInterface:    1.234,
		toInterfaces:   []interface{}{1.234},
		toInterfaceMap: map[string]interface{}{},
		toInt:          1,
		toInts:         []int{1},
		toIntMap:       map[string]int{},
		toFloat:        1.234,
		toFloats:       []float64{1.234},
		toFloatMap:     map[string]float64{},
		toBool:         true,
		toBools:        []bool{true},
		toBoolMap:      map[string]bool{},
		toString:       "1.234",
		toStrings:      []string{"1.234"},
		toStringMap:    map[string]string{},
	},
	{
		from:           "true",
		toInterface:    "true",
		toInterfaces:   []interface{}{"true"},
		toInterfaceMap: map[string]interface{}{},
		toInt:          1,
		toInts:         []int{1},
		toIntMap:       map[string]int{},
		toFloat:        1,
		toFloats:       []float64{1},
		toFloatMap:     map[string]float64{},
		toBool:         true,
		toBools:        []bool{true},
		toBoolMap:      map[string]bool{},
		toString:       "true",
		toStrings:      []string{"true"},
		toStringMap:    map[string]string{},
	},
	{
		from:           "123.45",
		toInterface:    "123.45",
		toInterfaces:   []interface{}{"123.45"},
		toInterfaceMap: map[string]interface{}{},
		toInt:          123,
		toInts:         []int{123},
		toIntMap:       map[string]int{},
		toFloat:        123.45,
		toFloats:       []float64{123.45},
		toFloatMap:     map[string]float64{},
		toBool:         true,
		toBools:        []bool{true},
		toBoolMap:      map[string]bool{},
		toString:       "123.45",
		toStrings:      []string{"123.45"},
		toStringMap:    map[string]string{},
	},
	/*
	------------------------------------
	  ARRAY
	------------------------------------
	*/
	{
		from:           []int{1, 2},
		toInterface:    []int{1, 2},
		toInterfaces:   []interface{}{int(1), int(2)},
		toInterfaceMap: map[string]interface{}{},
		toInt:          0,
		toInts:         []int{1, 2},
		toIntMap:       map[string]int{},
		toFloat:        0,
		toFloats:       []float64{1, 2},
		toFloatMap:     map[string]float64{},
		toBool:         false,
		toBools:        []bool{true, true},
		toBoolMap:      map[string]bool{},
		toString:       "",
		toStrings:      []string{"1", "2"},
		toStringMap:    map[string]string{},
	},
	{
		from:           []float64{1.1, 2.2},
		toInterface:    []float64{1.1, 2.2},
		toInterfaces:   []interface{}{float64(1.1), float64(2.2)},
		toInterfaceMap: map[string]interface{}{},
		toInt:          0,
		toInts:         []int{1, 2},
		toIntMap:       map[string]int{},
		toFloat:        0,
		toFloats:       []float64{1.1, 2.2},
		toFloatMap:     map[string]float64{},
		toBool:         false,
		toBools:        []bool{true, true},
		toBoolMap:      map[string]bool{},
		toString:       "",
		toStrings:      []string{"1.1", "2.2"},
		toStringMap:    map[string]string{},
	},
	{
		from:           []string{"1.1", "2.2"},
		toInterface:    []string{"1.1", "2.2"},
		toInterfaces:   []interface{}{"1.1", "2.2"},
		toInterfaceMap: map[string]interface{}{},
		toInt:          0,
		toInts:         []int{1, 2},
		toIntMap:       map[string]int{},
		toFloat:        0,
		toFloats:       []float64{1.1, 2.2},
		toFloatMap:     map[string]float64{},
		toBool:         false,
		toBools:        []bool{true, true},
		toBoolMap:      map[string]bool{},
		toString:       "",
		toStrings:      []string{"1.1", "2.2"},
		toStringMap:    map[string]string{},
	},
	/*
	------------------------------------
	  MAP
	------------------------------------
	*/
	{
		from:           map[string]interface{}{"foo": 1},
		toInterface:    map[string]interface{}{"foo": 1},
		toInterfaces:   []interface{}{map[string]interface{}{"foo": 1}},
		toInterfaceMap: map[string]interface{}{"foo": 1},
		toInt:          0,
		toInts:         []int{0},
		toIntMap:       map[string]int{"foo": 1},
		toFloat:        0,
		toFloats:       []float64{0},
		toFloatMap:     map[string]float64{"foo": 1},
		toBool:         false,
		toBools:        []bool{false},
		toBoolMap:      map[string]bool{"foo": true},
		toString:       "",
		toStrings:      []string{""},
		toStringMap:    map[string]string{"foo": "1"},
	},
}

func TestValue(t *testing.T) {
	Convey("Accessing value casters", t, func() {
		for i, test := range testsValue {
			o := spew.Sdump(test.from)
			o = strings.Replace(o, "\n", " ", -1)
			Convey(fmt.Sprintf("%d: \"%s\"", i, o), func() {
				v := NewValue(test.from)
				//fmt.Printf("\nFROM: %s\nTO  : %s\n\n", spew.Sdump(test.toInterface), spew.Sdump(v.Interface()))
				So(test.toInterface, ShouldResemble, v.Interface())
				So(test.toInterfaces, ShouldResemble, v.Interfaces())
				So(test.toInterfaceMap, ShouldResemble, v.InterfaceMap())
				So(test.toInt, ShouldEqual, v.Int())
				So(test.toInts, ShouldResemble, v.Ints())
				So(test.toIntMap, ShouldResemble, v.IntMap())
				So(test.toFloat, ShouldEqual, v.Float())
				So(test.toFloats, ShouldResemble, v.Floats())
				So(test.toFloatMap, ShouldResemble, v.FloatMap())
				So(test.toBool, ShouldEqual, v.Bool())
				So(test.toBools, ShouldResemble, v.Bools())
				So(test.toBoolMap, ShouldResemble, v.BoolMap())
				So(test.toString, ShouldEqual, v.String())
				So(test.toStrings, ShouldResemble, v.Strings())
				So(test.toStringMap, ShouldResemble, v.StringMap())
			})
		}
	})
}