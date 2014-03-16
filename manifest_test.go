package main

import (
	"fmt"
	"encoding/json"
	"strings"

	"testing"
)

func manifestDefinition() (str string) {
	str = `{"id": "addon-name",
		"api": {
			"regions": ""
		}}`
	return str
}

func deleteKey(key string, jsondef string) (str string) {
	var dat map[string]interface{}
	if err := json.Unmarshal([]byte(jsondef), &dat); err != nil {
		panic(err)
	}
	if key == "api['regions']" {
		api := dat["api"].(map[string]interface{})
		delete(api, "regions")
		dat["api"] = api
	} else {
		delete(dat, key)
	}
	byt, err := json.Marshal(dat)
	if err != nil {
		panic(err)
	}
	str = string(byt[:])
	return str
}

func testKeyExists(t *testing.T, jsondef string, key string) {
	s := []string{"Missing \"", key, "\""}
	errMsg := strings.Join(s, "")
	jsondef = deleteKey(key, jsondef)
	m := Manifest{Contents: []byte(jsondef)}
	_, err := m.IsValid()
	if err != nil && err.Error() == errMsg {
		// Successfully checked
	} else {
		fmt.Println(err)
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
	testKeyExists(t, jsonDef, "id")
}

func TestRequiresApi(t *testing.T) {
	jsonDef := manifestDefinition()
	testKeyExists(t, jsonDef, "api")
}

func TestRequiresApiRegions(t *testing.T) {
	jsonDef := manifestDefinition()
	testKeyExists(t, jsonDef, "api['regions']")
}
