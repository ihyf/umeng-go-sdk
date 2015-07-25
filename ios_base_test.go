package umeng_test
import (
	"github.com/netroby/umeng-go-sdk"
	"testing"
	"fmt"
)

func TestIOSPayload (t *testing.T) {
	p := &umeng.IOSPayload{
		Alert: "Nomatter",
		Badge: "Nsdf",
		Sound: "default",
		Category: "default",
		ContentAvailable: "default",
		CustomFields: []umeng.CustomFields{
			umeng.CustomFields{"key1", "value1"},
			umeng.CustomFields{"key2", "value2"},
		},
	}
	p.SetExtraField("key3", "value3")
	p.CustomFields = append(p.CustomFields, umeng.CustomFields{"key4", "value4"})

	jStr, _ := p.MarshalPayload()
	fmt.Println(string(jStr))
}