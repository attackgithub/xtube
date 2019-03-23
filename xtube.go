package xtube

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

const apiURL = "http://www.xtube.com/webmaster/api.php"
const APITimeout = 5

func GetVideoByID(ID string) XtubeVideo {
	timeout := time.Duration(APITimeout * time.Second)
	client := http.Client{
		Timeout: timeout,
	}
	resp, _ := client.Get(fmt.Sprintf(apiURL+"?action=getVideoById&video_id=%s", ID))
	b, _ := ioutil.ReadAll(resp.Body)
	var result XtubeVideo
	err := json.Unmarshal(b, &result)
	if err != nil {
		log.Println(err)
	}
	return result
}
