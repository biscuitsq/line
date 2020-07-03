package gopkg

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

func MessageSend(token string, msg string) (string, error) {

	result := ""

	jsonStr := "message=" + "message=" + url.QueryEscape(msg)

	req, err := http.NewRequest(
		"POST",
		"https://notify-api.line.me/api/notify",
		bytes.NewBuffer([]byte(jsonStr)),
	)
	if err != nil {
		log.Fatal(err)
	}

	//ヘッダーつける場合
	req.Host = "notify-api.line.me"
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Authorization", "Bearer "+token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	result = string(byteArray)
	//fmt.Println(result)

	return result, nil
}
