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
	//Map stores callbacks pointers, to protect them from GC.
	callbackProtectMap map[C.gpointer]*garCallback
	protectMapLock	   sync.RWMutex
)

func init() {
	callbackProtectMap = make(map[C.gpointer]*garCallback)
}

//export _go_gasyncreadycallback_call
func _go_gasyncreadycallback_call(cbinfoRaw C.gpointer, cresult unsafe.Pointer) {
	result := (*C.GAsyncResult)(cresult)
	cbinfo := (*garCallback)(unsafe.Pointer(cbinfoRaw))
	cbinfo.f.Call([]reflect.Value{reflect.ValueOf(result)})
	// protect callback from Garbage collection
	protectMapLock.Lock()
	delete(callbackProtectMap, cbinfoRaw)
	protectMapLock.Unlock()
}

func newGAsyncReadyCallback(f interface{}) (cCallback C.GAsyncReadyCallback, userData C.gpointer, err error) {
	rf := reflect.ValueOf(f)
	if rf.Kind() != reflect.Func {
		return nil, nil, errors.New("f is not a function")
	}
	cbinfo := &garCallback{rf}
	cbinfoRaw := C.gpointer(unsafe.Pointer(cbinfo))
	// protect callback from Garbage collection
	protectMapLock.Lock()
	callbackProtectMap[cbinfoRaw] = cbinfo
	protectMapLock.Unlock()
	return C.GAsyncReadyCallback(C._gasyncreadycallback_call), cbinfoRaw, nil
}
