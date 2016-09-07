package webkit2

import (
	"runtime"

	"github.com/visionect/gotk3/gtk"
)

func init() {
	runtime.LockOSThread()
	gtk.Init(nil)
}
