package nulls

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	n := New(10)
	assert.Equal(t, 10, n.Value())
	assert.True(t, n.valid)
	v, ok := n.Ok()
	assert.Equal(t, 10, v)
	assert.True(t, ok)
}

func TestNull(t *testing.T) {
	n := Null[int]()
	assert.Equal(t, 0, n.Value()) // zero value of int
	assert.False(t, n.valid)
	v, ok := n.Ok()
	assert.Equal(t, 0, v) // zero value of int
	assert.False(t, ok)
}

func TestString(t *testing.T) {
	n := New(10)
	assert.Equal(t, "10", n.String())

	nullN := Null[int]()
	assert.Equal(t, "", nullN.String())
}

func TestIsNull(t *testing.T) {
	n := New(10)
	assert.False(t, n.IsNull())

	nullN := Null[int]()
	assert.True(t, nullN.IsNull())
}

func TestMarshalJSON(t *testing.T) {
	n := New(10)
	b, err := json.Marshal(n)
	assert.NoError(t, err)
	assert.Equal(t, "10", string(b))

	nullN := Null[int]()
	b, err = json.Marshal(nullN)
	assert.NoError(t, err)
	assert.Equal(t, "null", string(b))

	type TestStruct struct {
		Value Nullable[int] `json:"value"`
	}

	testStruct := TestStruct{Value: New(5)}
	b, err = json.Marshal(testStruct)
	assert.NoError(t, err)
	assert.Equal(t, `{"value":5}`, string(b))

	testStructNull := TestStruct{Value: Null[int]()}
	b, err = json.Marshal(testStructNull)
	assert.NoError(t, err)
	assert.Equal(t, `{"value":null}`, string(b))

	testString := New("test")
	b, err = json.Marshal(testString)
	assert.NoError(t, err)
	assert.Equal(t, `"test"`, string(b))

	testStringNull := Null[string]()
	b, err = json.Marshal(testStringNull)
	assert.NoError(t, err)
	assert.Equal(t, `null`, string(b))
}

func TestUnmarshalJSON(t *testing.T) {
	var n Nullable[int]
	err := json.Unmarshal([]byte("10"), &n)
	assert.NoError(t, err)
	assert.Equal(t, 10, n.Value())
	assert.True(t, n.valid)

	err = json.Unmarshal([]byte("null"), &n)
	assert.NoError(t, err)
	assert.False(t, n.valid)

	type TestStruct struct {
		Value Nullable[int] `json:"value"`
	}

	var testStruct TestStruct
	err = json.Unmarshal([]byte(`{"value":5}`), &testStruct)
	assert.NoError(t, err)
	assert.Equal(t, 5, testStruct.Value.Value())
	assert.True(t, testStruct.Value.valid)

	var testStructNull TestStruct
	err = json.Unmarshal([]byte(`{"value":null}`), &testStructNull)
	assert.NoError(t, err)
	assert.False(t, testStructNull.Value.valid)

	var stringNullable Nullable[string]
	err = json.Unmarshal([]byte(`"test"`), &stringNullable)
	assert.NoError(t, err)
	assert.Equal(t, "test", stringNullable.Value())
	assert.True(t, stringNullable.valid)

	var stringNullableNull Nullable[string]
	err = json.Unmarshal([]byte(`null`), &stringNullableNull)
	assert.NoError(t, err)
	assert.False(t, stringNullableNull.valid)
}
