package reflekt

import (
	"fmt"
	"reflect"
	"strings"
)

func structElemAs(f reflect.Value, lc bool, m map[string]interface{}) interface{} {
	for f.Kind() == reflect.Ptr || f.Kind() == reflect.Interface {
		f = f.Elem()
	}
	switch f.Kind() {
	case reflect.Struct:
		return structAsMap(f.Interface(), lc, m)
	case reflect.Slice:
		s := make([]interface{}, f.Len())
		for i := 0; i < f.Len(); i++ {
			x := f.Index(i)
			for x.Kind() == reflect.Ptr || x.Kind() == reflect.Interface {
				x = x.Elem()
			}
			s[i] = structElemAs(x, lc, nil)
		}
		return s
	default:
		return f.Interface()
	}
}

func structAsMap(v interface{}, sc bool, res map[string]interface{}) map[string]interface{} {
	if res == nil {
		res = make(map[string]interface{})
	}
	r := reflect.ValueOf(v)
	k := r.Kind()
	for k == reflect.Ptr || k == reflect.Interface {
		r = r.Elem()
		k = r.Kind()
	}
	switch k {
	case reflect.Struct:
		t := r.Type()
		for i := 0; i < r.NumField(); i++ {
			fv := r.Field(i)
			ft := t.Field(i)
			if ft.PkgPath != "" {
				continue
			}
			n := ft.Name
			if sc {
				n = snakeCase(n)
			}
			if fv.Kind() == reflect.Struct && ft.Anonymous {
				structAsMap(fv.Interface(), sc, res)
			} else {
				res[n] = structElemAs(fv, sc, nil)
			}
		}
	}
	return res
}

// StructAsMap converts given struct into `map[string]interface{}`
func StructAsMap(v interface{}, snakeCase ...bool) map[string]interface{} {
	return structAsMap(v, len(snakeCase) > 0 && snakeCase[0], nil)
}

// StructFiller is going to be an awesome tool to simply fill structs from maps (thereby JSON and all that).. just not yet
type StructFiller struct {
	m map[reflect.Type]func(v interface{}) reflect.Type
}

// NewStructFiller generates new filler.. DONT USE IT, YET
func NewStructFiller() *StructFiller {
	return &StructFiller{
		m: make(map[reflect.Type]func(v interface{}) reflect.Type),
	}
}

// Register assigns concrete structs to interfaces.. since interface-type values cannot be created
func (this *StructFiller) Register(iface reflect.Type, determine func(v interface{}) reflect.Type) *StructFiller {
	this.m[iface] = determine
	return this
}

func (this *StructFiller) set(fv reflect.Value, ft reflect.StructField, d map[string]interface{}, p, prefix string) error {
	if !fv.CanSet() {
		fmt.Printf("- Cannot set %s\n", ft.Name)
		return nil
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
			} else if fk == reflect.Struct {
				if ft.Anonymous {
					// TODO: ..
				} else if vv.Kind() == reflect.Map {
					sub := reflect.New(fv.Type())
					this.fill(sub.Interface(), AsInterfaceMap(v), p+n+":")
					fv.Set(sub.Elem())
				} else {
					return fmt.Errorf(prefix+"Cannot fill sub-struct %s (%s) from %s", n, fk, vv.Kind())
				}
			} else if fk == reflect.Ptr {
				if vv.Kind() == reflect.Map {
					sub := reflect.New(fv.Type().Elem())
					this.fill(sub.Interface(), AsInterfaceMap(v), p+n+":")
					fv.Set(sub)
				} else {
					return fmt.Errorf(prefix+"Cannot fill sub-struct ptr %s (%s) from %s", n, fk, vv.Kind())
				}
			} else if fk == reflect.Interface {
				if cast, ok := this.m[fv.Type()]; !ok {
					return fmt.Errorf(prefix+"Not found registererd cast for interface %s for %s", fv.Kind(), n)
				} else if vv.Kind() == reflect.Map {
					sub := reflect.New(cast(v))
					this.fill(sub.Interface(), AsInterfaceMap(v), p+n+":")
					fv.Set(sub)
				} else {
					return fmt.Errorf(prefix+"Cannot fill sub-struct ptr %s (%s) from %s", n, fk, vv.Kind())
				}
			} else if fk == reflect.Slice {
				if vv.Kind() != reflect.Slice {
					return fmt.Errorf(prefix+"Cannot fill slice %s (%s) from %s", n, fk, vv.Kind())
				}
				pt := ft.Type.Elem()
				var et reflect.Type
				if pt.Kind() == reflect.Interface {
					if cast, ok := this.m[pt]; !ok {
						return fmt.Errorf(prefix+"Not found registererd cast for interface %s for slice %s", fv.Kind(), n)
					} else {
						et = cast(v)
					}
				} else {
					et = pt
					for et.Kind() == reflect.Ptr {
						et = et.Elem()
					}
				}
				st := reflect.SliceOf(et)
				l := fv.Len()
				reflect.MakeSlice(st, l, l)
				if et.Kind() == reflect.Struct {
					if vv.Kind() != reflect.Map {
						return fmt.Errorf(prefix+"Cannot fill slice of %s with for slice %s", et, vv.Kind(), n)
					} else {
						for i := 0; i < l; i++ {
							e := reflect.New(et)
							this.fill(e.Interface(), AsInterfaceMap(v), fmt.Sprintf("%s%s.%d:", p, n, i))
							fv.Field(i).Set(e)
						}
					}
				} else {
					for i := 0; i < l; i++ {
						e := reflect.New(et)
						fv.Field(i).Set(e)
					}
				}

				//s := reflect.MakeSlice()
				//fv.Set(vv)
			} else if fk == vv.Kind() {
				fv.Set(vv)
			} else {
				return fmt.Errorf(prefix+"Cannot fill %s (%s) from %s", n, fk, vv.Kind())
			}
			return nil
		}
	}

	return nil
}

func (this *StructFiller) fill(s interface{}, d map[string]interface{}, p string) error {
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
		return fmt.Errorf(prefix+"Expected (ptr|interface)+ -> struct, got %s -> %s", strings.Join(w, " -> "), r.Kind())
	}
	t := r.Type()
	if r.Kind() == reflect.Interface {
		t = r.Elem().Type()
	}

	for i := 0; i < t.NumField(); i++ {
		fv := r.Field(i)
		ft := t.Field(i)
		this.set(fv, ft, d, p, prefix)

	}
	return nil
}

// Fill takes map and populates the struct..
func (this *StructFiller) Fill(s interface{}, d map[string]interface{}) error {
	return this.fill(s, d, "")
}