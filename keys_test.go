package umeng_test
import (
	"testing"
	"fmt"
	"github.com/netroby/umeng-go-sdk"
)

func TestVar (t *testing.T) {
	fmt.Println(umeng.DATA_KEYS)
	fmt.Println(umeng.POLICY_KEYS)
	fmt.Println(umeng.Host)
	fmt.Println(umeng.UploadPath)
	fmt.Println(umeng.PostPath)
	fmt.Println(umeng.AppMasterSecret)
}