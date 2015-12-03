package webkit2

import (
	"errors"
	"image"
	"net/http"
	"reflect"
	"testing"

	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
	"github.com/sqs/gojs"
)

func TestNewWebView(t *testing.T) {
	webView := NewWebView()
	defer webView.Destroy()
}

func TestNewWebViewWithContext(t *testing.T) {
	cx := DefaultWebContext()
	webView := NewWebViewWithContext(cx)
	defer webView.Destroy()
}

func TestWebView_Context(t *testing.T) {
	webView := NewWebView()
	defer webView.Destroy()
	webView.Context()
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
	webView.Connect("load-changed", func(_ *glib.Object, loadEvent LoadEvent) {
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
	webView.Connect("load-changed", func(_ *glib.Object, loadEvent LoadEvent) {
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
	webView.Connect("load-changed", func(_ *glib.Object, loadEvent LoadEvent) {
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
	webView.Connect("notify::uri", func() {
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

func TestWebView_Settings(t *testing.T) {
	webView := NewWebView()
	defer webView.Destroy()

	webView.Settings()
}

func TestWebView_JavaScriptGlobalContext(t *testing.T) {
	webView := NewWebView()
	defer webView.Destroy()

	webView.JavaScriptGlobalContext()
}

func TestWebView_RunJavaScript(t *testing.T) {
	webView := NewWebView()
	defer webView.Destroy()

	wantResultString := "abc"
	webView.Connect("load-changed", func(_ *glib.Object, loadEvent LoadEvent) {
		switch loadEvent {
		case LoadFinished:
			webView.RunJavaScript(`document.getElementById("foo").innerHTML`, func(result *gojs.Value, err error) {
				if err != nil {
					t.Errorf("RunJavaScript error: %s", err)
				}
				resultString := webView.JavaScriptGlobalContext().ToStringOrDie(result)
				if wantResultString != resultString {
					t.Errorf("want result string %q, got %q", wantResultString, resultString)
				}
				gtk.MainQuit()
			})
		}
	})

	glib.IdleAdd(func() bool {
		webView.LoadHTML(`<p id=foo>abc</p>`, "")
		return false
	})

	gtk.Main()
}

func TestWebView_RunJavaScript_exception(t *testing.T) {
	webView := NewWebView()
	defer webView.Destroy()

	wantErr := errors.New("An exception was raised in JavaScript")
	webView.Connect("load-changed", func(_ *glib.Object, loadEvent LoadEvent) {
		switch loadEvent {
		case LoadFinished:
			webView.RunJavaScript(`throw new Error("foo")`, func(result *gojs.Value, err error) {
				if result != nil {
					ctx := webView.JavaScriptGlobalContext()
					t.Errorf("want result == nil, got %q", ctx.ToStringOrDie(result))
				}
				if !reflect.DeepEqual(wantErr, err) {
					t.Errorf("want error %q, got %q", wantErr, err)
				}
				gtk.MainQuit()
			})
		}
	})

	glib.IdleAdd(func() bool {
		webView.LoadHTML(`<p></p>`, "")
		return false
	})

	gtk.Main()
}

func TestWebView_GetSnapshot(t *testing.T) {
	webView := NewWebView()
	defer webView.Destroy()

	webView.Connect("load-changed", func(_ *glib.Object, loadEvent LoadEvent) {
		switch loadEvent {
		case LoadFinished:
			webView.GetSnapshot(func(img *image.RGBA, err error) {
				if err != nil {
					t.Errorf("GetSnapshot error: %q", err)
				}
				if img.Pix == nil {
					t.Error("!img.Pix")
				}
				if img.Stride == 0 || img.Rect.Max.X == 0 || img.Rect.Max.Y == 0 {
					t.Error("!img.Stride or !img.Rect.Max.X or !img.Rect.Max.Y")
				}
				gtk.MainQuit()
			})
		}
	})

	glib.IdleAdd(func() bool {
		webView.LoadHTML(`<p id=foo>abc</p>`, "")
		return false
	})

	gtk.Main()
}
