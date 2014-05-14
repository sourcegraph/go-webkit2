package webkit2_test

import (
	"fmt"
	"github.com/sqs/gojs"
	"github.com/visionect/go-webkit2/webkit2"
	"github.com/visionect/gotk3/glib"
	"github.com/visionect/gotk3/gtk"
	"runtime"
)

func Example() {
	runtime.LockOSThread()
	gtk.Init(nil)

	webView := webkit2.NewWebView()
	defer webView.Destroy()

	webView.Connect("load-failed", func() {
		fmt.Println("Load failed.")
	})
	webView.Connect("load-changed", func(_ *glib.Object, event int) {
		loadEvent := webkit2.LoadEvent(event)
		switch loadEvent {
		case webkit2.LoadFinished:
			fmt.Println("Load finished.")
			fmt.Printf("Title: %q\n", webView.Title())
			fmt.Printf("URI: %s\n", webView.URI())
			webView.RunJavaScript("window.location.hostname", func(val *gojs.Value, err error) {
				if err != nil {
					fmt.Println("JavaScript error.")
				} else {
					fmt.Printf("Hostname (from JavaScript): %q\n", val)
				}
				gtk.MainQuit()
			})
		}
	})

	glib.IdleAdd(func() bool {
		webView.LoadURI("https://www.google.com/")
		return false
	})

	gtk.Main()

	// output:
	// Load finished.
	// Title: "Google"
	// URI: https://www.google.com/
	// Hostname (from JavaScript): "www.google.com"
}
