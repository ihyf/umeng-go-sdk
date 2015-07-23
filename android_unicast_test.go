package umeng_test

import (
	"encoding/json"
	"fmt"
	"github.com/netroby/umeng-go-sdk"
	"testing"
)

func TestAndroidPlayBody(t *testing.T) {
	p := new(umeng.AndroidPayloadBody)
	p.Ticker = "hello"
	p.Title = "testtitle"
	p.Text = "TextTest"
	p.Play_vibrate = "text"
	p.Play_lights = "xxlight"
	p.Play_sound = "xxx"
	p.After_open = "xxx"
	jstr, _ := json.Marshal(p)
	fmt.Println(string(jstr))
}
func TestAndroidPayload(t *testing.T) {
	p := &umeng.AndroidPayloadBody{
		Ticker: "hello",
		Title: "testtitle",
		Text: "testText",
		Play_vibrate: "testPlayVibrate",
		Play_lights: "testPlayLights",
		After_open: "testAfterOpen",
	}
	ap := &umeng.AndroidPayload{
		Display_type: "testDisplayType",
		Body: *p,
	}
	jstr, _ := json.Marshal(ap)
	fmt.Println(string(jstr))
}
