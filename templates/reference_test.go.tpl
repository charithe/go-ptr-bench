package bench

import (
	"fmt"
	"testing"
	"unsafe"
)

const mapSize = 10
var refSum = 0

{{- range $i := until .NumFields }}
{{- $i = add1 $i | int }}
type R{{ $i }} struct {
	{{- range $n := until $i }}
	f{{ $n }} map[string]int
	{{- end }}
}

//go:noinline
func referencePassByValue{{ $i }}(r R{{ $i }}) int {
	return r.f0["key"]
}

//go:noinline
func referencePassByPtr{{ $i }}(r *R{{ $i }}) int {
	return r.f0["key"]
}
{{- end }}

func BenchmarkReferenceFields(b *testing.B) {
	b.ReportAllocs()

	{{- range $i := until .NumFields }}
	{{- $i = add1 $i | int }}
	b.Run("method=value/fields={{ $i }}", func(b *testing.B) {
		r := R{{ $i }}{
			{{- range $n := until $i }}
			f{{ $n }}: make(map[string]int, mapSize),
			{{- end }}
		}

        r.f0["x"] = 10

		size := unsafe.Sizeof(r)
		b.Run(fmt.Sprintf("size=%d", size), func(b *testing.B) {
			refSum = referencePassByValue{{ $i }}(r)
		})
	})

	b.Run("method=ptr/fields=1", func(b *testing.B) {
		r := R{{ $i }}{
			{{- range $n := until $i }}
			f{{ $n }}: make(map[string]int, mapSize),
			{{- end }}
		}

        r.f0["x"] = 10

		size := unsafe.Sizeof(r)
		b.Run(fmt.Sprintf("size=%d", size), func(b *testing.B) {
			refSum = referencePassByPtr{{ $i }}(&r)
		})
	})
	{{- end }}
}
