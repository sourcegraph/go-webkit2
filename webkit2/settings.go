package webkit2

// #include <webkit2/webkit2.h>
import "C"
import "unsafe"
import "github.com/gotk3/gotk3/glib"

type Settings struct {
	*glib.Object
	settings *C.WebKitSettings
}

// newSettings creates a new Settings with default values.
//
// See also: webkit_settings_new at
// http://webkitgtk.org/reference/webkit2gtk/stable/WebKitSettings.html#webkit-settings-new.
func newSettings(settings *C.WebKitSettings) *Settings {
	return &Settings{&glib.Object{glib.ToGObject(unsafe.Pointer(settings))}, settings}
}

// GetAutoLoadImages returns the "auto-load-images" property.
//
// See also: webkit_settings_get_auto_load_images at
// http://webkitgtk.org/reference/webkit2gtk/stable/WebKitSettings.html#webkit-settings-get-auto-load-images
func (s *Settings) GetAutoLoadImages() bool {
	return gobool(C.webkit_settings_get_auto_load_images(s.settings))
}

// SetAutoLoadImages sets the "auto-load-images" property.
//
// See also: webkit_settings_get_auto_load_images at
// http://webkitgtk.org/reference/webkit2gtk/stable/WebKitSettings.html#webkit-settings-set-auto-load-images
func (s *Settings) SetAutoLoadImages(autoLoad bool) {
	C.webkit_settings_set_auto_load_images(s.settings, gboolean(autoLoad))
}

// SetUserAgentWithApplicationDetails sets the "user-agent" property by
// appending the application details to the default user agent.
//
// See also: webkit_settings_set_user_agent_with_application_details at
// http://webkitgtk.org/reference/webkit2gtk/unstable/WebKitSettings.html#webkit-settings-set-user-agent-with-application-details
func (s *Settings) SetUserAgentWithApplicationDetails(appName, appVersion string) {
	C.webkit_settings_set_user_agent_with_application_details(s.settings, (*C.gchar)(C.CString(appName)), (*C.gchar)(C.CString(appVersion)))
}
