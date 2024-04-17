package encodingx

import (
	"bytes"
	"encoding/binary"
	"math"
	"reflect"
	"testing"
	"time"
)

func TestMarshal(t *testing.T) {
	testMarshal(t, big, s)
}

func TestMarshalStructInlne(t *testing.T) {
	testMarshal(t, bigA, sa)
}

func TestUnmarshal(t *testing.T) {
	var s1 Struct
	testUnmarshal(t, big, &s1, &s)
}

func TestUnmarshal2(t *testing.T) {
	var s StructA
	testUnmarshal(t, bigA, &s, &sa)
}

func checkResult(t *testing.T, dir string, err error, have, want any) {
	if err != nil {
		t.Errorf("%v: %v", dir, err)
		return
	}
	if !reflect.DeepEqual(have, want) {
		t.Errorf("%v:\n\thave %+v\n\twant %+v", dir, have, want)
	}
}

type Struct struct {
	Int8       int8
	Int16      int16
	Int32      int32
	Int64      int64
	Uint8      uint8
	Uint16     uint16
	Uint32     uint32
	Uint64     uint64
	Float32    float32
	Float64    float64
	Complex64  complex64
	Complex128 complex128
	Array      [4]uint8
	Bool       bool
	BoolArray  [4]bool
}

var s = Struct{
	0x01,
	0x0203,
	0x04050607,
	0x08090a0b0c0d0e0f,
	0x10,
	0x1112,
	0x13141516,
	0x1718191a1b1c1d1e,

	math.Float32frombits(0x1f202122),
	math.Float64frombits(0x232425262728292a),
	complex(
		math.Float32frombits(0x2b2c2d2e),
		math.Float32frombits(0x2f303132),
	),
	complex(
		math.Float64frombits(0x333435363738393a),
		math.Float64frombits(0x3b3c3d3e3f404142),
	),

	[4]uint8{0x43, 0x44, 0x45, 0x46},

	true,
	[4]bool{true, false, true, false},
}

var big = []byte{
	1,
	2, 3,
	4, 5, 6, 7,
	8, 9, 10, 11, 12, 13, 14, 15,
	16,
	17, 18,
	19, 20, 21, 22,
	23, 24, 25, 26, 27, 28, 29, 30,

	31, 32, 33, 34,
	35, 36, 37, 38, 39, 40, 41, 42,
	43, 44, 45, 46, 47, 48, 49, 50,
	51, 52, 53, 54, 55, 56, 57, 58, 59, 60, 61, 62, 63, 64, 65, 66,

	67, 68, 69, 70,

	1,
	1, 0, 1, 0,
}

type StructA struct {
	Struct
	Data []byte
}

var sa = StructA{
	Struct: s,
	Data:   []byte{0x43, 0x44, 0x45, 0x46, 0x43, 0x44, 0x45, 0x46},
}

var bigA = []byte{
	1,
	2, 3,
	4, 5, 6, 7,
	8, 9, 10, 11, 12, 13, 14, 15,
	16,
	17, 18,
	19, 20, 21, 22,
	23, 24, 25, 26, 27, 28, 29, 30,

	31, 32, 33, 34,
	35, 36, 37, 38, 39, 40, 41, 42,
	43, 44, 45, 46, 47, 48, 49, 50,
	51, 52, 53, 54, 55, 56, 57, 58, 59, 60, 61, 62, 63, 64, 65, 66,

	67, 68, 69, 70,

	1,
	1, 0, 1, 0,

	67, 68, 69, 70, 67, 68, 69, 70,
}

func testMarshal(t *testing.T, b []byte, s1 any) {
	bs, err := Marshal(s1)
	checkResult(t, "Unmarshal", err, bs, b)
}

func testUnmarshal(t *testing.T, b []byte, hav, want any) {
	err := Unmarshal(b, hav)
	checkResult(t, "Unmarshal", err, hav, want)
}

func TestDateTimeStr(t *testing.T) {
	buf := []byte{0, 2, 3, 23, 3, 4, 12, 33, 3}

	var str string
	if !DateTimeStr(3, buf, len(buf), &str) {
		t.Fail()
	}
	if str != "2023-03-04 12:33:03" {
		t.Fail()
	}
}

func TestDateTime(t *testing.T) {

	n := time.Date(2023, 3, 3, 19, 12, 3, 0, time.Local)

	buf := []byte{0, 1, 2, 23, 3, 3, 19, 12, 3}

	var st time.Time
	if !DateTime(3, buf, len(buf), &st) {
		t.Fail()
	}
	if !st.Equal(n) {
		t.Fail()
	}
}

type Tcase[T comparable] struct {
	v      T
	target T
	offset int
}

func (tc *Tcase[T]) Test(t *testing.T, offset int, wb *bytes.Buffer) {

	binary.Write(wb, byteOrder, tc.v)

	if !Value(offset, wb.Bytes(), wb.Len(), &tc.target) {
		t.Fatal(tc)
	}
	if tc.target != tc.v {
		t.Fail()
	}

}

func TestValue(t *testing.T) {

	wb := bytes.NewBuffer([]byte{0, 0})

	t1 := Tcase[int8]{-123, 0, 1}
	t2 := Tcase[int16]{-123, 0, 2}
	t3 := Tcase[int32]{-123, 0, 4}
	t4 := Tcase[int64]{-123, 0, 8}
	t5 := Tcase[float32]{-12223.12333333, 0, 4}
	t6 := Tcase[float64]{-12223.12333333, 0, 8}

	offset := 2

	t1.Test(t, offset, wb)
	offset += t1.offset

	t2.Test(t, offset, wb)
	offset += t2.offset

	t3.Test(t, offset, wb)
	offset += t3.offset
	t4.Test(t, offset, wb)
	offset += t4.offset
	t5.Test(t, offset, wb)
	offset += t5.offset
	t6.Test(t, offset, wb)
	offset += t6.offset

}

func TestInt8ToFloat(t *testing.T) {
	var i int8 = -99
	buf := []byte{byte(i)}
	var v float64
	if !Int8ToFloat(0, buf, 1, 0.01, &v) {
		t.Fail()
	}
	if v != -0.99 {
		t.Fail()
	}

}

func TestInt16ToInt(t *testing.T) {
	var i1 int16 = -230
	var buf []byte = make([]byte, 2)
	binary.BigEndian.PutUint16(buf, uint16(i1))
	var r int
	if !Int16ToInt(0, buf, len(buf), &r) {
		t.Fail()
	}
	if r != -230 {
		t.Fail()
	}
}
