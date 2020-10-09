# Go Pointer Benchmark

When should you use pointers to data when calling functions? There are lots of conflicting views on this topic around the internet. The arguments _against_ pointers are:

- Pointers should only be used for mutable data
- Pointers are not necessarily faster than copies because passing pointers places extra burden on the garbage collector

The first point, about signalling mutability, is a weak argument because if you have reference types (arrays, maps, pointers) in the data, they can still be modified even if the parent container was not passed as a pointer. Consider the following example:

```go
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
```

The second point is more interesting because there must be an inflexion point after which the cost of copying a large struct outweighs the cost of GC and other considerations. This repository contains some code I wrote to try and figure this out. This is very much a work in progress and there are lots of subtle issues to work around (such as inlining) that skew the results.


## Benchmarks


### Struct containing all primitive fields

Attempts to measure the cost of passing a struct containing all primitive `uint64` (8 bytes) fields as the number of fields increase. 


```shell
# Go up to 15 fields, running each benchmark 10 times
make draw N=15 COUNT=10 BENCHMARK=Primitive
```

![graph](results/Primitive.svg)


After about 32 bytes, copying seems to slow things down quite a bit. Pointers, unsurprisingly, have constant performance.

TODO: How to measure GC pressure?


### Struct containing all reference fields

Attempts to measure the cost of passing a struct containing all `map[string]int` fields as the number of fields increase. 


```shell
# Go up to 15 fields, running each benchmark 10 times
make draw N=15 COUNT=10 BENCHMARK=Reference
```

![graph](results/Reference.svg)

This...seems suspcious. I must have screwed up somewhere.

TODO: Figure out what’s wrong
