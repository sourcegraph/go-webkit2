package webkit2

import (
	"testing"
)

func TestSettings_AutoLoadImages(t *testing.T) {
	s := NewWebView().Settings()

	autoLoad := s.AutoLoadImages()
	wantAutoLoad := !autoLoad
	s.SetAutoLoadImages(wantAutoLoad)

	autoLoad = s.AutoLoadImages()
	if wantAutoLoad != autoLoad {
		t.Errorf("want changed AutoLoad == %d, got %d", wantAutoLoad, autoLoad)
	}

	// Revert to original setting.
	s.SetAutoLoadImages(!autoLoad)
}

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
