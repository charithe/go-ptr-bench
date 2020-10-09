package bench

import (
	"fmt"
	"testing"
	"unsafe"
)

var sum uint64

type S1 struct {
	f0 uint64
}

//go:noinline
func primitivePassByValue1(s S1) uint64 {
	return s.f0 * 2
}

//go:noinline
func primitivePassByPtr1(s *S1) uint64 {
	return s.f0 * 2
}

type S2 struct {
	f0 uint64
	f1 uint64
}

//go:noinline
func primitivePassByValue2(s S2) uint64 {
	return s.f0 * 2
}

//go:noinline
func primitivePassByPtr2(s *S2) uint64 {
	return s.f0 * 2
}

type S3 struct {
	f0 uint64
	f1 uint64
	f2 uint64
}

//go:noinline
func primitivePassByValue3(s S3) uint64 {
	return s.f0 * 2
}

//go:noinline
func primitivePassByPtr3(s *S3) uint64 {
	return s.f0 * 2
}

type S4 struct {
	f0 uint64
	f1 uint64
	f2 uint64
	f3 uint64
}

//go:noinline
func primitivePassByValue4(s S4) uint64 {
	return s.f0 * 2
}

//go:noinline
func primitivePassByPtr4(s *S4) uint64 {
	return s.f0 * 2
}

type S5 struct {
	f0 uint64
	f1 uint64
	f2 uint64
	f3 uint64
	f4 uint64
}

//go:noinline
func primitivePassByValue5(s S5) uint64 {
	return s.f0 * 2
}

//go:noinline
func primitivePassByPtr5(s *S5) uint64 {
	return s.f0 * 2
}

type S6 struct {
	f0 uint64
	f1 uint64
	f2 uint64
	f3 uint64
	f4 uint64
	f5 uint64
}

//go:noinline
func primitivePassByValue6(s S6) uint64 {
	return s.f0 * 2
}

//go:noinline
func primitivePassByPtr6(s *S6) uint64 {
	return s.f0 * 2
}

type S7 struct {
	f0 uint64
	f1 uint64
	f2 uint64
	f3 uint64
	f4 uint64
	f5 uint64
	f6 uint64
}

//go:noinline
func primitivePassByValue7(s S7) uint64 {
	return s.f0 * 2
}

//go:noinline
func primitivePassByPtr7(s *S7) uint64 {
	return s.f0 * 2
}

type S8 struct {
	f0 uint64
	f1 uint64
	f2 uint64
	f3 uint64
	f4 uint64
	f5 uint64
	f6 uint64
	f7 uint64
}

//go:noinline
func primitivePassByValue8(s S8) uint64 {
	return s.f0 * 2
}

//go:noinline
func primitivePassByPtr8(s *S8) uint64 {
	return s.f0 * 2
}

type S9 struct {
	f0 uint64
	f1 uint64
	f2 uint64
	f3 uint64
	f4 uint64
	f5 uint64
	f6 uint64
	f7 uint64
	f8 uint64
}

//go:noinline
func primitivePassByValue9(s S9) uint64 {
	return s.f0 * 2
}

//go:noinline
func primitivePassByPtr9(s *S9) uint64 {
	return s.f0 * 2
}

type S10 struct {
	f0 uint64
	f1 uint64
	f2 uint64
	f3 uint64
	f4 uint64
	f5 uint64
	f6 uint64
	f7 uint64
	f8 uint64
	f9 uint64
}

//go:noinline
func primitivePassByValue10(s S10) uint64 {
	return s.f0 * 2
}

//go:noinline
func primitivePassByPtr10(s *S10) uint64 {
	return s.f0 * 2
}

type S11 struct {
	f0  uint64
	f1  uint64
	f2  uint64
	f3  uint64
	f4  uint64
	f5  uint64
	f6  uint64
	f7  uint64
	f8  uint64
	f9  uint64
	f10 uint64
}

//go:noinline
func primitivePassByValue11(s S11) uint64 {
	return s.f0 * 2
}

//go:noinline
func primitivePassByPtr11(s *S11) uint64 {
	return s.f0 * 2
}

type S12 struct {
	f0  uint64
	f1  uint64
	f2  uint64
	f3  uint64
	f4  uint64
	f5  uint64
	f6  uint64
	f7  uint64
	f8  uint64
	f9  uint64
	f10 uint64
	f11 uint64
}

//go:noinline
func primitivePassByValue12(s S12) uint64 {
	return s.f0 * 2
}

//go:noinline
func primitivePassByPtr12(s *S12) uint64 {
	return s.f0 * 2
}

type S13 struct {
	f0  uint64
	f1  uint64
	f2  uint64
	f3  uint64
	f4  uint64
	f5  uint64
	f6  uint64
	f7  uint64
	f8  uint64
	f9  uint64
	f10 uint64
	f11 uint64
	f12 uint64
}

//go:noinline
func primitivePassByValue13(s S13) uint64 {
	return s.f0 * 2
}

//go:noinline
func primitivePassByPtr13(s *S13) uint64 {
	return s.f0 * 2
}

type S14 struct {
	f0  uint64
	f1  uint64
	f2  uint64
	f3  uint64
	f4  uint64
	f5  uint64
	f6  uint64
	f7  uint64
	f8  uint64
	f9  uint64
	f10 uint64
	f11 uint64
	f12 uint64
	f13 uint64
}

//go:noinline
func primitivePassByValue14(s S14) uint64 {
	return s.f0 * 2
}

//go:noinline
func primitivePassByPtr14(s *S14) uint64 {
	return s.f0 * 2
}

type S15 struct {
	f0  uint64
	f1  uint64
	f2  uint64
	f3  uint64
	f4  uint64
	f5  uint64
	f6  uint64
	f7  uint64
	f8  uint64
	f9  uint64
	f10 uint64
	f11 uint64
	f12 uint64
	f13 uint64
	f14 uint64
}

//go:noinline
func primitivePassByValue15(s S15) uint64 {
	return s.f0 * 2
}

//go:noinline
func primitivePassByPtr15(s *S15) uint64 {
	return s.f0 * 2
}

func BenchmarkPrimitiveFields(b *testing.B) {
	b.ReportAllocs()
	b.Run("method=value/fields=1", func(b *testing.B) {
		s := S1{
			f0: 0,
		}

		size := unsafe.Sizeof(s)
		b.Run(fmt.Sprintf("size=%d", size), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				sum = primitivePassByValue1(s)
			}
		})

	})

	b.Run("method=ptr/fields=1", func(b *testing.B) {
		s := S1{
			f0: 0,
		}

		size := unsafe.Sizeof(s)
		b.Run(fmt.Sprintf("size=%d", size), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				sum = primitivePassByPtr1(&s)
			}
		})
	})
	b.Run("method=value/fields=2", func(b *testing.B) {
		s := S2{
			f0: 0,
			f1: 1,
		}

		size := unsafe.Sizeof(s)
		b.Run(fmt.Sprintf("size=%d", size), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				sum = primitivePassByValue2(s)
			}
		})

	})

	b.Run("method=ptr/fields=2", func(b *testing.B) {
		s := S2{
			f0: 0,
			f1: 1,
		}

		size := unsafe.Sizeof(s)
		b.Run(fmt.Sprintf("size=%d", size), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				sum = primitivePassByPtr2(&s)
			}
		})
	})
	b.Run("method=value/fields=3", func(b *testing.B) {
		s := S3{
			f0: 0,
			f1: 1,
			f2: 2,
		}

		size := unsafe.Sizeof(s)
		b.Run(fmt.Sprintf("size=%d", size), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				sum = primitivePassByValue3(s)
			}
		})

	})

	b.Run("method=ptr/fields=3", func(b *testing.B) {
		s := S3{
			f0: 0,
			f1: 1,
			f2: 2,
		}

		size := unsafe.Sizeof(s)
		b.Run(fmt.Sprintf("size=%d", size), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				sum = primitivePassByPtr3(&s)
			}
		})
	})
	b.Run("method=value/fields=4", func(b *testing.B) {
		s := S4{
			f0: 0,
			f1: 1,
			f2: 2,
			f3: 3,
		}

		size := unsafe.Sizeof(s)
		b.Run(fmt.Sprintf("size=%d", size), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				sum = primitivePassByValue4(s)
			}
		})

	})

	b.Run("method=ptr/fields=4", func(b *testing.B) {
		s := S4{
			f0: 0,
			f1: 1,
			f2: 2,
			f3: 3,
		}

		size := unsafe.Sizeof(s)
		b.Run(fmt.Sprintf("size=%d", size), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				sum = primitivePassByPtr4(&s)
			}
		})
	})
	b.Run("method=value/fields=5", func(b *testing.B) {
		s := S5{
			f0: 0,
			f1: 1,
			f2: 2,
			f3: 3,
			f4: 4,
		}

		size := unsafe.Sizeof(s)
		b.Run(fmt.Sprintf("size=%d", size), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				sum = primitivePassByValue5(s)
			}
		})

	})

	b.Run("method=ptr/fields=5", func(b *testing.B) {
		s := S5{
			f0: 0,
			f1: 1,
			f2: 2,
			f3: 3,
			f4: 4,
		}

		size := unsafe.Sizeof(s)
		b.Run(fmt.Sprintf("size=%d", size), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				sum = primitivePassByPtr5(&s)
			}
		})
	})
	b.Run("method=value/fields=6", func(b *testing.B) {
		s := S6{
			f0: 0,
			f1: 1,
			f2: 2,
			f3: 3,
			f4: 4,
			f5: 5,
		}

		size := unsafe.Sizeof(s)
		b.Run(fmt.Sprintf("size=%d", size), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				sum = primitivePassByValue6(s)
			}
		})

	})

	b.Run("method=ptr/fields=6", func(b *testing.B) {
		s := S6{
			f0: 0,
			f1: 1,
			f2: 2,
			f3: 3,
			f4: 4,
			f5: 5,
		}

		size := unsafe.Sizeof(s)
		b.Run(fmt.Sprintf("size=%d", size), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				sum = primitivePassByPtr6(&s)
			}
		})
	})
	b.Run("method=value/fields=7", func(b *testing.B) {
		s := S7{
			f0: 0,
			f1: 1,
			f2: 2,
			f3: 3,
			f4: 4,
			f5: 5,
			f6: 6,
		}

		size := unsafe.Sizeof(s)
		b.Run(fmt.Sprintf("size=%d", size), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				sum = primitivePassByValue7(s)
			}
		})

	})

	b.Run("method=ptr/fields=7", func(b *testing.B) {
		s := S7{
			f0: 0,
			f1: 1,
			f2: 2,
			f3: 3,
			f4: 4,
			f5: 5,
			f6: 6,
		}

		size := unsafe.Sizeof(s)
		b.Run(fmt.Sprintf("size=%d", size), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				sum = primitivePassByPtr7(&s)
			}
		})
	})
	b.Run("method=value/fields=8", func(b *testing.B) {
		s := S8{
			f0: 0,
			f1: 1,
			f2: 2,
			f3: 3,
			f4: 4,
			f5: 5,
			f6: 6,
			f7: 7,
		}

		size := unsafe.Sizeof(s)
		b.Run(fmt.Sprintf("size=%d", size), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				sum = primitivePassByValue8(s)
			}
		})

	})

	b.Run("method=ptr/fields=8", func(b *testing.B) {
		s := S8{
			f0: 0,
			f1: 1,
			f2: 2,
			f3: 3,
			f4: 4,
			f5: 5,
			f6: 6,
			f7: 7,
		}

		size := unsafe.Sizeof(s)
		b.Run(fmt.Sprintf("size=%d", size), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				sum = primitivePassByPtr8(&s)
			}
		})
	})
	b.Run("method=value/fields=9", func(b *testing.B) {
		s := S9{
			f0: 0,
			f1: 1,
			f2: 2,
			f3: 3,
			f4: 4,
			f5: 5,
			f6: 6,
			f7: 7,
			f8: 8,
		}

		size := unsafe.Sizeof(s)
		b.Run(fmt.Sprintf("size=%d", size), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				sum = primitivePassByValue9(s)
			}
		})

	})

	b.Run("method=ptr/fields=9", func(b *testing.B) {
		s := S9{
			f0: 0,
			f1: 1,
			f2: 2,
			f3: 3,
			f4: 4,
			f5: 5,
			f6: 6,
			f7: 7,
			f8: 8,
		}

		size := unsafe.Sizeof(s)
		b.Run(fmt.Sprintf("size=%d", size), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				sum = primitivePassByPtr9(&s)
			}
		})
	})
	b.Run("method=value/fields=10", func(b *testing.B) {
		s := S10{
			f0: 0,
			f1: 1,
			f2: 2,
			f3: 3,
			f4: 4,
			f5: 5,
			f6: 6,
			f7: 7,
			f8: 8,
			f9: 9,
		}

		size := unsafe.Sizeof(s)
		b.Run(fmt.Sprintf("size=%d", size), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				sum = primitivePassByValue10(s)
			}
		})

	})

	b.Run("method=ptr/fields=10", func(b *testing.B) {
		s := S10{
			f0: 0,
			f1: 1,
			f2: 2,
			f3: 3,
			f4: 4,
			f5: 5,
			f6: 6,
			f7: 7,
			f8: 8,
			f9: 9,
		}

		size := unsafe.Sizeof(s)
		b.Run(fmt.Sprintf("size=%d", size), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				sum = primitivePassByPtr10(&s)
			}
		})
	})
	b.Run("method=value/fields=11", func(b *testing.B) {
		s := S11{
			f0:  0,
			f1:  1,
			f2:  2,
			f3:  3,
			f4:  4,
			f5:  5,
			f6:  6,
			f7:  7,
			f8:  8,
			f9:  9,
			f10: 10,
		}

		size := unsafe.Sizeof(s)
		b.Run(fmt.Sprintf("size=%d", size), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				sum = primitivePassByValue11(s)
			}
		})

	})

	b.Run("method=ptr/fields=11", func(b *testing.B) {
		s := S11{
			f0:  0,
			f1:  1,
			f2:  2,
			f3:  3,
			f4:  4,
			f5:  5,
			f6:  6,
			f7:  7,
			f8:  8,
			f9:  9,
			f10: 10,
		}

		size := unsafe.Sizeof(s)
		b.Run(fmt.Sprintf("size=%d", size), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				sum = primitivePassByPtr11(&s)
			}
		})
	})
	b.Run("method=value/fields=12", func(b *testing.B) {
		s := S12{
			f0:  0,
			f1:  1,
			f2:  2,
			f3:  3,
			f4:  4,
			f5:  5,
			f6:  6,
			f7:  7,
			f8:  8,
			f9:  9,
			f10: 10,
			f11: 11,
		}

		size := unsafe.Sizeof(s)
		b.Run(fmt.Sprintf("size=%d", size), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				sum = primitivePassByValue12(s)
			}
		})

	})

	b.Run("method=ptr/fields=12", func(b *testing.B) {
		s := S12{
			f0:  0,
			f1:  1,
			f2:  2,
			f3:  3,
			f4:  4,
			f5:  5,
			f6:  6,
			f7:  7,
			f8:  8,
			f9:  9,
			f10: 10,
			f11: 11,
		}

		size := unsafe.Sizeof(s)
		b.Run(fmt.Sprintf("size=%d", size), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				sum = primitivePassByPtr12(&s)
			}
		})
	})
	b.Run("method=value/fields=13", func(b *testing.B) {
		s := S13{
			f0:  0,
			f1:  1,
			f2:  2,
			f3:  3,
			f4:  4,
			f5:  5,
			f6:  6,
			f7:  7,
			f8:  8,
			f9:  9,
			f10: 10,
			f11: 11,
			f12: 12,
		}

		size := unsafe.Sizeof(s)
		b.Run(fmt.Sprintf("size=%d", size), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				sum = primitivePassByValue13(s)
			}
		})

	})

	b.Run("method=ptr/fields=13", func(b *testing.B) {
		s := S13{
			f0:  0,
			f1:  1,
			f2:  2,
			f3:  3,
			f4:  4,
			f5:  5,
			f6:  6,
			f7:  7,
			f8:  8,
			f9:  9,
			f10: 10,
			f11: 11,
			f12: 12,
		}

		size := unsafe.Sizeof(s)
		b.Run(fmt.Sprintf("size=%d", size), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				sum = primitivePassByPtr13(&s)
			}
		})
	})
	b.Run("method=value/fields=14", func(b *testing.B) {
		s := S14{
			f0:  0,
			f1:  1,
			f2:  2,
			f3:  3,
			f4:  4,
			f5:  5,
			f6:  6,
			f7:  7,
			f8:  8,
			f9:  9,
			f10: 10,
			f11: 11,
			f12: 12,
			f13: 13,
		}

		size := unsafe.Sizeof(s)
		b.Run(fmt.Sprintf("size=%d", size), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				sum = primitivePassByValue14(s)
			}
		})

	})

	b.Run("method=ptr/fields=14", func(b *testing.B) {
		s := S14{
			f0:  0,
			f1:  1,
			f2:  2,
			f3:  3,
			f4:  4,
			f5:  5,
			f6:  6,
			f7:  7,
			f8:  8,
			f9:  9,
			f10: 10,
			f11: 11,
			f12: 12,
			f13: 13,
		}

		size := unsafe.Sizeof(s)
		b.Run(fmt.Sprintf("size=%d", size), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				sum = primitivePassByPtr14(&s)
			}
		})
	})
	b.Run("method=value/fields=15", func(b *testing.B) {
		s := S15{
			f0:  0,
			f1:  1,
			f2:  2,
			f3:  3,
			f4:  4,
			f5:  5,
			f6:  6,
			f7:  7,
			f8:  8,
			f9:  9,
			f10: 10,
			f11: 11,
			f12: 12,
			f13: 13,
			f14: 14,
		}

		size := unsafe.Sizeof(s)
		b.Run(fmt.Sprintf("size=%d", size), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				sum = primitivePassByValue15(s)
			}
		})

	})

	b.Run("method=ptr/fields=15", func(b *testing.B) {
		s := S15{
			f0:  0,
			f1:  1,
			f2:  2,
			f3:  3,
			f4:  4,
			f5:  5,
			f6:  6,
			f7:  7,
			f8:  8,
			f9:  9,
			f10: 10,
			f11: 11,
			f12: 12,
			f13: 13,
			f14: 14,
		}

		size := unsafe.Sizeof(s)
		b.Run(fmt.Sprintf("size=%d", size), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				sum = primitivePassByPtr15(&s)
			}
		})
	})
}
