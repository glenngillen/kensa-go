package main

import (
	"fmt"
	"errors"
	"encoding/json"
)

type ManifestAPI struct {
	Regions	 *string   `json:"regions,omitempty"`
}
type Manifest struct {
	Id	 *string   `json:"id,omitempty"`
	Api	 *ManifestAPI   `json:"api,omitempty"`
	Contents []byte
}

func (m Manifest) IsValidJSON() (isValid bool) {
	var dat map[string]interface{}
	if err := json.Unmarshal(m.Contents, &dat); err != nil {
		isValid = false
	} else {
		isValid = true
	}
	return isValid
}

func (m Manifest) IsValid() (isValid bool, err error) {
	isValid = true
	o := &Manifest{}
	if err = json.Unmarshal(m.Contents, &o); err != nil {
		isValid = false
		fmt.Println("%s", err)
		err = errors.New("JSON invalid")
	} else {
		if o.Id == nil {
			isValid = false
			err = errors.New("Missing \"id\"")
		}
		if o.Api == nil {
			isValid = false
			err = errors.New("Missing \"api\"")
		} else {
			if o.Api.Regions == nil {
				isValid = false
				err = errors.New("Missing \"api['regions']\"")
			}
		}
	}
	return isValid, err
}
