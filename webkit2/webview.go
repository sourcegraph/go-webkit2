package webkit2

// #include <stdlib.h>
// #include <webkit2/webkit2.h>
//
// static WebKitWebView* to_WebKitWebView(GtkWidget* w) { return WEBKIT_WEB_VIEW(w); }
//
// #cgo pkg-config: webkit2gtk-3.0
import "C"

import (
	"errors"
	"github.com/sqs/gojs"
	"github.com/sqs/gotk3/glib"
	"github.com/sqs/gotk3/gtk"
	"unsafe"
)

// WebView represents a WebKit WebView.
//
// See also: WebView at
// http://webkitgtk.org/reference/webkit2gtk/stable/WebKitWebView.html.
type WebView struct {
	*gtk.Widget
	webView *C.WebKitWebView
}

// NewWebView creates a new WebView with the default WebContext and the default
// WebViewGroup.
//
// See also: webkit_web_view_new at
// http://webkitgtk.org/reference/webkit2gtk/stable/WebKitWebView.html#webkit-web-view-new.
func NewWebView() *WebView {
	return newWebView(C.webkit_web_view_new())
}

// NewWebViewWithContext creates a new WebView with the given WebContext and the
// default WebViewGroup.
//
// See also: webkit_web_view_new_with_context at
// http://webkitgtk.org/reference/webkit2gtk/stable/WebKitWebView.html#webkit-web-view-new-with-context.
func NewWebViewWithContext(ctx *WebContext) *WebView {
	return newWebView(C.webkit_web_view_new_with_context(ctx.webContext))
}

func newWebView(webViewWidget *C.GtkWidget) *WebView {
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(webViewWidget))}
	return &WebView{&gtk.Widget{glib.InitiallyUnowned{obj}}, C.to_WebKitWebView(webViewWidget)}
}

// Context returns the current WebContext of the WebView.
//
// See also: webkit_web_view_get_context at
// http://webkitgtk.org/reference/webkit2gtk/stable/WebKitWebView.html#webkit-web-view-get-context.
func (v *WebView) Context() *WebContext {
	return &WebContext{C.webkit_web_view_get_context(v.webView)}
}

// LoadURI requests loading of the specified URI string.
//
// See also: webkit_web_view_load_uri at
// http://webkitgtk.org/reference/webkit2gtk/stable/WebKitWebView.html#webkit-web-view-load-uri
func (v *WebView) LoadURI(uri string) {
	C.webkit_web_view_load_uri(v.webView, (*C.gchar)(C.CString(uri)))
}

// LoadHTML loads the given content string with the specified baseURI. The MIME
// type of the document will be "text/html".
//
// See also: webkit_web_view_load_html at
// http://webkitgtk.org/reference/webkit2gtk/stable/WebKitWebView.html#webkit-web-view-load-html
func (v *WebView) LoadHTML(content, baseURI string) {
	C.webkit_web_view_load_html(v.webView, (*C.gchar)(C.CString(content)), (*C.gchar)(C.CString(baseURI)))
}

// Settings returns the current active settings of this WebView's WebViewGroup.
//
// See also: webkit_web_view_get_settings at
// http://webkitgtk.org/reference/webkit2gtk/stable/WebKitWebView.html#webkit-web-view-get-settings.
func (v *WebView) Settings() *Settings {
	return newSettings(C.webkit_web_view_get_settings(v.webView))
}

// Title returns the current active title of the WebView.
//
// See also: webkit_web_view_get_title at
// http://webkitgtk.org/reference/webkit2gtk/stable/WebKitWebView.html#webkit-web-view-get-title.
func (v *WebView) Title() string {
	return C.GoString((*C.char)(C.webkit_web_view_get_title(v.webView)))
}

// URI returns the current active URI of the WebView.
//
// See also: webkit_web_view_get_uri at
// http://webkitgtk.org/reference/webkit2gtk/stable/WebKitWebView.html#webkit-web-view-get-uri.
func (v *WebView) URI() string {
	return C.GoString((*C.char)(C.webkit_web_view_get_uri(v.webView)))
}

// JavaScriptGlobalContext returns the global JavaScript context used by
// WebView.
//
// See also: webkit_web_view_get_javascript_global_context at
// http://webkitgtk.org/reference/webkit2gtk/stable/WebKitWebView.html#webkit-web-view-get-javascript-global-context
func (v *WebView) JavaScriptGlobalContext() *gojs.Context {
	return gojs.NewContextFrom(gojs.RawContext(C.webkit_web_view_get_javascript_global_context(v.webView)))
}

// RunJavaScript runs script asynchronously in the context of the current page
// in the WebView. Upon completion, resultCallback will be called with the
// result of evaluating the script, or with an error encountered during
// execution. To get the stack trace and other error logs, use the
// ::console-message signal.
//
// See also: webkit_web_view_run_javascript at
// http://webkitgtk.org/reference/webkit2gtk/stable/WebKitWebView.html#webkit-web-view-run-javascript
func (v *WebView) RunJavaScript(script string, resultCallback func(result *gojs.Value, err error)) {
	var cCallback C.GAsyncReadyCallback
	var userData C.gpointer
	var err error
	if resultCallback != nil {
		callback := func(result *C.GAsyncResult) {
			var jserr *C.GError
			jsResult := C.webkit_web_view_run_javascript_finish(v.webView, result, &jserr)
			if jsResult == nil {
				defer C.g_error_free(jserr)
				msg := C.GoString((*C.char)(jserr.message))
				resultCallback(nil, errors.New(msg))
				return
			}
			ctxRaw := gojs.RawContext(C.webkit_javascript_result_get_global_context(jsResult))
			jsValRaw := gojs.RawValue(C.webkit_javascript_result_get_value(jsResult))
			ctx := gojs.NewContextFrom(ctxRaw)
			jsVal := ctx.NewValueFrom(jsValRaw)
			resultCallback(jsVal, nil)
		}
		cCallback, userData, err = newGAsyncReadyCallback(callback)
		if err != nil {
			panic(err)
		}
	}
	C.webkit_web_view_run_javascript(v.webView, (*C.gchar)(C.CString(script)), nil, cCallback, userData)
}

// Destroy destroys the WebView's corresponding GtkWidget and marks its internal
// WebKitWebView as nil so that it can't be accidentally reused.
func (v *WebView) Destroy() {
	v.Widget.Destroy()
	v.webView = nil
}

// LoadEvent denotes the different events that happen during a WebView load
// operation.
//
// See also: WebKitLoadEvent at
// http://webkitgtk.org/reference/webkit2gtk/stable/WebKitWebView.html#WebKitLoadEvent.
type LoadEvent int

// LoadEvent enum values are described at
// http://webkitgtk.org/reference/webkit2gtk/stable/WebKitWebView.html#WebKitLoadEvent.
const (
	LoadStarted LoadEvent = iota
	LoadRedirected
	LoadCommitted
	LoadFinished
)
