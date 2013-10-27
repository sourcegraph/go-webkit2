package webkit2

import (
	"github.com/sqs/gotk3/gtk"
	"testing"
)


func TestNewSettings(t *testing.T) {
	NewSettings()
}

func TestSettings_AutoLoadImages(t *testing.T) {
	s := NewSettings()

	autoLoad := s.AutoLoadImages()
	wantAutoLoad := !autoLoad
	s.SetAutoLoadImages(wantAutoLoad)

	autoLoad = s.AutoLoadImages()
	if wantAutoLoad != autoLoad {
		t.Errorf("want changed AutoLoad == %d, got %d", wantAutoLoad, autoLoad)
	}
}
