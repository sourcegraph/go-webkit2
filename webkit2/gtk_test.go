package webkit2

import (
	"github.com/sqs/gotk3/gtk"
	"runtime"
)

func init() {
	runtime.LockOSThread()
	gtk.Init(nil)
}
