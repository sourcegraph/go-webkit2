package webkit2

// #include <gio/gio.h>
// #include "gasyncreadycallback.go.h"
import "C"
import (
	"errors"
	"reflect"
	"unsafe"
	"sync"
)

type garCallback struct {
	f reflect.Value
}

var (
	//Map stores callback pointers to protect callbacks from GC.
	CallbackProtectMap map[C.gpointer]*garCallback
	ProtectMapLock	   sync.RWMutex
)

func init() {
	CallbackProtectMap = make(map[C.gpointer]*garCallback)
}

//export _go_gasyncreadycallback_call
func _go_gasyncreadycallback_call(cbinfoRaw C.gpointer, cresult unsafe.Pointer) {
	result := (*C.GAsyncResult)(cresult)
	cbinfo := (*garCallback)(unsafe.Pointer(cbinfoRaw))
	cbinfo.f.Call([]reflect.Value{reflect.ValueOf(result)})
	// protect callback from Garbage collection
	ProtectMapLock.Lock()
	delete(CallbackProtectMap, cbinfoRaw)
	ProtectMapLock.Unlock()
}

func newGAsyncReadyCallback(f interface{}) (cCallback C.GAsyncReadyCallback, userData C.gpointer, err error) {
	rf := reflect.ValueOf(f)
	if rf.Kind() != reflect.Func {
		return nil, nil, errors.New("f is not a function")
	}
	cbinfo := &garCallback{rf}
	cbinfoRaw := C.gpointer(unsafe.Pointer(cbinfo))
	// protect callback from Garbage collection
	ProtectMapLock.Lock()
	CallbackProtectMap[cbinfoRaw] = cbinfo
	ProtectMapLock.Unlock()
	return C.GAsyncReadyCallback(C._gasyncreadycallback_call), cbinfoRaw, nil
}
