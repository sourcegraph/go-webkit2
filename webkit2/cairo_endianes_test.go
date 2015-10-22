package webkit2

import (
	"math/rand"
	"reflect"
	"testing"
)

func generateImageARGB32(w, h int) []byte {
	r := rand.New(rand.NewSource(666))
	pix := make([]byte, 4*w*h)
	for i, _ := range pix {
		pix[i] = byte(r.Intn(255))
	}
	return pix
}

func TestEndianDependeARGB32ToRGBA(t *testing.T) {
	pix := generateImageARGB32(1800, 1600)
	pix2 := make([]byte, len(pix))
	copy(pix2, pix)
	t.Log(pix2[:12])
	CairoEndianDependedARGB32ToRGBA(pix, pix)
	CairoEndianDependedARGB32ToRGBASmart(pix2, pix2)
	if !reflect.DeepEqual(pix, pix2) {
		t.Fatal("Endianes correction function comparison mismatch!")
	}
	t.Log(pix[:12])
	t.Log(pix2[:12])
}

func BenchmarkCairoEndianDependeARGB32ToRGBA(b *testing.B) {
	pix := generateImageARGB32(1800, 1600)
	rgba := make([]uint8, len(pix))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		CairoEndianDependedARGB32ToRGBA(pix, rgba)
	}
}

func BenchmarkCairoEndianDependeARGB32ToRGBASamrt(b *testing.B) {
	pix := generateImageARGB32(1800, 1600)
	rgba := make([]uint8, len(pix))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		CairoEndianDependedARGB32ToRGBASmart(pix, rgba)
	}
}
