package encodingx

import (
	"bytes"
	"encoding/binary"
	"math"
	"reflect"
	"sync"
)

var (
	byteOrder = binary.BigEndian
)

func Marshal(data interface{}) ([]byte, error) {
	// Fast path for basic types and slices.
	if n := intDataSize(data); n != 0 {
		bs := make([]byte, n)
		switch v := data.(type) {
		case *bool:
			if *v {
				bs[0] = 1
			} else {
				bs[0] = 0
			}
		case bool:
			if v {
				bs[0] = 1
			} else {
				bs[0] = 0
			}
		case []bool:
			for i, x := range v {
				if x {
					bs[i] = 1
				} else {
					bs[i] = 0
				}
			}
		case *int8:
			bs[0] = byte(*v)
		case int8:
			bs[0] = byte(v)
		case []int8:
			for i, x := range v {
				bs[i] = byte(x)
			}
		case *uint8:
			bs[0] = *v
		case uint8:
			bs[0] = v
		case []uint8:
			bs = v
		case *int16:
			byteOrder.PutUint16(bs, uint16(*v))
		case int16:
			byteOrder.PutUint16(bs, uint16(v))
		case []int16:
			for i, x := range v {
				byteOrder.PutUint16(bs[2*i:], uint16(x))
			}
		case *uint16:
			byteOrder.PutUint16(bs, *v)
		case uint16:
			byteOrder.PutUint16(bs, v)
		case []uint16:
			for i, x := range v {
				byteOrder.PutUint16(bs[2*i:], x)
			}
		case *int32:
			byteOrder.PutUint32(bs, uint32(*v))
		case int32:
			byteOrder.PutUint32(bs, uint32(v))
		case []int32:
			for i, x := range v {
				byteOrder.PutUint32(bs[4*i:], uint32(x))
			}
		case *uint32:
			byteOrder.PutUint32(bs, *v)
		case uint32:
			byteOrder.PutUint32(bs, v)
		case []uint32:
			for i, x := range v {
				byteOrder.PutUint32(bs[4*i:], x)
			}
		case *int64:
			byteOrder.PutUint64(bs, uint64(*v))
		case int64:
			byteOrder.PutUint64(bs, uint64(v))
		case []int64:
			for i, x := range v {
				byteOrder.PutUint64(bs[8*i:], uint64(x))
			}
		case *uint64:
			byteOrder.PutUint64(bs, *v)
		case uint64:
			byteOrder.PutUint64(bs, v)
		case []uint64:
			for i, x := range v {
				byteOrder.PutUint64(bs[8*i:], x)
			}
		case *float32:
			byteOrder.PutUint32(bs, math.Float32bits(*v))
		case float32:
			byteOrder.PutUint32(bs, math.Float32bits(v))
		case []float32:
			for i, x := range v {
				byteOrder.PutUint32(bs[4*i:], math.Float32bits(x))
			}
		case *float64:
			byteOrder.PutUint64(bs, math.Float64bits(*v))
		case float64:
			byteOrder.PutUint64(bs, math.Float64bits(v))
		case []float64:
			for i, x := range v {
				byteOrder.PutUint64(bs[8*i:], math.Float64bits(x))
			}
		}
		return bs, nil
	}

	v := reflect.Indirect(reflect.ValueOf(data))
	e := newEncoder()
	e.value(v)

	buf := append([]byte(nil), e.Bytes()...)
	encoderPool.Put(e)
	return buf, nil
}

var encoderPool sync.Pool

type encoder struct {
	bytes.Buffer
	scratch [64]byte
}

func newEncoder() *encoder {
	if v := encoderPool.Get(); v != nil {
		e := v.(*encoder)
		e.Reset()
		return e
	}
	return &encoder{}
}

func (e *encoder) bool(x bool) {
	if x {
		e.WriteByte(1)
	} else {
		e.WriteByte(0)
	}
}

func (e *encoder) uint8(x uint8) {
	e.WriteByte(x)
}

func (e *encoder) uint16(x uint16) {
	bs := e.scratch[:2]
	byteOrder.PutUint16(bs, x)
	e.Write(bs)
}

func (e *encoder) uint32(x uint32) {
	bs := e.scratch[:4]
	byteOrder.PutUint32(bs, x)
	e.Write(bs)
}

func (e *encoder) uint64(x uint64) {
	bs := e.scratch[:8]
	byteOrder.PutUint64(bs, x)
	e.Write(bs)
}

func (e *encoder) int8(x int8) { e.uint8(uint8(x)) }

func (e *encoder) int16(x int16) { e.uint16(uint16(x)) }

func (e *encoder) int32(x int32) { e.uint32(uint32(x)) }

func (e *encoder) int64(x int64) { e.uint64(uint64(x)) }

func (e *encoder) value(v reflect.Value) {
	switch v.Kind() {
	case reflect.Array:
		l := v.Len()
		for i := 0; i < l; i++ {
			e.value(v.Index(i))
		}
	case reflect.Struct:
		//t := v.Type()
		l := v.NumField()
		for i := 0; i < l; i++ {
			// see comment for corresponding code in decoder.value()
			//if v := v.Field(i); v.CanSet() {
			//	e.value(v)
			//}
			e.value(v.Field(i))
		}

	case reflect.Slice:
		l := v.Len()
		for i := 0; i < l; i++ {
			e.value(v.Index(i))
		}

	case reflect.Bool:
		e.bool(v.Bool())

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		switch v.Type().Kind() {
		case reflect.Int8:
			e.int8(int8(v.Int()))
		case reflect.Int16:
			e.int16(int16(v.Int()))
		case reflect.Int32:
			e.int32(int32(v.Int()))
		case reflect.Int64:
			e.int64(v.Int())
		}

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		switch v.Type().Kind() {
		case reflect.Uint8:
			e.uint8(uint8(v.Uint()))
		case reflect.Uint16:
			e.uint16(uint16(v.Uint()))
		case reflect.Uint32:
			e.uint32(uint32(v.Uint()))
		case reflect.Uint64:
			e.uint64(v.Uint())
		}

	case reflect.Float32, reflect.Float64:
		switch v.Type().Kind() {
		case reflect.Float32:
			e.uint32(math.Float32bits(float32(v.Float())))
		case reflect.Float64:
			e.uint64(math.Float64bits(v.Float()))
		}

	case reflect.Complex64, reflect.Complex128:
		switch v.Type().Kind() {
		case reflect.Complex64:
			x := v.Complex()
			e.uint32(math.Float32bits(float32(real(x))))
			e.uint32(math.Float32bits(float32(imag(x))))
		case reflect.Complex128:
			x := v.Complex()
			e.uint64(math.Float64bits(real(x)))
			e.uint64(math.Float64bits(imag(x)))
		}
	}
}

// intDataSize returns the size of the data required to represent the data when encoded.
// It returns zero if the type cannot be implemented by the fast path in Read or Write.
func intDataSize(data any) int {
	switch data := data.(type) {
	case bool, int8, uint8, *bool, *int8, *uint8:
		return 1
	case []bool:
		return len(data)
	case []int8:
		return len(data)
	case []uint8:
		return len(data)
	case int16, uint16, *int16, *uint16:
		return 2
	case []int16:
		return 2 * len(data)
	case []uint16:
		return 2 * len(data)
	case int32, uint32, *int32, *uint32:
		return 4
	case []int32:
		return 4 * len(data)
	case []uint32:
		return 4 * len(data)
	case int64, uint64, *int64, *uint64:
		return 8
	case []int64:
		return 8 * len(data)
	case []uint64:
		return 8 * len(data)
	case float32, *float32:
		return 4
	case float64, *float64:
		return 8
	case []float32:
		return 4 * len(data)
	case []float64:
		return 8 * len(data)
	}
	return 0
}
