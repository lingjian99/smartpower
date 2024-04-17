package encodingx

import (
	"math"
	"reflect"
	"sort"
	"strconv"
	"sync"
)

var blTag = "bl"

func Unmarshal(data []byte, v interface{}) error {
	// Fast path for basic types and slices.
	if n := intDataSize(v); n != 0 {
		bs := data
		if n > len(data) {
			bs = make([]byte, n)
			copy(bs, data)
		}

		switch vv := v.(type) {
		case *bool:
			*vv = bs[0] != 0
		case *int8:
			*vv = int8(bs[0])
		case *uint8:
			*vv = bs[0]
		case *int16:
			*vv = int16(byteOrder.Uint16(bs))
		case *uint16:
			*vv = byteOrder.Uint16(bs)
		case *int32:
			*vv = int32(byteOrder.Uint32(bs))
		case *uint32:
			*vv = byteOrder.Uint32(bs)
		case *int64:
			*vv = int64(byteOrder.Uint64(bs))
		case *uint64:
			*vv = byteOrder.Uint64(bs)
		case *float32:
			*vv = math.Float32frombits(byteOrder.Uint32(bs))
		case *float64:
			*vv = math.Float64frombits(byteOrder.Uint64(bs))
		case []bool:
			for i, x := range bs { // Easier to loop over the input for 8-bit values.
				vv[i] = x != 0
			}
		case []int8:
			for i, x := range bs {
				vv[i] = int8(x)
			}
		case []uint8:
			copy(vv, bs)
		case []int16:
			for i := range vv {
				vv[i] = int16(byteOrder.Uint16(bs[2*i:]))
			}
		case []uint16:
			for i := range vv {
				vv[i] = byteOrder.Uint16(bs[2*i:])
			}
		case []int32:
			for i := range vv {
				vv[i] = int32(byteOrder.Uint32(bs[4*i:]))
			}
		case []uint32:
			for i := range vv {
				vv[i] = byteOrder.Uint32(bs[4*i:])
			}
		case []int64:
			for i := range vv {
				vv[i] = int64(byteOrder.Uint64(bs[8*i:]))
			}
		case []uint64:
			for i := range vv {
				vv[i] = byteOrder.Uint64(bs[8*i:])
			}
		case []float32:
			for i := range vv {
				vv[i] = math.Float32frombits(byteOrder.Uint32(bs[4*i:]))
			}
		case []float64:
			for i := range vv {
				vv[i] = math.Float64frombits(byteOrder.Uint64(bs[8*i:]))
			}
		default:
			n = 0 // fast path doesn't apply
		}
		if n != 0 {
			return nil
		}
	}

	rv := reflect.ValueOf(v)
	if rv.Kind() != reflect.Pointer || rv.IsNil() {
		return &InvalidUnmarshalError{reflect.TypeOf(v)}
	}
	if rv.Kind() == reflect.Pointer {
		rv = rv.Elem()
	}

	var d decoder
	d.init(data)

	// Fallback to reflect-based decoding.
	//rv := reflect.ValueOf(v)
	//size := -1
	//switch rv.Kind() {
	//case reflect.Pointer:
	//	rv = rv.Elem()
	//	size = dataSize(rv)
	//case reflect.Slice:
	//}
	//if size < 0 {
	//	return errors.New("binary.Read: invalid type " + reflect.TypeOf(v).String())
	//}
	//d := &decoder{order: byteOrder, buf:bs}
	//
	d.value(rv)
	return nil
}

type InvalidUnmarshalError struct {
	Type reflect.Type
}

func (e *InvalidUnmarshalError) Error() string {
	if e.Type == nil {
		return "encodingx: Unmarshal(nil)"
	}

	if e.Type.Kind() != reflect.Pointer {
		return "encodingx: Unmarshal(non-pointer " + e.Type.String() + ")"
	}
	return "encodingx: Unmarshal(nil " + e.Type.String() + ")"
}

type decoder struct {
	data   []byte
	offset int // next read offset in data
}

func (d *decoder) eof() bool {
	return d.offset >= len(d.data)
}

// readIndex returns the position of the last byte read.
func (d *decoder) readIndex() int {
	return d.offset - 1
}

func (d *decoder) init(data []byte) *decoder {
	d.data = data
	d.offset = 0
	return d
}

func (d *decoder) bool() bool {
	if d.offset < len(d.data) {
		x := d.data[d.offset]
		d.offset++
		return x != 0
	}
	return false
}

func (d *decoder) uint8() uint8 {
	if d.offset < len(d.data) {
		x := d.data[d.offset]
		d.offset++
		return x
	}
	return 0
}

func (d *decoder) uint16() uint16 {
	if d.offset+2 <= len(d.data) {
		x := byteOrder.Uint16(d.data[d.offset : d.offset+2])
		d.offset += 2
		return x
	}
	return 0
}

func (d *decoder) uint32() uint32 {
	if d.offset+4 <= len(d.data) {
		x := byteOrder.Uint32(d.data[d.offset : d.offset+4])
		d.offset += 4
		return x
	}
	return 0
}

func (d *decoder) uint64() uint64 {
	if d.offset+8 <= len(d.data) {
		x := byteOrder.Uint64(d.data[d.offset : d.offset+8])
		d.offset += 8
		return x
	}
	return 0
}

func (d *decoder) int8() int8 { return int8(d.uint8()) }

func (d *decoder) int16() int16 { return int16(d.uint16()) }

func (d *decoder) int32() int32 { return int32(d.uint32()) }

func (d *decoder) int64() int64 { return int64(d.uint64()) }

func (d *decoder) value(v reflect.Value) {
	if d.offset >= len(d.data) {
		return
	}

	switch v.Kind() {
	case reflect.Array:
		l := v.Len()
		for i := 0; i < l; i++ {
			d.value(v.Index(i))
		}

	case reflect.Struct:
		t := v.Type()
		fields := cachedTypeFields(t)
		//subv := v
		for i := range fields.list {
			f := &fields.list[i]
			// Find the nested struct field by following f.index.
			fv := v
			for _, i := range f.index {
				if fv.Kind() == reflect.Pointer {
					if fv.IsNil() {
						fv.Set(reflect.New(v.Type().Elem()))
					}
					fv = fv.Elem()
				}
				fv = fv.Field(i)
			}

			if f.byteLen != 0 {
				d.valueLen(fv, int(f.byteLen))
			} else if fv.Kind() == reflect.Slice {
				// 将余下数据写入 []byte
				d.valueLen(fv, len(d.data)-d.offset)
			} else {
				d.value(fv)
			}
		}
	case reflect.Slice:
		l := v.Len()
		for i := 0; i < l; i++ {
			d.value(v.Index(i))
		}

	case reflect.Bool:
		v.SetBool(d.bool())

	case reflect.Int8:
		v.SetInt(int64(d.int8()))
	case reflect.Int16:
		v.SetInt(int64(d.int16()))
	case reflect.Int32:
		v.SetInt(int64(d.int32()))
	case reflect.Int64:
		v.SetInt(d.int64())

	case reflect.Uint8:
		v.SetUint(uint64(d.uint8()))
	case reflect.Uint16:
		v.SetUint(uint64(d.uint16()))
	case reflect.Uint32:
		v.SetUint(uint64(d.uint32()))
	case reflect.Uint64:
		v.SetUint(d.uint64())

	case reflect.Float32:
		v.SetFloat(float64(math.Float32frombits(d.uint32())))
	case reflect.Float64:
		v.SetFloat(math.Float64frombits(d.uint64()))

	case reflect.Complex64:
		v.SetComplex(complex(
			float64(math.Float32frombits(d.uint32())),
			float64(math.Float32frombits(d.uint32())),
		))
	case reflect.Complex128:
		v.SetComplex(complex(
			math.Float64frombits(d.uint64()),
			math.Float64frombits(d.uint64()),
		))
	}
}

func (d *decoder) valueLen(v reflect.Value, byteLen int) {
	// 处理 String and Slice
	if d.offset >= len(d.data) {
		return
	}
	t := v.Type()

	switch t.Kind() {
	case reflect.String:
		if d.offset+byteLen <= len(d.data) {
			s := string(d.data[d.offset : d.offset+byteLen])
			v.SetString(s)
			d.offset += byteLen
		} else {
			s := string(d.data[d.offset:])
			v.SetString(s)
			d.offset = len(d.data)
		}
	case reflect.Slice:
		if t.Elem().Kind() == reflect.Uint8 {
			if v.Cap() < byteLen {
				v.Set(reflect.MakeSlice(v.Type(), byteLen, byteLen))
			}
		}
		for i := 0; i < v.Len(); i++ {
			d.value(v.Index(i))
		}
	}
}

type structFields struct {
	list []field
}

// A field represents a single field found in a struct.
type field struct {
	name string

	index []int
	typ   reflect.Type

	byteLen int64
}

// byIndex sorts field by index sequence.
type byIndex []field

func (x byIndex) Len() int { return len(x) }

func (x byIndex) Swap(i, j int) { x[i], x[j] = x[j], x[i] }

func (x byIndex) Less(i, j int) bool {
	for k, xik := range x[i].index {
		if k >= len(x[j].index) {
			return false
		}
		if xik != x[j].index[k] {
			return xik < x[j].index[k]
		}
	}
	return len(x[i].index) < len(x[j].index)
}

// typeFields returns a list of fields that JSON should recognize for the given type.
// The algorithm is breadth-first search over the set of structs to include - the top struct
// and then any reachable anonymous structs.
func typeFields(t reflect.Type) structFields {
	// Anonymous fields to explore at the current level and the next.
	current := []field{}
	next := []field{{typ: t}}

	// Count of queued names for current level and the next.
	var count, nextCount map[reflect.Type]int

	// Types already visited at an earlier level.
	visited := map[reflect.Type]bool{}

	// Fields found.
	var fields []field

	for len(next) > 0 {
		current, next = next, current[:0]
		count, nextCount = nextCount, map[reflect.Type]int{}

		for _, f := range current {
			if visited[f.typ] {
				continue
			}
			visited[f.typ] = true

			// Scan f.typ for fields to include.
			for i := 0; i < f.typ.NumField(); i++ {
				sf := f.typ.Field(i)
				if sf.Anonymous {
					t := sf.Type
					if t.Kind() == reflect.Pointer {
						t = t.Elem()
					}
					if !sf.IsExported() && t.Kind() != reflect.Struct {
						// Ignore embedded fields of unexported non-struct types.
						continue
					}
					// Do not ignore embedded fields of unexported struct types
					// since they may have exported fields.
				} else if !sf.IsExported() {
					// Ignore unexported non-embedded fields.
					continue
				}

				tag := sf.Tag.Get(blTag)
				if tag == "-" {
					continue
				}
				blen, _ := strconv.ParseInt(tag, 10, 0)
				index := make([]int, len(f.index)+1)
				copy(index, f.index)
				index[len(f.index)] = i

				ft := sf.Type
				if ft.Name() == "" && ft.Kind() == reflect.Pointer {
					// Follow pointer.
					ft = ft.Elem()
				}

				// Record found field and index sequence.
				if !sf.Anonymous || ft.Kind() != reflect.Struct {
					field := field{
						name:    sf.Name,
						index:   index,
						typ:     ft,
						byteLen: blen,
					}

					fields = append(fields, field)
					if count[f.typ] > 1 {
						// If there were multiple instances, add a second,
						// so that the annihilation code will see a duplicate.
						// It only cares about the distinction between 1 or 2,
						// so don't bother generating any more copies.
						fields = append(fields, fields[len(fields)-1])
					}
					continue
				}

				// Record new anonymous struct to explore in next round.
				nextCount[ft]++
				if nextCount[ft] == 1 {
					next = append(next, field{name: ft.Name(), index: index, typ: ft})
				}
			}
		}
	}
	sort.Sort(byIndex(fields))

	//nameIndex := make(map[string]int, len(fields))
	//for i, field := range fields {
	//	nameIndex[field.name] = i
	//}
	return structFields{fields}
}

var fieldCache sync.Map // map[reflect.Type]structFields

// cachedTypeFields is like typeFields but uses a cache to avoid repeated work.
func cachedTypeFields(t reflect.Type) structFields {
	if f, ok := fieldCache.Load(t); ok {
		return f.(structFields)
	}
	f, _ := fieldCache.LoadOrStore(t, typeFields(t))
	return f.(structFields)
}
