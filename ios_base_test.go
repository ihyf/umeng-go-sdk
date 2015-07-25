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
		CustomFields:map[string]interface{}{
			"key1":"value1",
			"key2":"value2",
		},
	}
	jStr, _ := p.MarshalPayload()
	fmt.Println(string(jStr))
}