package webkit2

// #include <webkit2/webkit2.h>
import "C"
import "unsafe"
import "github.com/sqs/gotk3/glib"

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

// GetEnableWriteConsoleMessagesToStdout returns the
// "enable-write-console-messages-to-stdout" property.
//
// See also: webkit_settings_get_enable_write_console_messages_to_stdout at
// http://webkitgtk.org/reference/webkit2gtk/stable/WebKitSettings.html#webkit-settings-get-enable-write-console-messages-to-stdout
func (s *Settings) GetEnableWriteConsoleMessagesToStdout() bool {
	return gobool(C.webkit_settings_get_enable_write_console_messages_to_stdout(s.settings))
}

// SetEnableWriteConsoleMessagesToStdout sets the
// "enable-write-console-messages-to-stdout" property.
//
// See also: webkit_settings_set_enable_write_console_messages_to_stdout at
// http://webkitgtk.org/reference/webkit2gtk/stable/WebKitSettings.html#webkit-settings-set-enable-write-console-messages-to-stdout
func (s *Settings) SetEnableWriteConsoleMessagesToStdout(write bool) {
	C.webkit_settings_set_enable_write_console_messages_to_stdout(s.settings, gboolean(write))
}
