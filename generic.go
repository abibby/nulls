package nulls

import (
	"bytes"
	"encoding/json"
	"fmt"
)

var nullBytes = []byte("null")

type Nullable[T any] struct {
	value T
	valid bool
}

var _ json.Marshaler = Nullable[int]{}
var _ json.Unmarshaler = (*Nullable[int])(nil)

func New[T any](n T) Nullable[T] {
	return Nullable[T]{
		value: n,
		valid: true,
	}
}

func Null[T any]() Nullable[T] {
	return Nullable[T]{
		valid: false,
	}
}

func (n Nullable[T]) Value() T {
	return n.value
}

func (n Nullable[T]) String() string {
	if !n.valid {
		return ""
	}
	return fmt.Sprint(n.value)
}

func (n Nullable[T]) IsNull() bool {
	return !n.valid
}

func (n Nullable[T]) Ok() (T, bool) {
	return n.value, n.valid
}

// MarshalJSON implements json.Marshaler.
func (n Nullable[T]) MarshalJSON() ([]byte, error) {
	if !n.valid {
		return nullBytes, nil
	}
	return json.Marshal(n.value)
}

// UnmarshalJSON implements json.Unmarshaler.
func (n *Nullable[T]) UnmarshalJSON(b []byte) error {
	if bytes.Equal(b, nullBytes) {
		var zero T
		n.value = zero
		n.valid = false
		return nil
	}

	n.valid = true
	return json.Unmarshal(b, &n.value)
}
