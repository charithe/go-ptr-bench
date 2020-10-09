package mutability

import (
	"testing"

	"github.com/stretchr/testify/require"
)

type MyStruct struct {
	intField int
	mapField map[string]int
}

func passValue(m MyStruct) {
	m.intField++
	m.mapField["key"] = 10
}

func passPointer(m *MyStruct) {
	m.intField++
	m.mapField["key"] = 10
}

func TestMutability(t *testing.T) {
	t.Run("method=by_value", func(t *testing.T) {
		m := MyStruct{intField: 1, mapField: map[string]int{"key": 1}}
		passValue(m)
		t.Logf("Value: %+v", m)

		// primitive field does not get modified
		require.Equal(t, 1, m.intField)
		// reference field does get modified
		require.Equal(t, 10, m.mapField["key"])
	})

	t.Run("method=by_pointer", func(t *testing.T) {
		m := MyStruct{intField: 1, mapField: map[string]int{"key": 1}}
		passPointer(&m)
		t.Logf("Value: %+v", m)

		// both fields get modified
		require.Equal(t, 2, m.intField)
		require.Equal(t, 10, m.mapField["key"])
	})
}
