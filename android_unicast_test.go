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
