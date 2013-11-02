package webkit2

import (
	"testing"
)

func TestSettings_AutoLoadImages(t *testing.T) {
	s := NewWebView().Settings()

	autoLoad := s.GetAutoLoadImages()
	wantAutoLoad := !autoLoad
	s.SetAutoLoadImages(wantAutoLoad)

	autoLoad = s.GetAutoLoadImages()
	if wantAutoLoad != autoLoad {
		t.Errorf("want changed AutoLoad == %d, got %d", wantAutoLoad, autoLoad)
	}

	// Revert to original setting.
	s.SetAutoLoadImages(!autoLoad)
}

func TestSettings_SetUserAgentWithApplicationDetails(t *testing.T) {
	s := NewWebView().Settings()
	s.SetUserAgentWithApplicationDetails("myApp", "myVersion")
}
