package webkit2

import (
	"github.com/sqs/gotk3/glib"
	"github.com/sqs/gotk3/gtk"
	"net/http"
	"runtime"
	"testing"
)

func init() {
	runtime.LockOSThread()
	gtk.Init(nil)
}

func TestNewWebView(t *testing.T) {
	webView := NewWebView()
	defer webView.Destroy()
}

func TestNewWebViewWithContext(t *testing.T) {
	cx := DefaultWebContext()
	webView := NewWebViewWithContext(cx)
	defer webView.Destroy()
}

func TestWebView_LoadURI(t *testing.T) {
	setup()
	defer teardown()

	responseOk := false
	mux.HandleFunc("/", func(w http.ResponseWriter, _ *http.Request) {
		w.Write([]byte("abc"))
		responseOk = true
	})

	loadFinished := false
	webView.Connect("load-failed", func() {
		t.Errorf("load failed")
	})
	webView.Connect("load-changed", func(ctx *glib.CallbackContext) {
		loadEvent := LoadEvent(ctx.Arg(0).Int())
		switch loadEvent {
		case LoadFinished:
			loadFinished = true
			gtk.MainQuit()
		}
	})

	glib.IdleAdd(func() bool {
		webView.LoadURI(server.URL)
		return false
	})

	gtk.Main()

	if !responseOk {
		t.Error("!responseOk")
	}
	if !loadFinished {
		t.Error("!loadFinished")
	}
}

func TestWebView_LoadURI_load_failed(t *testing.T) {
	webView := NewWebView()
	defer webView.Destroy()

	loadFailed := false
	loadFinished := false
	webView.Connect("load-failed", func() {
		loadFailed = true
	})
	webView.Connect("load-changed", func(ctx *glib.CallbackContext) {
		loadEvent := LoadEvent(ctx.Arg(0).Int())
		switch loadEvent {
		case LoadFinished:
			loadFinished = true
			gtk.MainQuit()
		}
	})

	glib.IdleAdd(func() bool {
		// Load a bad URL to trigger load failure.
		webView.LoadURI("http://127.0.0.1:99999")
		return false
	})

	gtk.Main()

	if !loadFailed {
		t.Error("!loadFailed")
	}
	if !loadFinished {
		t.Error("!loadFinished")
	}
}

func TestWebView_LoadHTML(t *testing.T) {
	webView := NewWebView()
	defer webView.Destroy()

	loadOk := false
	webView.Connect("load-failed", func() {
		t.Errorf("load failed")
	})
	webView.Connect("load-changed", func(ctx *glib.CallbackContext) {
		loadEvent := LoadEvent(ctx.Arg(0).Int())
		switch loadEvent {
		case LoadFinished:
			loadOk = true
			gtk.MainQuit()
		}
	})

	glib.IdleAdd(func() bool {
		webView.LoadHTML("<p>hello</p>", "")
		return false
	})

	gtk.Main()

	if !loadOk {
		t.Error("!loadOk")
	}
}

func TestWebView_Title(t *testing.T) {
	webView := NewWebView()
	defer webView.Destroy()

	wantTitle := "foo"
	var gotTitle string
	webView.Connect("notify::title", func() {
		glib.IdleAdd(func() bool {
			gotTitle = webView.Title()
			if gotTitle != "" {
				gtk.MainQuit()
			}
			return false
		})
	})

	glib.IdleAdd(func() bool {
		webView.LoadHTML("<html><head><title>"+wantTitle+"</title></head><body></body></html>", "")
		return false
	})

	gtk.Main()

	if wantTitle != gotTitle {
		t.Errorf("want title %q, got %q", wantTitle, gotTitle)
	}
}

func TestWebView_URI(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/", func(w http.ResponseWriter, _ *http.Request) {})

	wantURI := server.URL + "/"
	var gotURI string
	webView.Connect("notify::uri", func(ctx *glib.CallbackContext) {
		glib.IdleAdd(func() bool {
			gotURI = webView.URI()
			if gotURI != "" {
				gtk.MainQuit()
			}
			return false
		})
	})

	glib.IdleAdd(func() bool {
		webView.LoadURI(server.URL)
		return false
	})

	gtk.Main()

	if wantURI != gotURI {
		t.Errorf("want URI %q, got %q", wantURI, gotURI)
	}
}
