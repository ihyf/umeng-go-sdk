package umeng

type IOSPayloadAPS struct {
	Alert string `json:"alert"`
	badge string `json:"badge"`
	sound string `json:"sound"`
	ContentAvailable string `json:"content-available"`
	Category string `json:"category"`
}
type IOSPayload struct {
	APS IOSPayloadAPS `json:"aps"`
}