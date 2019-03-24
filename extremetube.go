package extremetube

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

const apiURL = "https://www.extremetube.com/api/HubTrafficApiCall?"
const APITimeout = 5

func GetVideoByID(ID string) ExtremetubeSingleVideo {
	timeout := time.Duration(APITimeout * time.Second)
	client := http.Client{
		Timeout: timeout,
	}
	resp, _ := client.Get(fmt.Sprintf(apiURL+"data=getVideoById&output=json&video_id=%s", ID))
	b, _ := ioutil.ReadAll(resp.Body)
	var result ExtremetubeSingleVideo
	err := json.Unmarshal(b, &result)
	if err != nil {
		log.Println(err)
	}
	return result

}

func GetVideoEmbedCode(ID string) ExtremetubeEmbedCode {
	timeout := time.Duration(APITimeout * time.Second)
	client := http.Client{
		Timeout: timeout,
	}
	resp, _ := client.Get(fmt.Sprintf(apiURL+"data=getVideoEmbedCode&video_id=%s", ID))
	b, _ := ioutil.ReadAll(resp.Body)
	var result ExtremetubeEmbedCode
	err := json.Unmarshal(b, &result)
	if err != nil {
		log.Println(err)
	}
	return result
}
