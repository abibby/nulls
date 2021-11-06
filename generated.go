package nulls

import "fmt"

type Bool bool

func NewBool(v bool) *Bool {
	return (*Bool)(&v)
}

func (v *Bool) Value() bool {
	if v == nil {
		var zero bool
		return zero
	}
	return bool(*v)
}

func (v *Bool) Bool() bool {
	return v.Value()
}


func (v *Bool) String() string {
	if v == nil {
		return ""
	}
	return fmt.Sprint(v.Value())
}


func (v *Bool) IsNull() bool {
	return v == nil
}

func (v *Bool) Ok() (bool, bool) {
	return v.Value(), !v.IsNull()
}

type String string

func NewString(v string) *String {
	return (*String)(&v)
}

func (v *String) Value() string {
	if v == nil {
		var zero string
		return zero
	}
	return string(*v)
}

func (v *String) String() string {
	return v.Value()
}



func (v *String) IsNull() bool {
	return v == nil
}

func (v *String) Ok() (string, bool) {
	return v.Value(), !v.IsNull()
}

type Int int

func NewInt(v int) *Int {
	return (*Int)(&v)
}

func (v *Int) Value() int {
	if v == nil {
		var zero int
		return zero
	}
	return int(*v)
}

func (v *Int) Int() int {
	return v.Value()
}


func (v *Int) String() string {
	if v == nil {
		return ""
	}
	return fmt.Sprint(v.Value())
}


func (v *Int) IsNull() bool {
	return v == nil
}

func (v *Int) Ok() (int, bool) {
	return v.Value(), !v.IsNull()
}

type Int8 int8

func NewInt8(v int8) *Int8 {
	return (*Int8)(&v)
}

func (v *Int8) Value() int8 {
	if v == nil {
		var zero int8
		return zero
	}
	return int8(*v)
}

func (v *Int8) Int8() int8 {
	return v.Value()
}


func (v *Int8) String() string {
	if v == nil {
		return ""
	}
	return fmt.Sprint(v.Value())
}


func (v *Int8) IsNull() bool {
	return v == nil
}

func (v *Int8) Ok() (int8, bool) {
	return v.Value(), !v.IsNull()
}

type Int16 int16

func NewInt16(v int16) *Int16 {
	return (*Int16)(&v)
}

func (v *Int16) Value() int16 {
	if v == nil {
		var zero int16
		return zero
	}
	return int16(*v)
}

func (v *Int16) Int16() int16 {
	return v.Value()
}


func (v *Int16) String() string {
	if v == nil {
		return ""
	}
	return fmt.Sprint(v.Value())
}


func (v *Int16) IsNull() bool {
	return v == nil
}

func (v *Int16) Ok() (int16, bool) {
	return v.Value(), !v.IsNull()
}

type Int32 int32

func NewInt32(v int32) *Int32 {
	return (*Int32)(&v)
}

func (v *Int32) Value() int32 {
	if v == nil {
		var zero int32
		return zero
	}
	return int32(*v)
}

func (v *Int32) Int32() int32 {
	return v.Value()
}


func (v *Int32) String() string {
	if v == nil {
		return ""
	}
	return fmt.Sprint(v.Value())
}


func (v *Int32) IsNull() bool {
	return v == nil
}

func (v *Int32) Ok() (int32, bool) {
	return v.Value(), !v.IsNull()
}

type Int64 int64

func NewInt64(v int64) *Int64 {
	return (*Int64)(&v)
}

func (v *Int64) Value() int64 {
	if v == nil {
		var zero int64
		return zero
	}
	return int64(*v)
}

func (v *Int64) Int64() int64 {
	return v.Value()
}


func (v *Int64) String() string {
	if v == nil {
		return ""
	}
	return fmt.Sprint(v.Value())
}


func (v *Int64) IsNull() bool {
	return v == nil
}

func (v *Int64) Ok() (int64, bool) {
	return v.Value(), !v.IsNull()
}

type Uint uint

func NewUint(v uint) *Uint {
	return (*Uint)(&v)
}

func (v *Uint) Value() uint {
	if v == nil {
		var zero uint
		return zero
	}
	return uint(*v)
}

func (v *Uint) Uint() uint {
	return v.Value()
}


func (v *Uint) String() string {
	if v == nil {
		return ""
	}
	return fmt.Sprint(v.Value())
}


func (v *Uint) IsNull() bool {
	return v == nil
}

func (v *Uint) Ok() (uint, bool) {
	return v.Value(), !v.IsNull()
}

type Uint8 uint8

func NewUint8(v uint8) *Uint8 {
	return (*Uint8)(&v)
}

func (v *Uint8) Value() uint8 {
	if v == nil {
		var zero uint8
		return zero
	}
	return uint8(*v)
}

func (v *Uint8) Uint8() uint8 {
	return v.Value()
}


func (v *Uint8) String() string {
	if v == nil {
		return ""
	}
	return fmt.Sprint(v.Value())
}


func (v *Uint8) IsNull() bool {
	return v == nil
}

func (v *Uint8) Ok() (uint8, bool) {
	return v.Value(), !v.IsNull()
}

type Uint16 uint16

func NewUint16(v uint16) *Uint16 {
	return (*Uint16)(&v)
}

func (v *Uint16) Value() uint16 {
	if v == nil {
		var zero uint16
		return zero
	}
	return uint16(*v)
}

func (v *Uint16) Uint16() uint16 {
	return v.Value()
}


func (v *Uint16) String() string {
	if v == nil {
		return ""
	}
	return fmt.Sprint(v.Value())
}


func (v *Uint16) IsNull() bool {
	return v == nil
}

func (v *Uint16) Ok() (uint16, bool) {
	return v.Value(), !v.IsNull()
}

type Uint32 uint32

func NewUint32(v uint32) *Uint32 {
	return (*Uint32)(&v)
}

func (v *Uint32) Value() uint32 {
	if v == nil {
		var zero uint32
		return zero
	}
	return uint32(*v)
}

func (v *Uint32) Uint32() uint32 {
	return v.Value()
}


func (v *Uint32) String() string {
	if v == nil {
		return ""
	}
	return fmt.Sprint(v.Value())
}


func (v *Uint32) IsNull() bool {
	return v == nil
}

func (v *Uint32) Ok() (uint32, bool) {
	return v.Value(), !v.IsNull()
}

type Uint64 uint64

func NewUint64(v uint64) *Uint64 {
	return (*Uint64)(&v)
}

func (v *Uint64) Value() uint64 {
	if v == nil {
		var zero uint64
		return zero
	}
	return uint64(*v)
}

func (v *Uint64) Uint64() uint64 {
	return v.Value()
}


func (v *Uint64) String() string {
	if v == nil {
		return ""
	}
	return fmt.Sprint(v.Value())
}


func (v *Uint64) IsNull() bool {
	return v == nil
}

func (v *Uint64) Ok() (uint64, bool) {
	return v.Value(), !v.IsNull()
}

type Uintptr uintptr

func NewUintptr(v uintptr) *Uintptr {
	return (*Uintptr)(&v)
}

func (v *Uintptr) Value() uintptr {
	if v == nil {
		var zero uintptr
		return zero
	}
	return uintptr(*v)
}

func (v *Uintptr) Uintptr() uintptr {
	return v.Value()
}


func (v *Uintptr) String() string {
	if v == nil {
		return ""
	}
	return fmt.Sprint(v.Value())
}


func (v *Uintptr) IsNull() bool {
	return v == nil
}

func (v *Uintptr) Ok() (uintptr, bool) {
	return v.Value(), !v.IsNull()
}

type Byte byte

func NewByte(v byte) *Byte {
	return (*Byte)(&v)
}

func (v *Byte) Value() byte {
	if v == nil {
		var zero byte
		return zero
	}
	return byte(*v)
}

func (v *Byte) Byte() byte {
	return v.Value()
}


func (v *Byte) String() string {
	if v == nil {
		return ""
	}
	return fmt.Sprint(v.Value())
}


func (v *Byte) IsNull() bool {
	return v == nil
}

func (v *Byte) Ok() (byte, bool) {
	return v.Value(), !v.IsNull()
}

type Rune rune

func NewRune(v rune) *Rune {
	return (*Rune)(&v)
}

func (v *Rune) Value() rune {
	if v == nil {
		var zero rune
		return zero
	}
	return rune(*v)
}

func (v *Rune) Rune() rune {
	return v.Value()
}


func (v *Rune) String() string {
	if v == nil {
		return ""
	}
	return fmt.Sprint(v.Value())
}


func (v *Rune) IsNull() bool {
	return v == nil
}

func (v *Rune) Ok() (rune, bool) {
	return v.Value(), !v.IsNull()
}

type Float32 float32

func NewFloat32(v float32) *Float32 {
	return (*Float32)(&v)
}

func (v *Float32) Value() float32 {
	if v == nil {
		var zero float32
		return zero
	}
	return float32(*v)
}

func (v *Float32) Float32() float32 {
	return v.Value()
}


func (v *Float32) String() string {
	if v == nil {
		return ""
	}
	return fmt.Sprint(v.Value())
}


func (v *Float32) IsNull() bool {
	return v == nil
}

func (v *Float32) Ok() (float32, bool) {
	return v.Value(), !v.IsNull()
}

type Float64 float64

func NewFloat64(v float64) *Float64 {
	return (*Float64)(&v)
}

func (v *Float64) Value() float64 {
	if v == nil {
		var zero float64
		return zero
	}
	return float64(*v)
}

func (v *Float64) Float64() float64 {
	return v.Value()
}


func (v *Float64) String() string {
	if v == nil {
		return ""
	}
	return fmt.Sprint(v.Value())
}


func (v *Float64) IsNull() bool {
	return v == nil
}

func (v *Float64) Ok() (float64, bool) {
	return v.Value(), !v.IsNull()
}

type Complex64 complex64

func NewComplex64(v complex64) *Complex64 {
	return (*Complex64)(&v)
}

func (v *Complex64) Value() complex64 {
	if v == nil {
		var zero complex64
		return zero
	}
	return complex64(*v)
}

func (v *Complex64) Complex64() complex64 {
	return v.Value()
}


func (v *Complex64) String() string {
	if v == nil {
		return ""
	}
	return fmt.Sprint(v.Value())
}


func (v *Complex64) IsNull() bool {
	return v == nil
}

func (v *Complex64) Ok() (complex64, bool) {
	return v.Value(), !v.IsNull()
}

type Complex128 complex128

func NewComplex128(v complex128) *Complex128 {
	return (*Complex128)(&v)
}

func (v *Complex128) Value() complex128 {
	if v == nil {
		var zero complex128
		return zero
	}
	return complex128(*v)
}

func (v *Complex128) Complex128() complex128 {
	return v.Value()
}


func (v *Complex128) String() string {
	if v == nil {
		return ""
	}
	return fmt.Sprint(v.Value())
}


func (v *Complex128) IsNull() bool {
	return v == nil
}

func (v *Complex128) Ok() (complex128, bool) {
	return v.Value(), !v.IsNull()
}

