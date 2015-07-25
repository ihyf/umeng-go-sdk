package umeng

import (
	"encoding/json"
	"errors"
)

var (
	APS_KEYS = []string{
		"alert",
		"badge",
		"sound",
		"content-available",
	}
)

type CustomFields struct {
	Key string
	Val string
}
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
	CustomFields     []CustomFields
}
type IOSUnicastNotification struct {
	AppKey       string `json:"appkey"`
	Timestamp    string `json:"timestamp"`
	Type         string `json:"Type"`
	DeviceTokens string `json:"device_tokens"`
	AliasType    string `json:"alias_type"`
}
type IOSUnicast struct {
	AppMasterSecret string
	Notification    IOSUnicastNotification
}

func (c *IOSUnicast) SetAppMasterSecret(secret string) {
	c.AppMasterSecret = secret
}
func (c *IOSUnicast) SetPredefinedKeyValue(key, value string) {

}

/*
	SetExtraField For IOSPayload
*/
func (p *IOSPayload) SetExtraField(key, value string) {
	p.CustomFields = append(p.CustomFields, CustomFields{key, value})
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
func constructFullPayload(aps interface{}, customFields []CustomFields) (map[string]interface{}, error) {
	var fullPayload = make(map[string]interface{})
	fullPayload["aps"] = aps
	for i := range customFields {
		if customFields[i].Key == "aps" {
			return nil, errors.New("Cannot have a custom field named aps")
		}
		fullPayload[customFields[i].Key] = customFields[i].Val
	}
	return fullPayload, nil
}
