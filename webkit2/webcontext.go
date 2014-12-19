package webkit2

// #include <stdlib.h>
// #include <webkit2/webkit2.h>
import "C"
import (
	"unsafe"
	"github.com/visionect/gotk3/glib"
)
// WebContext manages all aspects common to all WebViews.
//
// See also: WebKitWebContext at
// http://webkitgtk.org/reference/webkit2gtk/stable/WebKitWebContext.html.
type WebContext struct {
	*glib.Object
	webContext *C.WebKitWebContext
}

func newWebContext(webContext *C.WebKitWebContext) *WebContext {
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(webContext))}
	return &WebContext{obj, webContext}
}

// DefaultWebContext returns the default WebContext.
//
// See also: webkit_web_context_get_default at
// http://webkitgtk.org/reference/webkit2gtk/stable/WebKitWebContext.html#webkit-web-context-get-default.
func DefaultWebContext() *WebContext {
	return newWebContext(C.webkit_web_context_get_default())
}

// CacheModel describes the caching behavior.
//
// See also: WebKitCacheModel at
// http://webkitgtk.org/reference/webkit2gtk/stable/WebKitWebContext.html#WebKitCacheModel.
type CacheModel int

// CacheModel enum values are described at
// http://webkitgtk.org/reference/webkit2gtk/stable/WebKitWebContext.html#WebKitCacheModel.
const (
	DocumentViewerCacheModel CacheModel = iota
	WebBrowserCacheModel
	DocumentBrowserCacheModel
)

// CacheModel returns the current cache model.
//
// See also: webkit_web_context_get_cache_model at
// http://webkitgtk.org/reference/webkit2gtk/stable/WebKitWebContext.html#webkit-web-context-get-cache-model.
func (wc *WebContext) CacheModel() CacheModel {
	return CacheModel(C.int(C.webkit_web_context_get_cache_model(wc.webContext)))
}

// SetCacheModel sets the current cache model.
//
// See also: webkit_web_context_set_cache_model at
// http://webkitgtk.org/reference/webkit2gtk/stable/WebKitWebContext.html#webkit-web-context-set-cache-model.
func (wc *WebContext) SetCacheModel(model CacheModel) {
	C.webkit_web_context_set_cache_model(wc.webContext, C.WebKitCacheModel(model))
}

// ClearCache clears all resources currently cached.
//
// See also: webkit_web_context_clear_cache at
// http://webkitgtk.org/reference/webkit2gtk/stable/WebKitWebContext.html#webkit-web-context-clear-cache.
func (wc *WebContext) ClearCache() {
	C.webkit_web_context_clear_cache(wc.webContext)
}

// SetDiskCacheDirectory sets the directory where disk cache files will be stored .
//
// See also: webkit_web_context_set_disk_cache_directory
// http://webkitgtk.org/reference/webkit2gtk/stable/WebKitWebContext.html#webkit-web-context-set-disk-cache-directory
func (wc *WebContext) SetDiskCacheDirectory(directory string) {
	cstr := C.CString(directory)
	defer C.free(unsafe.Pointer(cstr))
	C.webkit_web_context_set_disk_cache_directory(wc.webContext, (*C.gchar)(cstr))
}

func (wc *WebContext) GetCookieManager() *CookieManager {
	return newCookieManager(C.webkit_web_context_get_cookie_manager(wc.webContext))
}

// SetWebExtensionsDirectory sets the directory where WebKit will look for Web Extensions.
//
// See also: webkit_web_context_set_web_extensions_directory
// http://webkitgtk.org/reference/webkit2gtk/stable/WebKitWebContext.html#webkit-web-context-set-web-extensions-directory
func (wc *WebContext) SetWebExtensionsDirectory(directory string) {
	cstr := C.CString(directory)
	defer C.free(unsafe.Pointer(cstr))
	C.webkit_web_context_set_web_extensions_directory(wc.webContext, (*C.gchar)(cstr))
}

