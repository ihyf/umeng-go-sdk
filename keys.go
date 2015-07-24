package umeng

var (
	DATA_KEYS = []string{
		"appkey",
		"timestamp",
		"type",
		"device_tokens",
		"alias",
		"alias_type",
		"file_id",
		"filter",
		"production_model",
		"feedback",
		"description",
		"thirdparty_id",
	}
	POLICY_KEYS = []string{
		"start_time",
		"expire_time",
		"max_send_num",
	}
	Host            string = "http://msg.umeng.com"
	UploadPath      string = "/upload"
	PostPath        string = "/api/send"
	AppMasterSecret string
)

//The func which will export should begin with Uppercase word
func KInSlice(a string, b []string) bool {
	for i := range b {
		if a == b[i] {
			return true
		}
	}
	return false
}
