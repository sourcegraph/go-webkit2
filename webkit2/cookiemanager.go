package webkit2

// #include <stdlib.h>
// #include <webkit2/webkit2.h>
import "C"
import (
	"github.com/gotk3/gotk3/glib"
	"unsafe"
)

// CookieManager — Defines how to handle cookies in a WebContext
//
// See also: WebKitCookieManager at
// http://webkitgtk.org/reference/webkit2gtk/stable/WebKitCookieManager.html.
type CookieManager struct {
	*glib.Object
	cookieManager *C.WebKitCookieManager
}

func newCookieManager(cookieManager *C.WebKitCookieManager) *CookieManager {
	obj := &glib.Object{glib.ToGObject(unsafe.Pointer(cookieManager))}
	return &CookieManager{obj, cookieManager}
}

// CookiePersistentStorage values used to denote the cookie persistent storage types
// are described at
// http://webkitgtk.org/reference/webkit2gtk/stable/WebKitCookieManager.html#WebKitCookiePersistentStorage
type CookiePersistentStorage int

const (
	CookiePersistentStorageText CookiePersistentStorage = iota
	CookiePersistentStorageSqlite
)

// SetPersistentStorage sets the filename where non-session cookies are stored
// persistently using storage as the format to read/write the cookies.
//
// See also: webkit_cookie_manager_set_persistent_storage
// http://webkitgtk.org/reference/webkit2gtk/stable/WebKitCookieManager.html#webkit-cookie-manager-set-persistent-storage
func (cm *CookieManager) SetPersistentStorage(filename string, storage CookiePersistentStorage) {
	cstr := C.CString(filename)
	defer C.free(unsafe.Pointer(cstr))
	C.webkit_cookie_manager_set_persistent_storage(cm.cookieManager,
		(*C.gchar)(cstr),
		C.WebKitCookiePersistentStorage(storage))
}

// CookiePersistentStorage values used to denote the cookie acceptance policies.
// are described at
// http://webkitgtk.org/reference/webkit2gtk/stable/WebKitCookieManager.html#WebKitCookieAcceptPolicy
type CookieAcceptPolicy int

const (
	CookiePolicyAcceptAlways CookieAcceptPolicy = iota
	CookiePolicyAcceptNever
	CookiePolicyAcceptNoThirdParty
)

// SetAcceptPolicy set the cookie acceptance policy of CookieManager as policy .
//
// See also: webkit_cookie_manager_set_accept_policy
// http://webkitgtk.org/reference/webkit2gtk/stable/WebKitCookieManager.html#webkit-cookie-manager-set-accept-policy
func (cm *CookieManager) SetAcceptPolicy(policy CookieAcceptPolicy) {
	C.webkit_cookie_manager_set_accept_policy(cm.cookieManager,
		C.WebKitCookieAcceptPolicy(policy))
}

// DeleteCookiesForDomain Remove all cookies of CookieManager for the given domain.
//
// See also: webkit_cookie_manager_delete_cookies_for_domain
// http://webkitgtk.org/reference/webkit2gtk/stable/WebKitCookieManager.html#webkit-cookie-manager-delete-cookies-for-domain
func (cm *CookieManager) DeleteCookiesForDomain(domain string) {
	cstr := C.CString(domain)
	defer C.free(unsafe.Pointer(cstr))
	C.webkit_cookie_manager_delete_cookies_for_domain(cm.cookieManager,
		(*C.gchar)(cstr))
}

// DeleteAllCookies delete all cookies of CookieManager.
//
// See also: webkit_cookie_manager_delete_all_cookies
// http://webkitgtk.org/reference/webkit2gtk/stable/WebKitCookieManager.html#webkit-cookie-manager-delete-all-cookies
func (cm *CookieManager) DeleteAllCookies(domain string) {
	C.webkit_cookie_manager_delete_all_cookies(cm.cookieManager)
}
