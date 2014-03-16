package main

import (
	"testing"
)

func TestRejectsInvalidJSON(t *testing.T) {
	m := Manifest{Contents: []byte(`"foo": "adada"}`)}
	if m.IsValidJSON() != false {
		t.Errorf("JSON should not be parsed: '%s'", m.Contents)
	}
}

func TestRequiresId(t *testing.T) {
	m := Manifest{Contents: []byte(`{"foo": "adada"}`)}
	_, err := m.IsValid()
	if err != nil && err.Error() == "Missing ID" {
		// Successfully checked for ID
	} else {
		t.Errorf("Should require 'id' property to be defined")
	}
}
