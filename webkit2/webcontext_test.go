package webkit2

import (
	"testing"
)

func TestDefaultWebContext(t *testing.T) {
	DefaultWebContext()
}

func TestWebContext_CacheModel(t *testing.T) {
	wc := DefaultWebContext()

	// WebBrowserCacheModel is the default, per
	// http://webkitgtk.org/reference/webkit2gtk/stable/WebKitWebContext.html#webkit-web-context-set-cache-model.
	wantCacheModel := WebBrowserCacheModel
	cacheModel := wc.CacheModel()
	if wantCacheModel != cacheModel {
		t.Errorf("want default CacheModel == %d, got %d", wantCacheModel, cacheModel)
	}

	wantCacheModel = DocumentViewerCacheModel
	wc.SetCacheModel(DocumentViewerCacheModel)
	cacheModel = wc.CacheModel()
	if wantCacheModel != cacheModel {
		t.Errorf("want changed CacheModel == %d, got %d", wantCacheModel, cacheModel)
	}
}

func TestWebContext_ClearCache(t *testing.T) {
	DefaultWebContext().ClearCache()
}
