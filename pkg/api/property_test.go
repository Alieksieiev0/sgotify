package api

import (
	"testing"
)

func TestCreatBodyFromProperties(t *testing.T) {
	ps := []Property{DeviceIds([]string{"test"}), Play(false)}
	b, err := createBodyFromProperties(ps)
	if err != nil {
		t.Fatal(err)
	}

	if len(ps) == 0 {
		t.Errorf("properties was not added to body map")
	}

	if len(b) == 0 {
		t.Errorf("body was not created from properties")
	}
}
