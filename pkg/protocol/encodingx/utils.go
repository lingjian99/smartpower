package encodingx

import (
	"fmt"
	"math"
	"time"
)

func ReadChars(offset int, buf []byte, size int, charSize int, str *string) bool {
	if offset+charSize > size {
		return false
	}
	*str = string(buf[offset : offset+charSize])
	return true
}

func Value(offset int, buf []byte, size int, v interface{}) bool {
	n := intDataSize(v)
	if n == 0 {
		return false
	}
	if offset+n > size {
		return false
	}

	switch vv := v.(type) {
	case *bool:
		*vv = buf[offset] != 0
	case *int8:
		*vv = int8(buf[offset])
	case *uint8:
		*vv = buf[offset]
	case *int16:
		*vv = int16(byteOrder.Uint16(buf[offset:]))
	case *uint16:
		*vv = byteOrder.Uint16(buf[offset:])
	case *int32:
		*vv = int32(byteOrder.Uint32(buf[offset:]))
	case *uint32:
		*vv = byteOrder.Uint32(buf[offset:])
	case *int64:
		*vv = int64(byteOrder.Uint64(buf[offset:]))
	case *uint64:
		*vv = byteOrder.Uint64(buf[offset:])
	case *float32:
		*vv = math.Float32frombits(byteOrder.Uint32(buf[offset:]))
	case *float64:
		*vv = math.Float64frombits(byteOrder.Uint64(buf[offset:]))
	case []bool:
		for i := 0; i < n; i++ {
			vv[i] = buf[offset+i] != 0
		}

	case []int8:
		for i := 0; i < n; i++ {
			vv[i] = int8(buf[offset+i])
		}
	case []uint8:
		copy(vv, buf[offset:])
	case []int16:
		for i := 0; i < n; i++ {
			vv[i] = int16(byteOrder.Uint16(buf[offset+2*i:]))
		}
	case []uint16:
		for i := 0; i < n; i++ {
			vv[i] = byteOrder.Uint16(buf[offset+2*i:])
		}
	case []int32:
		for i := 0; i < n; i++ {
			vv[i] = int32(byteOrder.Uint32(buf[offset+4*i:]))
		}
	case []uint32:
		for i := 0; i < n; i++ {
			vv[i] = byteOrder.Uint32(buf[offset+4*i:])
		}
	case []int64:
		for i := 0; i < n; i++ {
			vv[i] = int64(byteOrder.Uint64(buf[offset+8*i:]))
		}
	case []uint64:
		for i := 0; i < n; i++ {
			vv[i] = byteOrder.Uint64(buf[offset+8*i:])
		}
	case []float32:
		for i := 0; i < n; i++ {
			vv[i] = math.Float32frombits(byteOrder.Uint32(buf[offset+4*i:]))
		}
	case []float64:
		for i := 0; i < n; i++ {
			vv[i] = math.Float64frombits(byteOrder.Uint64(buf[offset+8*i:]))
		}
	default:
		return false
	}

	return true
}

func Int32(offset int, buf []byte, size int, v *int32) bool {
	if offset+4 > size {
		return false
	}
	*v = int32(byteOrder.Uint32(buf[offset:]))
	return true
}

func Int32ToInt(offset int, buf []byte, size int, v *int) bool {
	if offset+4 > size {
		return false
	}
	*v = int(int32(byteOrder.Uint32(buf[offset:])))
	return true
}

func UInt32ToInt(offset int, buf []byte, size int, v *int) bool {
	if offset+4 > size {
		return false
	}
	*v = int(byteOrder.Uint32(buf[offset:]))
	return true
}

func Int16ToInt(offset int, buf []byte, size int, v *int) bool {
	if offset+2 > size {
		return false
	}
	*v = int(int16(byteOrder.Uint16(buf[offset:])))
	return true
}

func UInt16ToInt(offset int, buf []byte, size int, v *int) bool {
	if offset+2 > size {
		return false
	}
	*v = int(byteOrder.Uint16(buf[offset:]))
	return true
}

func UInt16ToInt64(offset int, buf []byte, size int, v *int64) bool {
	if offset+2 > size {
		return false
	}
	*v = int64(byteOrder.Uint16(buf[offset:]))
	return true
}

func Int8ToFloat(offset int, buf []byte, size int, scale float64, v *float64) bool {
	if offset+1 > size {
		return false
	}

	*v = float64(int8(buf[offset])) * scale
	return true
}

func Uint8ToFloat(offset int, buf []byte, size int, scale float64, v *float64) bool {
	if offset+1 > size {
		return false
	}

	*v = float64(buf[offset]) * scale
	return true
}

func Int16ToFloat(offset int, buf []byte, size int, scale float64, v *float64) bool {
	if offset+2 > size {
		return false
	}

	*v = float64(int16(byteOrder.Uint16(buf[offset:offset+2]))) * scale
	return true
}
func UInt16ToFloat(offset int, buf []byte, size int, scale float64, v *float64) bool {
	if offset+2 > size {
		return false
	}

	*v = float64(byteOrder.Uint16(buf[offset:offset+2])) * scale
	return true
}

func UInt32ToFloat(offset int, buf []byte, size int, scale float64, v *float64) bool {
	if offset+4 > size {
		return false
	}

	*v = float64(byteOrder.Uint32(buf[offset:])) * scale
	return true
}

func Int32ToFloat(offset int, buf []byte, size int, scale float64, v *float64) bool {
	if offset+4 > size {
		return false
	}

	*v = float64(int32(byteOrder.Uint32(buf[offset:]))) * scale
	return true
}

func Int64ToFloat(offset int, buf []byte, size int, scale float64, v *float64) bool {
	if offset+8 > size {
		return false
	}

	*v = float64(int64(byteOrder.Uint64(buf[offset:]))) * scale
	return true
}

func DateTimeStr(offset int, buf []byte, size int, str *string) bool {
	if offset+6 > size {
		return false
	}
	*str = fmt.Sprintf("%d-%02d-%02d %02d:%02d:%02d", 2000+int(buf[offset]), buf[offset+1], buf[offset+2], buf[offset+3], buf[offset+4], buf[offset+5])
	return true
}

func DateTime(offset int, buf []byte, size int, t *time.Time) bool {
	if offset+6 > size {
		return false
	}
	*t = time.Date(2000+int(buf[offset]), time.Month(buf[offset+1]), int(buf[offset+2]), int(buf[offset+3]), int(buf[offset+4]), int(buf[offset+5]), 0, time.Local)
	return true
}
