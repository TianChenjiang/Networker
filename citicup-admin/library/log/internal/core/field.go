package core

import (
	"fmt"
	"math"
	"time"

	xtime "go-web-demo/library/time"
)

// FieldType represent D value type
type FieldType int32

// DType enum
const (
	UnknownType FieldType = iota
	StringType
	IntType
	Int64Type
	UintType
	Uint64Type
	Float32Type
	Float64Type
	DurationType
)

// Field is for encoder
type Field struct {
	Key       string
	Value     interface{}
	Type      FieldType
	StringVal string
	Int64Val  int64
}

// AddTo exports a field through the ObjectEncoder interface. It's primarily
// useful to library authors, and shouldn't be necessary in most applications.
func (f Field) AddTo(enc ObjectEncoder) {
	if f.Type == UnknownType {
		f.assertAddTo(enc)
		return
	}
	switch f.Type {
	case StringType:
		AddString(f.Key, f.StringVal)
	case IntType:
		AddInt(f.Key, int(f.Int64Val))
	case Int64Type:
		AddInt64(f.Key, f.Int64Val)
	case UintType:
		AddUint(f.Key, uint(f.Int64Val))
	case Uint64Type:
		AddUint64(f.Key, uint64(f.Int64Val))
	case Float32Type:
		AddFloat32(f.Key, math.Float32frombits(uint32(f.Int64Val)))
	case Float64Type:
		AddFloat64(f.Key, math.Float64frombits(uint64(f.Int64Val)))
	case DurationType:
		AddDuration(f.Key, time.Duration(f.Int64Val))
	default:
		panic(fmt.Sprintf("unknown field type: %v", f))
	}
}

func (f Field) assertAddTo(enc ObjectEncoder) {
	// assert interface
	switch val := f.Value.(type) {
	case bool:
		AddBool(f.Key, val)
	case complex128:
		AddComplex128(f.Key, val)
	case complex64:
		AddComplex64(f.Key, val)
	case float64:
		AddFloat64(f.Key, val)
	case float32:
		AddFloat32(f.Key, val)
	case int:
		AddInt(f.Key, val)
	case int64:
		AddInt64(f.Key, val)
	case int32:
		AddInt32(f.Key, val)
	case int16:
		AddInt16(f.Key, val)
	case int8:
		AddInt8(f.Key, val)
	case string:
		AddString(f.Key, val)
	case uint:
		AddUint(f.Key, val)
	case uint64:
		AddUint64(f.Key, val)
	case uint32:
		AddUint32(f.Key, val)
	case uint16:
		AddUint16(f.Key, val)
	case uint8:
		AddUint8(f.Key, val)
	case []byte:
		AddByteString(f.Key, val)
	case uintptr:
		AddUintptr(f.Key, val)
	case time.Time:
		AddTime(f.Key, val)
	case xtime.Time:
		AddTime(f.Key, val.Time())
	case time.Duration:
		AddDuration(f.Key, val)
	case xtime.Duration:
		AddDuration(f.Key, time.Duration(val))
	case error:
		AddString(f.Key, val.Error())
	case fmt.Stringer:
		AddString(f.Key, val.String())
	default:
		err := AddReflected(f.Key, val)
		if err != nil {
			AddString(fmt.Sprintf("%sError", f.Key), err.Error())
		}
	}
}
