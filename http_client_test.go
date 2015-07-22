package umeng_test
import (
	"github.com/netroby/umeng-go-sdk"
	"testing"
)
func TestSend(t *testing.T) {
	umeng.Send("string", nil)
}