package main

import (
	"fmt"
	"errors"
	"encoding/json"
)

type ManifestAPIEndpoints struct {
	BaseURL		*string	  `json:"base_url,omitempty"`
	SSOUrl		*string	  `json:"sso_url,omitempty"`
}
type ManifestAPI struct {
	Regions		*string   `json:"regions,omitempty"`
	Password	*string   `json:"password,omitempty"`
	Production	*ManifestAPIEndpoints	`json:"production,omitempty"`
	Test		*ManifestAPIEndpoints	`json:"test,omitempty"`
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
			if o.Api.Password == nil {
				isValid = false
				err = errors.New("Missing \"api['password']\"")
			}
			if o.Api.Production == nil {
				isValid = false
				err = errors.New("Missing \"api['production']\"")
			}
			if o.Api.Test == nil {
				isValid = false
				err = errors.New("Missing \"api['test']\"")
			}
		}
	}
	return isValid, err
}
