package umeng_test

import (
	"fmt"
	"github.com/netroby/umeng-go-sdk"
	"testing"
)

func TestVar(t *testing.T) {
	fmt.Println(umeng.DATA_KEYS)
	fmt.Println(umeng.POLICY_KEYS)
	fmt.Println(umeng.Host)
	fmt.Println(umeng.UploadPath)
	fmt.Println(umeng.PostPath)
	fmt.Println(umeng.AppMasterSecret)
}

func TestKInSlice(t *testing.T) {
	fmt.Println(umeng.KInSlice("appkey", umeng.DATA_KEYS))     //This will return true
	fmt.Println(umeng.KInSlice("novalidkey", umeng.DATA_KEYS)) //This will return false
}
