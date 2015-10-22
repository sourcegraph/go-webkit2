package webkit2

// #include "cairo_endianes.h"
import "C"
import (
	"unsafe"
)

func CairoEndianDependedARGB32ToRGBA(data []byte, out []byte) {
	C.gowk2_cairo_endian_depended_ARGB32_to_RGBA((*C.uchar)(unsafe.Pointer(&data[0])), (*C.uchar)(unsafe.Pointer(&out[0])), (C.uint)(len(data)))
}

func CairoEndianDependedARGB32ToRGBASmart(data []byte, out []byte) {
	C.gowk2_cairo_endian_depended_ARGB32_to_RGBA_smart((*C.uchar)(unsafe.Pointer(&data[0])), (*C.uchar)(unsafe.Pointer(&out[0])), (C.uint)(len(data)))
}
