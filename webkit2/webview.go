package webkit2

// #include <stdlib.h>
// #include <webkit2/webkit2.h>
// #include <cairo/cairo.h>
//
// static WebKitWebView* to_WebKitWebView(GtkWidget* w) { return WEBKIT_WEB_VIEW(w); }
//
// #cgo pkg-config: webkit2gtk-4.0
import "C"

import (
	"bytes"
	"encoding/binary"
	"errors"
	"image"
	"unsafe"

	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
	"github.com/sqs/gojs"
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
	return (*gojs.Context)(gojs.NewGlobalContextFrom((gojs.RawGlobalContext)(unsafe.Pointer(C.webkit_web_view_get_javascript_global_context(v.webView)))))
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
			C.free(unsafe.Pointer(userData))
			var jserr *C.GError
			jsResult := C.webkit_web_view_run_javascript_finish(v.webView, result, &jserr)
			if jsResult == nil {
				defer C.g_error_free(jserr)
				msg := C.GoString((*C.char)(jserr.message))
				resultCallback(nil, errors.New(msg))
				return
			}
			ctxRaw := gojs.RawGlobalContext(unsafe.Pointer(C.webkit_javascript_result_get_global_context(jsResult)))
			jsValRaw := gojs.RawValue(unsafe.Pointer(C.webkit_javascript_result_get_value(jsResult)))
			ctx := (*gojs.Context)(gojs.NewGlobalContextFrom(ctxRaw))
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

// http://cairographics.org/manual/cairo-cairo-surface-t.html#cairo-surface-type-t
const cairoSurfaceTypeImage = 0

// http://cairographics.org/manual/cairo-Image-Surfaces.html#cairo-format-t
const cairoImageSurfaceFormatARGB32 = 0

// GetSnapshot runs asynchronously, taking a snapshot of the WebView.
// Upon completion, resultCallback will be called with a copy of the underlying
// bitmap backing store for the frame, or with an error encountered during
// execution.
//
// See also: webkit_web_view_get_snapshot at
// http://webkitgtk.org/reference/webkit2gtk/stable/WebKitWebView.html#webkit-web-view-get-snapshot
func (v *WebView) GetSnapshot(resultCallback func(result *image.RGBA, err error)) {
	var cCallback C.GAsyncReadyCallback
	var userData C.gpointer
	var err error
	if resultCallback != nil {
		callback := func(result *C.GAsyncResult) {
			C.free(unsafe.Pointer(userData))
			var snapErr *C.GError
			snapResult := C.webkit_web_view_get_snapshot_finish(v.webView, result, &snapErr)
			if snapResult == nil {
				defer C.g_error_free(snapErr)
				msg := C.GoString((*C.char)(snapErr.message))
				resultCallback(nil, errors.New(msg))
				return
			}
			defer C.cairo_surface_destroy(snapResult)

			if C.cairo_surface_get_type(snapResult) != cairoSurfaceTypeImage ||
				C.cairo_image_surface_get_format(snapResult) != cairoImageSurfaceFormatARGB32 {
				panic("Snapshot in unexpected format")
			}

			w := int(C.cairo_image_surface_get_width(snapResult))
			h := int(C.cairo_image_surface_get_height(snapResult))
			stride := int(C.cairo_image_surface_get_stride(snapResult))
			data := unsafe.Pointer(C.cairo_image_surface_get_data(snapResult))
			surfaceBytes := C.GoBytes(data, C.int(stride*h))
			// convert from b,g,r,a or a,r,g,b(local endianness) to r,g,b,a
			testint, _ := binary.ReadUvarint(bytes.NewBuffer([]byte{0x1, 0}))
			if testint == 0x1 {
				// Little: b,g,r,a -> r,g,b,a
				for i := 0; i < w*h; i++ {
					b := surfaceBytes[4*i+0]
					r := surfaceBytes[4*i+2]
					surfaceBytes[4*i+0] = r
					surfaceBytes[4*i+2] = b
				}
			} else {
				// Big: a,r,g,b -> r,g,b,a
				for i := 0; i < w*h; i++ {
					a := surfaceBytes[4*i+0]
					r := surfaceBytes[4*i+1]
					g := surfaceBytes[4*i+2]
					b := surfaceBytes[4*i+3]
					surfaceBytes[4*i+0] = r
					surfaceBytes[4*i+1] = g
					surfaceBytes[4*i+2] = b
					surfaceBytes[4*i+3] = a
				}
			}
			rgba := &image.RGBA{surfaceBytes, stride, image.Rect(0, 0, w, h)}
			resultCallback(rgba, nil)
		}
		cCallback, userData, err = newGAsyncReadyCallback(callback)
		if err != nil {
			panic(err)
		}
	}

	C.webkit_web_view_get_snapshot(v.webView,
		(C.WebKitSnapshotRegion)(1), // FullDocument is the only working region at this point
		(C.WebKitSnapshotOptions)(0),
		nil,
		cCallback,
		userData)
}
