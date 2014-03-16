package main

import (
	"encoding/json"

	"testing"
)

func manifestDefinition() (str string) {
	str = `{"id": "addon-name", "api": "", "regions": ""}`
	return str
}

func deleteKey(key string, jsondef string) (str string) {
	var dat map[string]interface{}
	if err := json.Unmarshal([]byte(jsondef), &dat); err != nil {
		panic(err)
	}
	delete(dat, key)
	byt, err := json.Marshal(dat)
	if err != nil {
		panic(err)
	}
	str = string(byt[:])
	return str
}

func testKeyExists(t *testing.T, m Manifest, errMsg string) {
	_, err := m.IsValid()
	if err != nil && err.Error() == errMsg {
		// Successfully checked
	} else {
		t.Errorf("Expected \"%s\" validation error to be raised", errMsg)
	}
}

func TestRejectsInvalidJSON(t *testing.T) {
	m := Manifest{Contents: []byte(`"foo": "adada"}`)}
	if m.IsValidJSON() != false {
		t.Errorf("JSON should not be parsed: '%s'", m.Contents)
	}
}

func TestValidates(t *testing.T) {
	jsonDef := manifestDefinition()
	m := Manifest{Contents: []byte(jsonDef)}
	_, err := m.IsValid()
	if err != nil {
		t.Errorf("Manifest did not validate: %s", err)
	}
}

func TestRequiresId(t *testing.T) {
	jsonDef := manifestDefinition()
	jsonDef = deleteKey("id", jsonDef)
	m := Manifest{Contents: []byte(jsonDef)}
	testKeyExists(t, m, "Missing 'id'")
}

func TestRequiresApi(t *testing.T) {
	jsonDef := manifestDefinition()
	jsonDef = deleteKey("api", jsonDef)
	m := Manifest{Contents: []byte(jsonDef)}
	testKeyExists(t, m, "Missing 'api'")
}

func TestRequiresRegions(t *testing.T) {
	jsonDef := manifestDefinition()
	jsonDef = deleteKey("regions", jsonDef)
	m := Manifest{Contents: []byte(jsonDef)}
	testKeyExists(t, m, "Missing 'regions'")
}
