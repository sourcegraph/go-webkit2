// +build !webkit2_2_1,!webkit2_2_0

package webkit2

import (
	"testing"
)

func TestSettings_EnableWriteConsoleMessagesToStdout(t *testing.T) {
	s := NewWebView().Settings()

	write := s.GetEnableWriteConsoleMessagesToStdout()
	wantWrite := !write
	s.SetEnableWriteConsoleMessagesToStdout(wantWrite)

	write = s.GetEnableWriteConsoleMessagesToStdout()
	if wantWrite != write {
		t.Errorf("want changed Write == %d, got %d", wantWrite, write)
	}

	// Revert to original setting.
	s.SetEnableWriteConsoleMessagesToStdout(!write)
}
