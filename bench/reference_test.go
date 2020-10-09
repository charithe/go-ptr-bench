package bench

import (
	"fmt"
	"testing"
	"unsafe"
)

const mapSize = 10

var refSum = 0

type R1 struct {
	f0 map[string]int
}

//go:noinline
func referencePassByValue1(r R1) int {
	return r.f0["key"]
}

//go:noinline
func referencePassByPtr1(r *R1) int {
	return r.f0["key"]
}

type R2 struct {
	f0 map[string]int
	f1 map[string]int
}

//go:noinline
func referencePassByValue2(r R2) int {
	return r.f0["key"]
}

//go:noinline
func referencePassByPtr2(r *R2) int {
	return r.f0["key"]
}

type R3 struct {
	f0 map[string]int
	f1 map[string]int
	f2 map[string]int
}

//go:noinline
func referencePassByValue3(r R3) int {
	return r.f0["key"]
}

//go:noinline
func referencePassByPtr3(r *R3) int {
	return r.f0["key"]
}

type R4 struct {
	f0 map[string]int
	f1 map[string]int
	f2 map[string]int
	f3 map[string]int
}

//go:noinline
func referencePassByValue4(r R4) int {
	return r.f0["key"]
}

//go:noinline
func referencePassByPtr4(r *R4) int {
	return r.f0["key"]
}

type R5 struct {
	f0 map[string]int
	f1 map[string]int
	f2 map[string]int
	f3 map[string]int
	f4 map[string]int
}

//go:noinline
func referencePassByValue5(r R5) int {
	return r.f0["key"]
}

//go:noinline
func referencePassByPtr5(r *R5) int {
	return r.f0["key"]
}

type R6 struct {
	f0 map[string]int
	f1 map[string]int
	f2 map[string]int
	f3 map[string]int
	f4 map[string]int
	f5 map[string]int
}

//go:noinline
func referencePassByValue6(r R6) int {
	return r.f0["key"]
}

//go:noinline
func referencePassByPtr6(r *R6) int {
	return r.f0["key"]
}

type R7 struct {
	f0 map[string]int
	f1 map[string]int
	f2 map[string]int
	f3 map[string]int
	f4 map[string]int
	f5 map[string]int
	f6 map[string]int
}

//go:noinline
func referencePassByValue7(r R7) int {
	return r.f0["key"]
}

//go:noinline
func referencePassByPtr7(r *R7) int {
	return r.f0["key"]
}

type R8 struct {
	f0 map[string]int
	f1 map[string]int
	f2 map[string]int
	f3 map[string]int
	f4 map[string]int
	f5 map[string]int
	f6 map[string]int
	f7 map[string]int
}

//go:noinline
func referencePassByValue8(r R8) int {
	return r.f0["key"]
}

//go:noinline
func referencePassByPtr8(r *R8) int {
	return r.f0["key"]
}

type R9 struct {
	f0 map[string]int
	f1 map[string]int
	f2 map[string]int
	f3 map[string]int
	f4 map[string]int
	f5 map[string]int
	f6 map[string]int
	f7 map[string]int
	f8 map[string]int
}

//go:noinline
func referencePassByValue9(r R9) int {
	return r.f0["key"]
}

//go:noinline
func referencePassByPtr9(r *R9) int {
	return r.f0["key"]
}

type R10 struct {
	f0 map[string]int
	f1 map[string]int
	f2 map[string]int
	f3 map[string]int
	f4 map[string]int
	f5 map[string]int
	f6 map[string]int
	f7 map[string]int
	f8 map[string]int
	f9 map[string]int
}

//go:noinline
func referencePassByValue10(r R10) int {
	return r.f0["key"]
}

//go:noinline
func referencePassByPtr10(r *R10) int {
	return r.f0["key"]
}

type R11 struct {
	f0  map[string]int
	f1  map[string]int
	f2  map[string]int
	f3  map[string]int
	f4  map[string]int
	f5  map[string]int
	f6  map[string]int
	f7  map[string]int
	f8  map[string]int
	f9  map[string]int
	f10 map[string]int
}

//go:noinline
func referencePassByValue11(r R11) int {
	return r.f0["key"]
}

//go:noinline
func referencePassByPtr11(r *R11) int {
	return r.f0["key"]
}

type R12 struct {
	f0  map[string]int
	f1  map[string]int
	f2  map[string]int
	f3  map[string]int
	f4  map[string]int
	f5  map[string]int
	f6  map[string]int
	f7  map[string]int
	f8  map[string]int
	f9  map[string]int
	f10 map[string]int
	f11 map[string]int
}

//go:noinline
func referencePassByValue12(r R12) int {
	return r.f0["key"]
}

//go:noinline
func referencePassByPtr12(r *R12) int {
	return r.f0["key"]
}

type R13 struct {
	f0  map[string]int
	f1  map[string]int
	f2  map[string]int
	f3  map[string]int
	f4  map[string]int
	f5  map[string]int
	f6  map[string]int
	f7  map[string]int
	f8  map[string]int
	f9  map[string]int
	f10 map[string]int
	f11 map[string]int
	f12 map[string]int
}

//go:noinline
func referencePassByValue13(r R13) int {
	return r.f0["key"]
}

//go:noinline
func referencePassByPtr13(r *R13) int {
	return r.f0["key"]
}

type R14 struct {
	f0  map[string]int
	f1  map[string]int
	f2  map[string]int
	f3  map[string]int
	f4  map[string]int
	f5  map[string]int
	f6  map[string]int
	f7  map[string]int
	f8  map[string]int
	f9  map[string]int
	f10 map[string]int
	f11 map[string]int
	f12 map[string]int
	f13 map[string]int
}

//go:noinline
func referencePassByValue14(r R14) int {
	return r.f0["key"]
}

//go:noinline
func referencePassByPtr14(r *R14) int {
	return r.f0["key"]
}

type R15 struct {
	f0  map[string]int
	f1  map[string]int
	f2  map[string]int
	f3  map[string]int
	f4  map[string]int
	f5  map[string]int
	f6  map[string]int
	f7  map[string]int
	f8  map[string]int
	f9  map[string]int
	f10 map[string]int
	f11 map[string]int
	f12 map[string]int
	f13 map[string]int
	f14 map[string]int
}

//go:noinline
func referencePassByValue15(r R15) int {
	return r.f0["key"]
}

//go:noinline
func referencePassByPtr15(r *R15) int {
	return r.f0["key"]
}

func BenchmarkReferenceFields(b *testing.B) {
	b.ReportAllocs()
	b.Run("method=value/fields=1", func(b *testing.B) {
		r := R1{
			f0: make(map[string]int, mapSize),
		}

		r.f0["x"] = 10

		size := unsafe.Sizeof(r)
		b.Run(fmt.Sprintf("size=%d", size), func(b *testing.B) {
			refSum = referencePassByValue1(r)
		})
	})

	b.Run("method=ptr/fields=1", func(b *testing.B) {
		r := R1{
			f0: make(map[string]int, mapSize),
		}

		r.f0["x"] = 10

		size := unsafe.Sizeof(r)
		b.Run(fmt.Sprintf("size=%d", size), func(b *testing.B) {
			refSum = referencePassByPtr1(&r)
		})
	})
	b.Run("method=value/fields=2", func(b *testing.B) {
		r := R2{
			f0: make(map[string]int, mapSize),
			f1: make(map[string]int, mapSize),
		}

		r.f0["x"] = 10

		size := unsafe.Sizeof(r)
		b.Run(fmt.Sprintf("size=%d", size), func(b *testing.B) {
			refSum = referencePassByValue2(r)
		})
	})

	b.Run("method=ptr/fields=1", func(b *testing.B) {
		r := R2{
			f0: make(map[string]int, mapSize),
			f1: make(map[string]int, mapSize),
		}

		r.f0["x"] = 10

		size := unsafe.Sizeof(r)
		b.Run(fmt.Sprintf("size=%d", size), func(b *testing.B) {
			refSum = referencePassByPtr2(&r)
		})
	})
	b.Run("method=value/fields=3", func(b *testing.B) {
		r := R3{
			f0: make(map[string]int, mapSize),
			f1: make(map[string]int, mapSize),
			f2: make(map[string]int, mapSize),
		}

		r.f0["x"] = 10

		size := unsafe.Sizeof(r)
		b.Run(fmt.Sprintf("size=%d", size), func(b *testing.B) {
			refSum = referencePassByValue3(r)
		})
	})

	b.Run("method=ptr/fields=1", func(b *testing.B) {
		r := R3{
			f0: make(map[string]int, mapSize),
			f1: make(map[string]int, mapSize),
			f2: make(map[string]int, mapSize),
		}

		r.f0["x"] = 10

		size := unsafe.Sizeof(r)
		b.Run(fmt.Sprintf("size=%d", size), func(b *testing.B) {
			refSum = referencePassByPtr3(&r)
		})
	})
	b.Run("method=value/fields=4", func(b *testing.B) {
		r := R4{
			f0: make(map[string]int, mapSize),
			f1: make(map[string]int, mapSize),
			f2: make(map[string]int, mapSize),
			f3: make(map[string]int, mapSize),
		}

		r.f0["x"] = 10

		size := unsafe.Sizeof(r)
		b.Run(fmt.Sprintf("size=%d", size), func(b *testing.B) {
			refSum = referencePassByValue4(r)
		})
	})

	b.Run("method=ptr/fields=1", func(b *testing.B) {
		r := R4{
			f0: make(map[string]int, mapSize),
			f1: make(map[string]int, mapSize),
			f2: make(map[string]int, mapSize),
			f3: make(map[string]int, mapSize),
		}

		r.f0["x"] = 10

		size := unsafe.Sizeof(r)
		b.Run(fmt.Sprintf("size=%d", size), func(b *testing.B) {
			refSum = referencePassByPtr4(&r)
		})
	})
	b.Run("method=value/fields=5", func(b *testing.B) {
		r := R5{
			f0: make(map[string]int, mapSize),
			f1: make(map[string]int, mapSize),
			f2: make(map[string]int, mapSize),
			f3: make(map[string]int, mapSize),
			f4: make(map[string]int, mapSize),
		}

		r.f0["x"] = 10

		size := unsafe.Sizeof(r)
		b.Run(fmt.Sprintf("size=%d", size), func(b *testing.B) {
			refSum = referencePassByValue5(r)
		})
	})

	b.Run("method=ptr/fields=1", func(b *testing.B) {
		r := R5{
			f0: make(map[string]int, mapSize),
			f1: make(map[string]int, mapSize),
			f2: make(map[string]int, mapSize),
			f3: make(map[string]int, mapSize),
			f4: make(map[string]int, mapSize),
		}

		r.f0["x"] = 10

		size := unsafe.Sizeof(r)
		b.Run(fmt.Sprintf("size=%d", size), func(b *testing.B) {
			refSum = referencePassByPtr5(&r)
		})
	})
	b.Run("method=value/fields=6", func(b *testing.B) {
		r := R6{
			f0: make(map[string]int, mapSize),
			f1: make(map[string]int, mapSize),
			f2: make(map[string]int, mapSize),
			f3: make(map[string]int, mapSize),
			f4: make(map[string]int, mapSize),
			f5: make(map[string]int, mapSize),
		}

		r.f0["x"] = 10

		size := unsafe.Sizeof(r)
		b.Run(fmt.Sprintf("size=%d", size), func(b *testing.B) {
			refSum = referencePassByValue6(r)
		})
	})

	b.Run("method=ptr/fields=1", func(b *testing.B) {
		r := R6{
			f0: make(map[string]int, mapSize),
			f1: make(map[string]int, mapSize),
			f2: make(map[string]int, mapSize),
			f3: make(map[string]int, mapSize),
			f4: make(map[string]int, mapSize),
			f5: make(map[string]int, mapSize),
		}

		r.f0["x"] = 10

		size := unsafe.Sizeof(r)
		b.Run(fmt.Sprintf("size=%d", size), func(b *testing.B) {
			refSum = referencePassByPtr6(&r)
		})
	})
	b.Run("method=value/fields=7", func(b *testing.B) {
		r := R7{
			f0: make(map[string]int, mapSize),
			f1: make(map[string]int, mapSize),
			f2: make(map[string]int, mapSize),
			f3: make(map[string]int, mapSize),
			f4: make(map[string]int, mapSize),
			f5: make(map[string]int, mapSize),
			f6: make(map[string]int, mapSize),
		}

		r.f0["x"] = 10

		size := unsafe.Sizeof(r)
		b.Run(fmt.Sprintf("size=%d", size), func(b *testing.B) {
			refSum = referencePassByValue7(r)
		})
	})

	b.Run("method=ptr/fields=1", func(b *testing.B) {
		r := R7{
			f0: make(map[string]int, mapSize),
			f1: make(map[string]int, mapSize),
			f2: make(map[string]int, mapSize),
			f3: make(map[string]int, mapSize),
			f4: make(map[string]int, mapSize),
			f5: make(map[string]int, mapSize),
			f6: make(map[string]int, mapSize),
		}

		r.f0["x"] = 10

		size := unsafe.Sizeof(r)
		b.Run(fmt.Sprintf("size=%d", size), func(b *testing.B) {
			refSum = referencePassByPtr7(&r)
		})
	})
	b.Run("method=value/fields=8", func(b *testing.B) {
		r := R8{
			f0: make(map[string]int, mapSize),
			f1: make(map[string]int, mapSize),
			f2: make(map[string]int, mapSize),
			f3: make(map[string]int, mapSize),
			f4: make(map[string]int, mapSize),
			f5: make(map[string]int, mapSize),
			f6: make(map[string]int, mapSize),
			f7: make(map[string]int, mapSize),
		}

		r.f0["x"] = 10

		size := unsafe.Sizeof(r)
		b.Run(fmt.Sprintf("size=%d", size), func(b *testing.B) {
			refSum = referencePassByValue8(r)
		})
	})

	b.Run("method=ptr/fields=1", func(b *testing.B) {
		r := R8{
			f0: make(map[string]int, mapSize),
			f1: make(map[string]int, mapSize),
			f2: make(map[string]int, mapSize),
			f3: make(map[string]int, mapSize),
			f4: make(map[string]int, mapSize),
			f5: make(map[string]int, mapSize),
			f6: make(map[string]int, mapSize),
			f7: make(map[string]int, mapSize),
		}

		r.f0["x"] = 10

		size := unsafe.Sizeof(r)
		b.Run(fmt.Sprintf("size=%d", size), func(b *testing.B) {
			refSum = referencePassByPtr8(&r)
		})
	})
	b.Run("method=value/fields=9", func(b *testing.B) {
		r := R9{
			f0: make(map[string]int, mapSize),
			f1: make(map[string]int, mapSize),
			f2: make(map[string]int, mapSize),
			f3: make(map[string]int, mapSize),
			f4: make(map[string]int, mapSize),
			f5: make(map[string]int, mapSize),
			f6: make(map[string]int, mapSize),
			f7: make(map[string]int, mapSize),
			f8: make(map[string]int, mapSize),
		}

		r.f0["x"] = 10

		size := unsafe.Sizeof(r)
		b.Run(fmt.Sprintf("size=%d", size), func(b *testing.B) {
			refSum = referencePassByValue9(r)
		})
	})

	b.Run("method=ptr/fields=1", func(b *testing.B) {
		r := R9{
			f0: make(map[string]int, mapSize),
			f1: make(map[string]int, mapSize),
			f2: make(map[string]int, mapSize),
			f3: make(map[string]int, mapSize),
			f4: make(map[string]int, mapSize),
			f5: make(map[string]int, mapSize),
			f6: make(map[string]int, mapSize),
			f7: make(map[string]int, mapSize),
			f8: make(map[string]int, mapSize),
		}

		r.f0["x"] = 10

		size := unsafe.Sizeof(r)
		b.Run(fmt.Sprintf("size=%d", size), func(b *testing.B) {
			refSum = referencePassByPtr9(&r)
		})
	})
	b.Run("method=value/fields=10", func(b *testing.B) {
		r := R10{
			f0: make(map[string]int, mapSize),
			f1: make(map[string]int, mapSize),
			f2: make(map[string]int, mapSize),
			f3: make(map[string]int, mapSize),
			f4: make(map[string]int, mapSize),
			f5: make(map[string]int, mapSize),
			f6: make(map[string]int, mapSize),
			f7: make(map[string]int, mapSize),
			f8: make(map[string]int, mapSize),
			f9: make(map[string]int, mapSize),
		}

		r.f0["x"] = 10

		size := unsafe.Sizeof(r)
		b.Run(fmt.Sprintf("size=%d", size), func(b *testing.B) {
			refSum = referencePassByValue10(r)
		})
	})

	b.Run("method=ptr/fields=1", func(b *testing.B) {
		r := R10{
			f0: make(map[string]int, mapSize),
			f1: make(map[string]int, mapSize),
			f2: make(map[string]int, mapSize),
			f3: make(map[string]int, mapSize),
			f4: make(map[string]int, mapSize),
			f5: make(map[string]int, mapSize),
			f6: make(map[string]int, mapSize),
			f7: make(map[string]int, mapSize),
			f8: make(map[string]int, mapSize),
			f9: make(map[string]int, mapSize),
		}

		r.f0["x"] = 10

		size := unsafe.Sizeof(r)
		b.Run(fmt.Sprintf("size=%d", size), func(b *testing.B) {
			refSum = referencePassByPtr10(&r)
		})
	})
	b.Run("method=value/fields=11", func(b *testing.B) {
		r := R11{
			f0:  make(map[string]int, mapSize),
			f1:  make(map[string]int, mapSize),
			f2:  make(map[string]int, mapSize),
			f3:  make(map[string]int, mapSize),
			f4:  make(map[string]int, mapSize),
			f5:  make(map[string]int, mapSize),
			f6:  make(map[string]int, mapSize),
			f7:  make(map[string]int, mapSize),
			f8:  make(map[string]int, mapSize),
			f9:  make(map[string]int, mapSize),
			f10: make(map[string]int, mapSize),
		}

		r.f0["x"] = 10

		size := unsafe.Sizeof(r)
		b.Run(fmt.Sprintf("size=%d", size), func(b *testing.B) {
			refSum = referencePassByValue11(r)
		})
	})

	b.Run("method=ptr/fields=1", func(b *testing.B) {
		r := R11{
			f0:  make(map[string]int, mapSize),
			f1:  make(map[string]int, mapSize),
			f2:  make(map[string]int, mapSize),
			f3:  make(map[string]int, mapSize),
			f4:  make(map[string]int, mapSize),
			f5:  make(map[string]int, mapSize),
			f6:  make(map[string]int, mapSize),
			f7:  make(map[string]int, mapSize),
			f8:  make(map[string]int, mapSize),
			f9:  make(map[string]int, mapSize),
			f10: make(map[string]int, mapSize),
		}

		r.f0["x"] = 10

		size := unsafe.Sizeof(r)
		b.Run(fmt.Sprintf("size=%d", size), func(b *testing.B) {
			refSum = referencePassByPtr11(&r)
		})
	})
	b.Run("method=value/fields=12", func(b *testing.B) {
		r := R12{
			f0:  make(map[string]int, mapSize),
			f1:  make(map[string]int, mapSize),
			f2:  make(map[string]int, mapSize),
			f3:  make(map[string]int, mapSize),
			f4:  make(map[string]int, mapSize),
			f5:  make(map[string]int, mapSize),
			f6:  make(map[string]int, mapSize),
			f7:  make(map[string]int, mapSize),
			f8:  make(map[string]int, mapSize),
			f9:  make(map[string]int, mapSize),
			f10: make(map[string]int, mapSize),
			f11: make(map[string]int, mapSize),
		}

		r.f0["x"] = 10

		size := unsafe.Sizeof(r)
		b.Run(fmt.Sprintf("size=%d", size), func(b *testing.B) {
			refSum = referencePassByValue12(r)
		})
	})

	b.Run("method=ptr/fields=1", func(b *testing.B) {
		r := R12{
			f0:  make(map[string]int, mapSize),
			f1:  make(map[string]int, mapSize),
			f2:  make(map[string]int, mapSize),
			f3:  make(map[string]int, mapSize),
			f4:  make(map[string]int, mapSize),
			f5:  make(map[string]int, mapSize),
			f6:  make(map[string]int, mapSize),
			f7:  make(map[string]int, mapSize),
			f8:  make(map[string]int, mapSize),
			f9:  make(map[string]int, mapSize),
			f10: make(map[string]int, mapSize),
			f11: make(map[string]int, mapSize),
		}

		r.f0["x"] = 10

		size := unsafe.Sizeof(r)
		b.Run(fmt.Sprintf("size=%d", size), func(b *testing.B) {
			refSum = referencePassByPtr12(&r)
		})
	})
	b.Run("method=value/fields=13", func(b *testing.B) {
		r := R13{
			f0:  make(map[string]int, mapSize),
			f1:  make(map[string]int, mapSize),
			f2:  make(map[string]int, mapSize),
			f3:  make(map[string]int, mapSize),
			f4:  make(map[string]int, mapSize),
			f5:  make(map[string]int, mapSize),
			f6:  make(map[string]int, mapSize),
			f7:  make(map[string]int, mapSize),
			f8:  make(map[string]int, mapSize),
			f9:  make(map[string]int, mapSize),
			f10: make(map[string]int, mapSize),
			f11: make(map[string]int, mapSize),
			f12: make(map[string]int, mapSize),
		}

		r.f0["x"] = 10

		size := unsafe.Sizeof(r)
		b.Run(fmt.Sprintf("size=%d", size), func(b *testing.B) {
			refSum = referencePassByValue13(r)
		})
	})

	b.Run("method=ptr/fields=1", func(b *testing.B) {
		r := R13{
			f0:  make(map[string]int, mapSize),
			f1:  make(map[string]int, mapSize),
			f2:  make(map[string]int, mapSize),
			f3:  make(map[string]int, mapSize),
			f4:  make(map[string]int, mapSize),
			f5:  make(map[string]int, mapSize),
			f6:  make(map[string]int, mapSize),
			f7:  make(map[string]int, mapSize),
			f8:  make(map[string]int, mapSize),
			f9:  make(map[string]int, mapSize),
			f10: make(map[string]int, mapSize),
			f11: make(map[string]int, mapSize),
			f12: make(map[string]int, mapSize),
		}

		r.f0["x"] = 10

		size := unsafe.Sizeof(r)
		b.Run(fmt.Sprintf("size=%d", size), func(b *testing.B) {
			refSum = referencePassByPtr13(&r)
		})
	})
	b.Run("method=value/fields=14", func(b *testing.B) {
		r := R14{
			f0:  make(map[string]int, mapSize),
			f1:  make(map[string]int, mapSize),
			f2:  make(map[string]int, mapSize),
			f3:  make(map[string]int, mapSize),
			f4:  make(map[string]int, mapSize),
			f5:  make(map[string]int, mapSize),
			f6:  make(map[string]int, mapSize),
			f7:  make(map[string]int, mapSize),
			f8:  make(map[string]int, mapSize),
			f9:  make(map[string]int, mapSize),
			f10: make(map[string]int, mapSize),
			f11: make(map[string]int, mapSize),
			f12: make(map[string]int, mapSize),
			f13: make(map[string]int, mapSize),
		}

		r.f0["x"] = 10

		size := unsafe.Sizeof(r)
		b.Run(fmt.Sprintf("size=%d", size), func(b *testing.B) {
			refSum = referencePassByValue14(r)
		})
	})

	b.Run("method=ptr/fields=1", func(b *testing.B) {
		r := R14{
			f0:  make(map[string]int, mapSize),
			f1:  make(map[string]int, mapSize),
			f2:  make(map[string]int, mapSize),
			f3:  make(map[string]int, mapSize),
			f4:  make(map[string]int, mapSize),
			f5:  make(map[string]int, mapSize),
			f6:  make(map[string]int, mapSize),
			f7:  make(map[string]int, mapSize),
			f8:  make(map[string]int, mapSize),
			f9:  make(map[string]int, mapSize),
			f10: make(map[string]int, mapSize),
			f11: make(map[string]int, mapSize),
			f12: make(map[string]int, mapSize),
			f13: make(map[string]int, mapSize),
		}

		r.f0["x"] = 10

		size := unsafe.Sizeof(r)
		b.Run(fmt.Sprintf("size=%d", size), func(b *testing.B) {
			refSum = referencePassByPtr14(&r)
		})
	})
	b.Run("method=value/fields=15", func(b *testing.B) {
		r := R15{
			f0:  make(map[string]int, mapSize),
			f1:  make(map[string]int, mapSize),
			f2:  make(map[string]int, mapSize),
			f3:  make(map[string]int, mapSize),
			f4:  make(map[string]int, mapSize),
			f5:  make(map[string]int, mapSize),
			f6:  make(map[string]int, mapSize),
			f7:  make(map[string]int, mapSize),
			f8:  make(map[string]int, mapSize),
			f9:  make(map[string]int, mapSize),
			f10: make(map[string]int, mapSize),
			f11: make(map[string]int, mapSize),
			f12: make(map[string]int, mapSize),
			f13: make(map[string]int, mapSize),
			f14: make(map[string]int, mapSize),
		}

		r.f0["x"] = 10

		size := unsafe.Sizeof(r)
		b.Run(fmt.Sprintf("size=%d", size), func(b *testing.B) {
			refSum = referencePassByValue15(r)
		})
	})

	b.Run("method=ptr/fields=1", func(b *testing.B) {
		r := R15{
			f0:  make(map[string]int, mapSize),
			f1:  make(map[string]int, mapSize),
			f2:  make(map[string]int, mapSize),
			f3:  make(map[string]int, mapSize),
			f4:  make(map[string]int, mapSize),
			f5:  make(map[string]int, mapSize),
			f6:  make(map[string]int, mapSize),
			f7:  make(map[string]int, mapSize),
			f8:  make(map[string]int, mapSize),
			f9:  make(map[string]int, mapSize),
			f10: make(map[string]int, mapSize),
			f11: make(map[string]int, mapSize),
			f12: make(map[string]int, mapSize),
			f13: make(map[string]int, mapSize),
			f14: make(map[string]int, mapSize),
		}

		r.f0["x"] = 10

		size := unsafe.Sizeof(r)
		b.Run(fmt.Sprintf("size=%d", size), func(b *testing.B) {
			refSum = referencePassByPtr15(&r)
		})
	})
}
