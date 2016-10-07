package webkit2

// #include <gio/gio.h>
// #include "gasyncreadycallback.go.h"
import "C"
import (
	"errors"
	"reflect"
	"unsafe"
)

type garCallback struct {
	f reflect.Value
}

//export _go_gasyncreadycallback_call
func _go_gasyncreadycallback_call(cbinfoRaw C.gpointer, cresult unsafe.Pointer) {
	result := (*C.GAsyncResult)(cresult)
	cbinfo := (*garCallback)(unsafe.Pointer(cbinfoRaw))
	cbinfo.f.Call([]reflect.Value{reflect.ValueOf(result)})
}

func newGAsyncReadyCallback(f interface{}) (cCallback C.GAsyncReadyCallback, userData C.gpointer, err error) {
	rf := reflect.ValueOf(f)
	if rf.Kind() != reflect.Func {
		return nil, nil, errors.New("f is not a function")
	}
	data := C.malloc(C.size_t(unsafe.Sizeof(garCallback{})))
	cbinfo := (*garCallback)(data)
	cbinfo.f = rf
	return C.GAsyncReadyCallback(C._gasyncreadycallback_call), C.gpointer(unsafe.Pointer(cbinfo)), nil
}
