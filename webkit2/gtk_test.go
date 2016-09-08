package webkit2

import (
	"runtime"

	"github.com/gotk3/gotk3/gtk"
)

func init() {
	runtime.LockOSThread()
	gtk.Init(nil)
}
