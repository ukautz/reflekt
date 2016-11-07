package reflekt

import (
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"reflect"
	"testing"
)

type t1 struct {
	A int
	B string
	C t2
}

type t2 struct {
	D int
	E *t3
}

type t3 struct {
	Foo string
}

type i1 interface {
	Foo() string
}

type t4 struct {
	T string
}

func (this *t4) Foo() string {
	return this.T
}

type t5 struct {
	I i1
}

type t6 struct {
	I []i1
}

type t7 struct {
	t3
	J string
}

var testsStructAsMap = []struct {
	from interface{}
	to   map[string]interface{}
	lc   bool
}{
	{
		from: 1,
		to:   map[string]interface{}{},
	},
	{
		from: struct{ X int }{X: 1},
		to:   map[string]interface{}{"X": 1},
	},
	{
		from: struct {
			X int
			Y string
		}{X: 1, Y: "bla"},
		to: map[string]interface{}{"x": 1, "y": "bla"},
		lc: true,
	},
	{
		from: t1{
			A: 1,
			B: "two",
			C: t2{
				D: 3,
				E: &t3{
					Foo: "Bar",
				},
			},
		},
		to: map[string]interface{}{
			"a": 1,
			"b": "two",
			"c": map[string]interface{}{
				"d": 3,
				"e": map[string]interface{}{
					"foo": "Bar",
				},
			},
		},
		lc: true,
	},
	{
		from: t5{
			I: &t4{"bar"},
		},
		to: map[string]interface{}{
			"I": map[string]interface{}{
				"T": "bar",
			},
		},
	},
	{
		from: t6{
			I: []i1{
				&t4{"bar"},
				&t4{"baz"},
			},
		},
		to: map[string]interface{}{
			"I": []interface{}{
				map[string]interface{}{"T": "bar"},
				map[string]interface{}{"T": "baz"},
			},
		},
	},
	{
		from: t7{
			t3: t3{
				Foo: "foo",
			},
			J: "j",
		},
		to: map[string]interface{}{
			"Foo": "foo",
			"J":   "j",
		},
	},
}

func TestStructAsMap(t *testing.T) {
	t.Skip("Still working on that")
	Convey("Try casting any map to map[string]interface{}", t, func() {
		for idx, test := range testsStructAsMap {
			r := reflect.ValueOf(test.from)
			Convey(fmt.Sprintf("%d) From %s", idx, r.Type().String()), func() {
				So(StructAsMap(test.from, test.lc), ShouldResemble, test.to)
			})
		}
	})
}

func TestFillStruct(t *testing.T) {
	t.Skip("Still working on that")
	Convey("Try filling struct from map[string]interface{}", t, func() {
		for idx, test := range testsStructAsMap {
			if idx != 5 {
				continue
			}
			r := reflect.ValueOf(test.from)
			Convey(fmt.Sprintf("%d) From %s", idx, r.Type().String()), func() {
				s := reflect.New(reflect.TypeOf(test.from))
				f := NewStructFiller()
				f.Register(reflect.TypeOf((*i1)(nil)).Elem(), func(v interface{}) reflect.Type {
					return reflect.TypeOf(t4{})
				})
				err := f.Fill(s.Interface(), test.to)
				fmt.Printf("SHOULD: %s\n", reflect.ValueOf(test.from).Kind())
				fmt.Printf("IS    : %s\n", reflect.ValueOf(s).Kind())
				So(err, ShouldBeNil)
				So(s.Elem().Interface(), ShouldResemble, test.from)
			})
		}
	})
}
