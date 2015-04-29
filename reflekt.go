package reflekt

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

// IsIntKind checks if provided kind is of any unsigned integer kind
func IsUintKind(k reflect.Kind) bool {
	return k == reflect.Uint || k == reflect.Uint8 || k == reflect.Uint16 || k == reflect.Uint32 || k == reflect.Uint64
}

// IsIntKind checks if provided kind is of any signed integer kind
func IsIntKind(k reflect.Kind) bool {
	return k == reflect.Int || k == reflect.Int8 || k == reflect.Int16 || k == reflect.Int32 || k == reflect.Int64
}

// IsInt checks if value is of any (signed or unsigned) integer kind
func IsInt(v interface{}) bool {
	var k reflect.Kind
	switch v.(type) {
	case reflect.Value:
		k = v.(reflect.Value).Kind()
	default:
		k = reflect.ValueOf(v).Kind()
	}
	return IsIntKind(k) || IsUintKind(k)
}

// IsFloatKind checks if provided kind is of an float kind
func IsFloatKind(k reflect.Kind) bool {
	return k == reflect.Float32 || k == reflect.Float64
}

// IsFloat checks if value is of any float kind
func IsFloat(v interface{}) bool {
	var k reflect.Kind
	switch v.(type) {
	case reflect.Value:
		k = v.(reflect.Value).Kind()
	default:
		k = reflect.ValueOf(v).Kind()
	}
	return IsFloatKind(k)
}

// AsInt tries to return or convert the value from anything to int
func AsInt(v interface{}) int {
	r := reflect.ValueOf(v)
	if r.Type().String() == `reflect.Value` {
		r = v.(reflect.Value)
	}
	k := r.Kind()
	switch {
	case IsIntKind(k):
		return int(r.Int())
	case IsUintKind(k):
		return int(r.Uint())
	case IsFloatKind(k):
		return int(r.Float())
	case k == reflect.Bool:
		if r.Bool() {
			return 1
		} else {
			return 0
		}
	case k == reflect.String:
		if i, e := strconv.ParseInt(r.String(), 10, 0); e != nil {
			if f, e := strconv.ParseFloat(r.String(), 64); e != nil {
				if b, _ := strconv.ParseBool(r.String()); b {
					return 1
				} else {
					return 0
				}
			} else {
				return int(f)
			}
		} else {
			return int(i)
		}
	case k == reflect.Interface:
		return AsInt(fmt.Sprintf("%v", r.Interface()))
	default:
		return 0
	}
}

// AsFloat tries to return or convert the value from anything to float64
func AsFloat(v interface{}) float64 {
	r := reflect.ValueOf(v)
	if r.Type().String() == `reflect.Value` {
		r = v.(reflect.Value)
	}
	k := r.Kind()
	switch {
	case IsFloatKind(k):
		return r.Float()
	case k == reflect.String:
		if f, e := strconv.ParseFloat(r.String(), 64); e != nil {
			if b, _ := strconv.ParseBool(r.String()); b {
				return float64(1)
			} else {
				return float64(0)
			}
		} else {
			return f
		}
	case k == reflect.Interface:
		return AsFloat(fmt.Sprintf("%v", r.Interface()))
	default:
		return float64(AsInt(v))
	}
}

// AsBool tries to return or convert the value from anything to bool
func AsBool(v interface{}) bool {
	r := reflect.ValueOf(v)
	if r.Type().String() == `reflect.Value` {
		r = v.(reflect.Value)
	}
	k := r.Kind()
	switch {
	case r.Kind() == reflect.Bool:
		return r.Bool()
	case k == reflect.String:
		if b, e := strconv.ParseBool(r.String()); e != nil {
			return AsFloat(v) > 0
		} else {
			return b
		}
	case k == reflect.Interface:
		return AsBool(fmt.Sprintf("%v", r.Interface()))
	default:
		return AsFloat(v) > 0
	}
}

// AsString tries to return or convert the value from anything to string
func AsString(v interface{}) string {
	r := reflect.ValueOf(v)
	if r.Type().String() == `reflect.Value` {
		r = v.(reflect.Value)
	}
	k := r.Kind()
	switch {
	case k == reflect.String:
		return r.String()
	case k == reflect.Interface:
		return fmt.Sprintf("%v", r.Interface())
	case r.Kind() == reflect.Bool:
		return fmt.Sprintf("%v", r.Bool())
	case IsInt(v):
		return fmt.Sprintf("%d", AsInt(r))
	case IsFloat(v):
		return fmt.Sprintf("%v", AsFloat(r))
	default:
		return ""
	}
}

// AsMap converts given map into other map
func AsMap(v interface{}, key reflect.Type, val reflect.Type, add func(to reflect.Value, key reflect.Value, val reflect.Value)) reflect.Value {
	res := reflect.MakeMap(reflect.MapOf(key, val))
	r := reflect.ValueOf(v)
	switch r.Kind() {
	case reflect.Map:
		for _, k := range r.MapKeys() {
			add(res, k, r.MapIndex(k))
		}
	}
	return res
}

// AsIntMap tries to return any map[interface{}]interface{} as map[string]int.
// Returns nil if v is not a map
func AsIntMap(v interface{}) map[string]int {
	m := AsMap(v, reflect.TypeOf(""), reflect.TypeOf(0), func(to reflect.Value, key reflect.Value, val reflect.Value) {
		k := reflect.ValueOf(AsString(key))
		v := reflect.ValueOf(AsInt(val))
		to.SetMapIndex(k, v)
	})
	return m.Interface().(map[string]int)
}

// AsFloatMap tries to return any map[interface{}]interface{} as map[string]float.
// Returns nil if v is not a map
func AsFloatMap(v interface{}) map[string]float64 {
	m := AsMap(v, reflect.TypeOf(""), reflect.TypeOf(0.0), func(to reflect.Value, key reflect.Value, val reflect.Value) {
		k := reflect.ValueOf(AsString(key))
		v := reflect.ValueOf(AsFloat(val))
		to.SetMapIndex(k, v)
	})
	return m.Interface().(map[string]float64)
}

// AsBoolMap tries to return any map[interface{}]interface{} as map[string]bool.
// Returns nil if v is not a map
func AsBoolMap(v interface{}) map[string]bool {
	m := AsMap(v, reflect.TypeOf(""), reflect.TypeOf(true), func(to reflect.Value, key reflect.Value, val reflect.Value) {
		k := reflect.ValueOf(AsString(key))
		v := reflect.ValueOf(AsBool(val))
		to.SetMapIndex(k, v)
	})
	return m.Interface().(map[string]bool)
}

// AsStringMap tries to return any map[interface{}]interface{} as map[string]string.
// Returns nil if v is not a map
func AsStringMap(v interface{}) map[string]string {
	m := AsMap(v, reflect.TypeOf(""), reflect.TypeOf(""), func(to reflect.Value, key reflect.Value, val reflect.Value) {
		k := reflect.ValueOf(AsString(key))
		v := reflect.ValueOf(AsString(val))
		to.SetMapIndex(k, v)
	})
	return m.Interface().(map[string]string)
}

// AsInterfaceMap tries to return any map[interface{}]interface{} as map[string]interface{}.
// Returns nil if v is not a map
func AsInterfaceMap(v interface{}) map[string]interface{} {
	i := reflect.TypeOf((*interface{})(nil)).Elem()
	m := AsMap(v, reflect.TypeOf(""), i, func(to reflect.Value, key reflect.Value, val reflect.Value) {
		k := reflect.ValueOf(AsString(key))
		to.SetMapIndex(k, val)
	})
	return m.Interface().(map[string]interface{})
}


func structElemAs(f reflect.Value, lc bool) interface{} {
	for f.Kind() == reflect.Ptr || f.Kind() == reflect.Interface {
		f = f.Elem()
	}
	switch f.Kind() {
	case reflect.Struct:
		return StructAsMap(f.Interface(), lc)
	case reflect.Slice:
		s := make([]interface{}, f.Len())
		for i := 0; i < f.Len(); i++ {
			x := f.Index(i)
			for x.Kind() == reflect.Ptr || x.Kind() == reflect.Interface {
				x = x.Elem()
			}
			s[i] = structElemAs(x, lc)
		}
		return s
	default:
		return f.Interface()
	}
}

// StructAsMap converts given struct into `map[string]interface{}`
func StructAsMap(v interface{}, lowerCase ...bool) map[string]interface{} {
	lc := false
	if len(lowerCase) > 0 && lowerCase[0] {
		lc = true
	}
	res := make(map[string]interface{})
	r := reflect.ValueOf(v)
	for r.Kind() == reflect.Ptr {
		r = r.Elem()
	}
	switch r.Kind() {
	case reflect.Struct:
		t := r.Type()
		for i := 0; i < r.NumField(); i++ {
			f := r.Field(i)
			n := t.Field(i).Name
			if lc {
				n = strings.ToLower(n)
			}
			res[n] = structElemAs(f, lc)
		}
	}
	return res
}

func fillStruct(s interface{}, d map[string]interface{}, p string) error {
	r := reflect.ValueOf(s)
	prefix := p
	if prefix != "" {
		prefix = prefix + " "
	}
	fmt.Printf("Starting with %s (%s)\n", r.Kind(), prefix)
	w := []string{}
	for r.Kind() == reflect.Ptr || r.Kind() == reflect.Interface {
		w = append(w, r.Kind().String())
		r = r.Elem()
		fmt.Printf(" > Hopping to %s\n", r.Kind())
	}
	fmt.Printf("Using now %s (%s)\n", r.Kind(), prefix)
	if r.Kind() != reflect.Struct {
		return fmt.Errorf(prefix+ "Expected (ptr|interface)+ -> struct, got %s -> %s", strings.Join(w, " -> "), r.Kind())
	}
	t := r.Type()
	if r.Kind() == reflect.Interface {
		t = r.Elem().Type()
	}

	for i := 0; i < t.NumField(); i++ {
		fv := r.Field(i)
		ft := t.Field(i)
		if !fv.CanSet() {
			fmt.Printf("- Cannot set %s\n", ft.Name)
			continue
		}
		for _, n := range []string{ft.Name, strings.ToLower(ft.Name)} {
			if v, ok := d[n]; ok {
				fk := fv.Kind()
				vv := reflect.ValueOf(v)
				if IsIntKind(fk) {
					fv.SetInt(int64(AsInt(v)))
				} else if IsFloatKind(fk) {
					fv.SetFloat(AsFloat(v))
				} else if fk == reflect.Bool {
					fv.SetBool(AsBool(v))
				} else if fk == reflect.String {
					fv.SetString(AsString(v))
				} else if fk == reflect.Struct || fk == reflect.Ptr || fk == reflect.Interface {
					if (fk == reflect.Ptr || fk == reflect.Interface) && fv.IsNil() {
						fmt.Printf("... NIL!\n")
					}
					if vv.Kind() == reflect.Map {
						var err error
						var sub reflect.Value
						if fk == reflect.Struct {
							fmt.Printf("  SUB IS STRUCT\n");
							sub = reflect.New(fv.Type())
						} else {
							fmt.Printf("  SUB IS %s (%s) (%s)\n", fk, ft.Type, fv.Type().Elem());
							sub = reflect.New(fv.Type().Elem())
						}
						err = fillStruct(sub.Interface(), AsInterfaceMap(v), p+ n+ ":")
						if err != nil {
							return err
						}
						if fk == reflect.Struct {
							fv.Set(sub)
						} else {
							fv.Set(sub.Addr())
						}
					} else {
						return fmt.Errorf(prefix+ "Cannot fill sub-struct %s from %s", n, vv.Kind())
					}
				} else if fk == vv.Kind() {
					fv.Set(vv)
				} else {
					return fmt.Errorf(prefix+ "Cannot fill %s (%s) from %s", n, fk, vv.Kind())
				}
				continue
			}
		}
	}
	return nil
}

func FillStruct(s interface{}, d map[string]interface{}) error {
	return fillStruct(s, d, "")
}


// MergeMaps takes arbitrary maps of the same type and merges them into a new one
// TODO: does not fit here -> new package!
func MergeMaps(v ...interface{}) (interface{}, error) {
	expect := ""
	var res reflect.Value
	for i, m := range v {
		r := reflect.ValueOf(m)
		if r.Kind() != reflect.Map {
			return nil, fmt.Errorf("Parameter %d is not a map but %s", i, r.Kind())
		}
		if expect == "" {
			expect = r.Type().String()
			res = reflect.New(r.Type())
		} else if r.Type().String() != expect {
			return nil, fmt.Errorf("All maps must be of the same type. Expect %s, got %s (parameter: %d)", expect, r.Type(), i)
		}
		for _, k := range r.MapKeys() {
			res.SetMapIndex(k, r.MapIndex(k))
		}
	}

	return res.Interface(), nil
}




