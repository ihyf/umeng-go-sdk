package umeng

import (
	"encoding/json"
	"errors"
)

type IOSPayloadAPS struct {
	Alert            string `json:"alert"`
	Badge            string `json:"badge"`
	Sound            string `json:"sound"`
	Category         string `json:"category"`
	ContentAvailable string `json:"content-available"`
}
type IOSPayload struct {
	Alert            string `json:"alert"`
	Badge            string `json:"badge"`
	Sound            string `json:"sound"`
	Category         string `json:"category"`
	ContentAvailable string `json:"content-available"`
	CustomFields     map[string]interface{}
}

func (p *IOSPayload) MarshalPayload() ([]byte, error) {
	var jsonStr []byte
	//use simple payload
	aps := IOSPayloadAPS{
		Alert:            p.Alert,
		Badge:            p.Badge,
		Sound:            p.Sound,
		Category:         p.Category,
		ContentAvailable: p.ContentAvailable,
	}

	fullPayload, err := constructFullPayload(aps, p.CustomFields)
	if err != nil {
		return nil, err
	}

	jsonStr, err = json.Marshal(fullPayload)
	if err != nil {
		return nil, err
	}
	fullPayload["aps"] = aps
	if err != nil {
		return nil, err
	}

	jsonStr, err = json.Marshal(fullPayload)
	if err != nil {
		return nil, err
	}
	return jsonStr, nil
}

//Helper method to generate a json compatible map with aps key + custom fields
//will return error if custom field named aps supplied
func constructFullPayload(aps interface{}, customFields map[string]interface{}) (map[string]interface{}, error) {
	var fullPayload = make(map[string]interface{})
	fullPayload["aps"] = aps
	for key, value := range customFields {
		if key == "aps" {
			return nil, errors.New("Cannot have a custom field named aps")
		}
		fullPayload[key] = value
	}
	return fullPayload, nil
}
