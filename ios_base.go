package umeng

type IOSPayloadAPS struct {
	Alert string `json:"alert"`
	badge string `json:"badge"`
	sound string `json:"sound"`
	ContentAvailable string `json:"content-available"`
	Category string `json:"category"`
}
type IOSPayload struct {
	Alert string `json:"alert"`
	badge string `json:"badge"`
	sound string `json:"sound"`
	ContentAvailable string `json:"content-available"`
	Category string `json:"category"`
	CustomFields map[string]interface{}
}

func (p *IOSPayload) marshalPayload() ([]byte, error) {
	var jsonStr []byte
	return jsonStr, nil
}