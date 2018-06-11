[![Build Status](https://travis-ci.org/ukautz/reflekt.svg?branch=v4)](https://travis-ci.org/ukautz/reflekt)
[![Coverage](https://gocover.io/_badge/github.com/ukautz/reflekt)](http://gocover.io/github.com/ukautz/reflekt)
[![GoDoc](https://godoc.org/github.com/ukautz/reflekt?status.svg)](https://godoc.org/github.com/ukautz/reflekt)



reflekt
=======

This package provides lazy developers shorthands for determining or casting
(to) primitive Go types.


Installation
------------

```bash
$ go get gopkg.in/ukautz/reflekt.v4
```


Documentation
-------------

GoDoc can be [found here](http://godoc.org/github.com/ukautz/reflekt)


Examples
--------

Determine whether anything is of a primitive type. 

```go
import "gopkg.in/ukautz/reflekt.v4"

if reflekt.IsInt(v) {
    fmt.Printf("Yes, %d is a native integer type", v)
} else if reflekt.IsFloat(v) {
    fmt.Printf("Yes, %f is a native float type", v)
}
```

(Try) Cast given value into respective type (tries to exhaust all possibilities):

### Casting scalars

``` go
import "gopkg.in/ukautz/reflekt.v4"

si := "1.2"
ii := 1
fi := 1.2
bi := true

// always as int
var i int
i = reflekt.AsInt(si) // 1
i = reflekt.AsInt(ii) // 1
i = reflekt.AsInt(fi) // 1
i = reflekt.AsInt(bi) // 1

// always as float
var f float64
f = reflekt.AsFloat(si) // 1.2
f = reflekt.AsFloat(ii) // 1.0
f = reflekt.AsFloat(fi) // 1.2
f = reflekt.AsFloat(bi) // 1.0

// always as string
var s string
s = reflekt.AsString(si) // "1.2"
s = reflekt.AsString(ii) // "1"
s = reflekt.AsString(fi) // "1.200000"
f = reflekt.AsString(bi) // "true"

// always as bool
bi := 1
bb := true
bs1 := "true" // see strconv.ParseBool
bs2 := "True" // see strconv.ParseBool
bs3 := "TRUE" // see strconv.ParseBool
var b bool
b = reflekt.AsBool(bi)  // true
b = reflekt.AsBool(bb)  // true
b = reflekt.AsBool(bs1) // true
b = reflekt.AsBool(bs2) // true
b = reflekt.AsBool(bs3) // true
```

### Casting maps

```go
import "gopkg.in/ukautz/reflekt.v4"

// object oriented
m := map[string]interface{}{"foo": 1}
reflekt.AsIntMap(m) // map[string]int{"foo":1}
reflekt.AsFloatMap(m) // map[string]float64{"foo":1}
reflekt.AsBoolMap(m) // map[string]bool{"foo":true}
reflekt.AsStringMap(m) // map[string]string{"foo":"1"}
```

### Using OO interface

```go
import "gopkg.in/ukautz/reflekt.v4"

// object oriented
v := reflekt.NewValue(1.23)
v.String() // == "1.23"
v.Strings() // == []string{"1.23"}
v.Int() // == int(1)
v.Ints() // == []int{1}
v.Float() // == float64(1.23)
v.Floats() // == []float64{1.23}
v.Bool() // == true
v.Bools() // == []bool{true}

// OO with maps
v := reflekt.NewValue(map[interface{}]interface{}{"foo": 1.23})
v.InterfaceMap() // map[string]interface{}{"foo": 1.23}
v.IntMap() // map[string]int{"foo": 1}
v.FloatMap() // map[string]float64{"foo": 1.23}
v.BoolMap() // map[string]bool{"foo": true}
v.StringMap() // map[string]string{"foo": "1.23"}
```
