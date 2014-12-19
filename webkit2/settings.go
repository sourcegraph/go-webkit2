package webkit2

// #include <stdlib.h>
// #include <webkit2/webkit2.h>
import "C"
import "unsafe"
import "github.com/visionect/gotk3/glib"

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

// GetEnableDeveloperExtras returns the "enable-developer-extras" property.
//
// See also: webkit_settings_set_enable_developer_extras at
// http://webkitgtk.org/reference/webkit2gtk/stable/WebKitSettings.html#webkit-settings-get-enable-developer-extras
func (s *Settings) GetEnableDeveloperExtras() bool {
	return gobool(C.webkit_settings_get_enable_developer_extras(s.settings))
}

// SetEnableDeveloperExtras sets the "enable-developer-extras" property.
//
// See also: webkit_settings_set_enable_developer_extras at
// http://webkitgtk.org/reference/webkit2gtk/stable/WebKitSettings.html#webkit-settings-set-enable-developer-extras
func (s *Settings) SetEnableDeveloperExtras(autoLoad bool) {
	C.webkit_settings_set_enable_developer_extras(s.settings, gboolean(autoLoad))
}

// SetUserAgentWithApplicationDetails sets the "user-agent" property by
// appending the application details to the default user agent.
//
// See also: webkit_settings_set_user_agent_with_application_details at
// http://webkitgtk.org/reference/webkit2gtk/unstable/WebKitSettings.html#webkit-settings-set-user-agent-with-application-details
func (s *Settings) SetUserAgentWithApplicationDetails(appName, appVersion string) {
	cName := C.CString(appName)
	defer C.free(unsafe.Pointer(cName))
	cVersion := C.CString(appVersion)
	defer C.free(unsafe.Pointer(cVersion))
	C.webkit_settings_set_user_agent_with_application_details(s.settings, (*C.gchar)(cName), (*C.gchar)(cVersion))
}
