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

// SendTextMessage does
func SendTextMessage(name string, email string, message string, site string, sid string, token string) {
	urlStr := "https://api.twilio.com/2010-04-01/Accounts/" + sid + "/Messages.json"

	// Create message body
	msgBody := fmt.Sprintf("You have been contacted by %v. %v. Message: %v. Site: %v", name, email, message, site)

	// Set up rand
	rand.Seed(time.Now().Unix())

	msgData := url.Values{}
	msgData.Set("To", fmt.Sprintf("+18437373287"))
	msgData.Set("From", "+17123509887")
	msgData.Set("Body", msgBody)
	msgDataReader := *strings.NewReader(msgData.Encode())

	client := &http.Client{}
	req, _ := http.NewRequest("POST", urlStr, &msgDataReader)
	req.SetBasicAuth(sid, token)
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
