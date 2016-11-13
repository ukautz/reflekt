package reflekt

import (
	"reflect"
	"github.com/davecgh/go-spew/spew"
)

type (
	Value struct {
		v interface{}
		r *reflect.Value
	}
)

func NewValue(v interface{}) *Value {
	return &Value{v: v}
}

func (this *Value) Dump() string {
	return spew.Sdump(this.v)
}

func (this *Value) Set(v interface{}) {
	this.v = v
	this.r = nil
}

func (this *Value) Reflect() *reflect.Value {
	if this.r == nil {
		r := reflect.ValueOf(this.v)
		this.r = &r
	}
	return this.r
}

func (this *Value) Interface() interface{} {
	return this.v
}

func (this *Value) Interfaces() []interface{} {
	return AsInterfaces(this.v)
}

func (this *Value) InterfaceMap() map[string]interface{} {
	return AsInterfaceMap(this.v)
}


func (this *Value) Int() int {
	return AsInt(this.v)
}

func (this *Value) Ints() []int {
	return AsInts(this.v)
}

func (this *Value) IntMap() map[string]int {
	return AsIntMap(this.v)
}

func (this *Value) Float() float64 {
	return AsFloat(this.v)
}

func (this *Value) Floats() []float64 {
	return AsFloats(this.v)
}

func (this *Value) FloatMap() map[string]float64 {
	return AsFloatMap(this.v)
}

func (this *Value) Bool() bool {
	return AsBool(this.v)
}

func (this *Value) Bools() []bool {
	return AsBools(this.v)
}

func (this *Value) BoolMap() map[string]bool {
	return AsBoolMap(this.v)
}

func (this *Value) String() string {
	return AsString(this.v)
}

func (this *Value) Strings() []string {
	return AsStrings(this.v)
}

func (this *Value) StringMap() map[string]string {
	return AsStringMap(this.v)
}
