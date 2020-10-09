package bench

import (
    "fmt"
    "testing"
    "unsafe"
)

var sum uint64

{{- range $i := until .NumFields }} 
{{- $i = add1 $i | int }}
type S{{ $i }} struct {
	{{- range $n := until $i }}
	f{{ $n }} uint64
	{{- end }}
}

//go:noinline
func primitivePassByValue{{ $i }}(s S{{ $i }}) uint64 {
	return s.f0 * 2
}

//go:noinline
func primitivePassByPtr{{ $i }}(s *S{{ $i }}) uint64 {
    return s.f0 * 2
}
{{- end }}


func BenchmarkPrimitiveFields(b *testing.B) {
	b.ReportAllocs()

{{- range $i := until .NumFields }}
{{- $i = add1 $i | int }}
	b.Run("method=value/fields={{ $i }}", func(b *testing.B) {
		s := S{{ $i }}{
		{{- range $n := until $i }}
		   f{{ $n }}: {{ $n }},
		{{- end }}
	    }

        size := unsafe.Sizeof(s)
        b.Run(fmt.Sprintf("size=%d", size), func(b *testing.B){
            for i := 0; i < b.N; i++ {
                sum = primitivePassByValue{{ $i }}(s)
            }
        })

	})

	b.Run("method=ptr/fields={{ $i }}", func(b *testing.B) {
		s := S{{ $i }}{
		{{- range $n := until $i }}
		   f{{ $n }}: {{ $n }},
		{{- end }}
	    }

        size := unsafe.Sizeof(s)
        b.Run(fmt.Sprintf("size=%d", size), func(b *testing.B){
            for i := 0; i < b.N; i++ {
                sum = primitivePassByPtr{{ $i }}(&s)
            }
        })
	})
{{- end }}
}

