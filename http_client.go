package umeng

import (
	"fmt"
	"net/http"
	"bytes"
	"encoding/json"
	"io/ioutil"
)
type Data struct {

}
func Send(url string, p *Data) {
	jstr, err := json.Marshal(p)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jstr))
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status", resp.Status)
	fmt.Println("response Headers", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("Response Body:", string(body))
}

