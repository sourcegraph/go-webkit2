package webkit2

import (
	"runtime"

	"github.com/conformal/gotk3/gtk"
)

func init() {
	runtime.LockOSThread()
	gtk.Init(nil)
}
