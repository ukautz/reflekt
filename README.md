reflekt
=======

This package provides lazy developers shorthands for determining or casting
(to) primitive Go types.

My use it mainly in the context of parsing non-uniform data.

Examples
--------

Determine whether anything is of a primitive type. 

``` go
if reflekt.IsInt(v) {
    fmt.Printf("Yes, %d is a native integer type", v)
} else if reflekt.IsFloat(v) {
    fmt.Printf("Yes, %f is a native float type", v)
}
```

(Try) Cast given value into respective type (tries to exhaust all possibilities):


``` go
si := "1.2"
ii := 1
fi := 1.2
bi := true

// always as int
var i int
i = AsInt(si) // 1
i = AsInt(ii) // 1
i = AsInt(fi) // 1
i = AsInt(bi) // 1

// always as float
var f float64
f = AsFloat(si) // 1.2
f = AsFloat(ii) // 1.0
f = AsFloat(fi) // 1.2

// always as string
var s string
s = AsString(si) // "1.2"
s = AsString(ii) // "1"
s = AsString(fi) // "1.200000"

// always as bool
bi := 1
bb := true
bs1 := "true" // see strconv.ParseBool
bs2 := "True" // see strconv.ParseBool
bs3 := "TRUE" // see strconv.ParseBool
var b bool
b = AsBool(bi)  // true
b = AsBool(bb)  // true
b = AsBool(bs1) // true
b = AsBool(bs2) // true
b = AsBool(bs3) // true

```
