// This file includes wrappers for symbols included since WebKit2GTK+ 2.2, and
// and should not be included in a build intended to target any older WebKit2
// versions.  To target an older build, such as 2.1, use
// 'go build -tags webkit2_2_1'.
// +build !webkit2_2_1,!webkit2_2_0

package webkit2

// #include <webkit2/webkit2.h>
import "C"

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
