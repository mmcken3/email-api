package twilio

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// Config contains all items needed for a twilio message.
type Config struct {
	Message    string
	SID        string
	Token      string
	ToNumber   string
	FromNumber string
}

// SendTextMessage does
func SendTextMessage(config Config) {
	urlStr := "https://api.twilio.com/2010-04-01/Accounts/" + config.SID + "/Messages.json"

	// Set up rand
	rand.Seed(time.Now().Unix())

	msgData := url.Values{}
	msgData.Set("To", config.ToNumber)
	msgData.Set("From", config.FromNumber)
	msgData.Set("Body", config.Message)
	msgDataReader := *strings.NewReader(msgData.Encode())

	client := &http.Client{}
	req, _ := http.NewRequest("POST", urlStr, &msgDataReader)
	req.SetBasicAuth(config.SID, config.Token)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, _ := client.Do(req)
	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		var data map[string]interface{}
		decoder := json.NewDecoder(resp.Body)
		err := decoder.Decode(&data)
		if err == nil {
			fmt.Println(data["sid"])
		}
	} else {
		fmt.Println(resp.Status)
	}
	return
}
