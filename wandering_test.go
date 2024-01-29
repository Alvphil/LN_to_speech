package main

import (
	"testing"
)

func TestWebDomain(t *testing.T) {
	cleaned := cleanDomain("https://wanderinginn.com/2022/06/18/interlude-singing-ships/")
	if !(cleaned == "wanderinginn.com") {
		t.Errorf("expected to get wanderinginn.com, got %s", cleaned)
	}
}
