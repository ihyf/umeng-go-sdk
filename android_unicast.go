package umeng

import (
//"fmt"
)

type AndroidPayloadBody struct {
	Ticker       string `json:"ticker"`
	Title        string `json:"title"`
	Text         string `json:"text"`
	Play_vibrate string `json:"play_vibrate"`
	Play_lights  string `json:"play_lights"`
	Play_sound   string `json:"play_sound"`
	After_open   string `json:"after_open"`
}
type AndroidPayload struct {
	Display_type string             `json:"display_type"`
	Body         AndroidPayloadBody `json:"body"`
}
type PayloadKeys struct {
	Display_type string `json:"display_type"`
}
