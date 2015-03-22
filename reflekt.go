package reflekt

import (
	"fmt"
	"reflect"
	"strconv"
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
	k := reflect.ValueOf(v).Kind()
	return IsIntKind(k) || IsUintKind(k)
}

// IsFloatKind checks if provided kind is of an float kind
func IsFloatKind(k reflect.Kind) bool {
	return k == reflect.Float32 || k == reflect.Float64
}

// IsFloat checks if value is of any float kind
func IsFloat(v interface{}) bool {
	return IsFloatKind(reflect.ValueOf(v).Kind())
}

// AsInt tries to return or convert the value from anything to int
func AsInt(v interface{}) int {
	r := reflect.ValueOf(v)
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
			if f, e := strconv.ParseFloat(r.String(), 64); e!= nil {
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
	default:
		return 0
	}
}

// AsFloat tries to return or convert the value from anything to float64
func AsFloat(v interface{}) float64 {
	r := reflect.ValueOf(v)
	k := r.Kind()
	switch {
	case IsFloatKind(k):
		return r.Float()
	case k == reflect.String:
		f, _ := strconv.ParseFloat(r.String(), 64)
		return f
	default:
		return float64(AsInt(v))
	}
}

// AsBool tries to return or convert the value from anything to bool
func AsBool(v interface{}) bool {
	r := reflect.ValueOf(v)
	k := r.Kind()
	switch {
	case r.Kind() == reflect.Bool:
		return r.Bool()
	case k == reflect.String:
		b, _ := strconv.ParseBool(r.String())
		return b
	default:
		return AsInt(v) > 0
	}
}

// AsBool tries to return or convert the value from anything to bool
func AsString(v interface{}) string {
	r := reflect.ValueOf(v)
	k := r.Kind()
	switch {
		case k == reflect.String:
		return v.(string)
		case r.Kind() == reflect.Bool:
		return fmt.Sprintf("%v", v)
		case IsInt(v):
		return fmt.Sprintf("%d", AsInt(v))
		case IsFloat(v):
		return fmt.Sprintf("%0.6f", AsFloat(v))
		default:
		return ""
	}
}
